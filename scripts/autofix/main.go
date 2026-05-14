// MIT License
// 
// Copyright (c) 2020–2026
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package main provides an auto-fixer for common Go test syntax errors.
// It iteratively parses files, detects known error patterns, applies fixes, and re-checks.
// Usage: go run ./scripts/autofix/ [--dry-run] [files or dirs...]
// If no args, defaults to tests/integratedtests/corestrtests/
package main

import (
	"fmt"
	"go/parser"
	"go/scanner"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var dryRun bool

const rulesReference = `────────────────────────────────────────────────────────────────────────────────
 SYNTAX RULES REFERENCE
────────────────────────────────────────────────────────────────────────────────

 BRACECHECK — Diagnostic Categories (scripts/bracecheck/main.go)
 ───────────────────────────────────────────────────────────────

 1. Missing trailing comma
    Pattern:  "missing ',' before newline in argument list"
    Cause:    Multi-line calls/literals need a trailing comma after the last arg.
    Example:
        assert.Equal(t,
            expected,
            actual     // ← missing comma
        )

 2. Unexpected closing paren (inside block)
    Pattern:  "expected statement, found ')'"
    Cause:    Extra ')' inside a closure, or '}' and ')' split across lines.
    Example:
        }
        )    // ← should be '})'

 3. Unexpected ')' at top level
    Pattern:  "expected declaration, found ')'"
    Cause:    Stray ')' at package level from mismatched safeTest closure.

 4. Expression count mismatch
    Pattern:  "expected 1 expression"
    Cause:    Return/assignment has wrong number of values on one side.
    Example:  return a, b    // but func returns 1 value

 5. Unclosed brace at EOF
    Pattern:  "expected '}', found 'EOF'"
    Cause:    A '{' block never closed before file end.

 6. Semicolon expected, comma found
    Pattern:  "expected ';', found ','"
    Cause:    A '}' closing a code block (for/if/func) has a stray trailing ','
              making the parser treat subsequent code as a composite literal.

 7. Other / unclassified
    Any Go parser error not matching the above patterns.

 AUTOFIX — Auto-Repair Rules (scripts/autofix/main.go)
 ─────────────────────────────────────────────────────

 A. missing-trailing-comma
    Trigger:  "missing ',' before newline in argument list"
    Action:   Appends ',' to previous non-empty/non-comment line.
    Before:   actual        After:   actual,

 B. unexpected-close-paren
    Trigger:  "expected statement, found ')'"
    Action:   Merges split '}' + ')' → '})'; removes stray ')'; or
              changes '})' → '}' when '}' closed a struct/map literal
              (not a func body) so ')' is extraneous.
    Before:   }             After:   })       (split lines)
              )
    Before:   })            After:   }        (struct close, not func close)

 C. stray-top-level-paren
    Trigger:  "expected declaration, found ')'"
    Action:   Removes bare ')' at top level.

 D. missing-close-brace-eof
    Trigger:  "expected '}', found 'EOF'"
    Action:   Appends '}' to end of file.

 E. expected-one-expression
    Trigger:  "expected 1 expression"
    Action:   Removes trailing comma from returns ("return nil," → "return nil")
              or trims multi-value return to first value ("return a, b" → "return a").

 F. expected-operand
    Trigger:  "expected operand, found '<token>'"
    Action:   Fixes double commas ("a,, b" → "a, b"), empty argument slots
              ("func(a, , b)" → "func(a, b)"), dangling operators (merges
              with next line), and stray tokens where operand expected.

 G. semicolon-expected-comma-found
    Trigger:  "expected ';', found ','"
    Action:   Removes trailing ',' from '},' lines where '}' closes a code block.
    Before:   },            After:   }

 H. stray-comma-on-statement (pre-pass, no parser trigger)
    Trigger:  Lines matching ':= <expr>,' or 'var <ident> <type>,' inside func bodies.
    Action:   Removes the trailing comma that would cause cascading parser errors.
    Before:   x := foo(),   After:   x := foo()

 safeTest CLOSURE PATTERN
 ────────────────────────
    Correct:   safeTest(t, "name", func() { ... })
    Wrong:     }  then  )  on separate lines → merge to  })
    Wrong:     })}  → normalize to  })
    Wrong:     stray  )  after  })  }  → remove the  )

────────────────────────────────────────────────────────────────────────────────

`

// fixRecord tracks a single fix applied or detected.
type fixRecord struct {
	File    string
	Line    int
	Rule    string
	Message string
}

var allRecords []fixRecord

func addRecord(file string, line int, rule, message string) {
	allRecords = append(allRecords, fixRecord{
		File:    file,
		Line:    line,
		Rule:    rule,
		Message: message,
	})
}

func main() {
	var args []string
	for _, a := range os.Args[1:] {
		if a == "--dry-run" {
			dryRun = true
		} else {
			args = append(args, a)
		}
	}
	if len(args) == 0 {
		args = []string{"tests/integratedtests/corestrtests/"}
	}

	var files []string
	for _, arg := range args {
		info, err := os.Stat(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s: %v\n", arg, err)
			os.Exit(1)
		}
		if info.IsDir() {
			entries, _ := os.ReadDir(arg)
			for _, e := range entries {
				if !e.IsDir() && strings.HasSuffix(e.Name(), ".go") {
					files = append(files, filepath.Join(arg, e.Name()))
				}
			}
		} else {
			files = append(files, arg)
		}
	}

	totalFixed := 0
	totalFiles := 0

	for _, f := range files {
		n := fixFile(f)
		if n > 0 {
			totalFiles++
			totalFixed += n
			if dryRun {
				fmt.Printf("  → %s: %d fix(es) would be applied\n", f, n)
			} else {
				fmt.Printf("  ✓ %s: %d fix(es) applied\n", f, n)
			}
		}
	}

	if totalFixed == 0 {
		fmt.Println("✓ No fixable issues found.")
	} else if dryRun {
		fmt.Printf("\n→ Would apply %d fix(es) across %d file(s). (dry-run, no files modified)\n", totalFixed, totalFiles)
	} else {
		fmt.Printf("\n✓ Applied %d fix(es) across %d file(s).\n", totalFixed, totalFiles)
		fmt.Println("  Run bracecheck again to verify: go run ./scripts/bracecheck/")
	}

	// Write syntax-issues.txt report to data/coverage/
	writeReport(totalFixed, totalFiles, len(files))
}

// fixFile attempts up to maxPasses of parse-fix cycles on a single file.
func fixFile(path string) int {
	const maxPasses = 10
	totalFixes := 0

	for pass := 0; pass < maxPasses; pass++ {
		src, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "  ✗ %s: read error: %v\n", path, err)
			return totalFixes
		}

		fset := token.NewFileSet()
		_, parseErr := parser.ParseFile(fset, path, src, parser.AllErrors)
		if parseErr == nil {
			return totalFixes // clean
		}

		errList, ok := parseErr.(scanner.ErrorList)
		if !ok {
			return totalFixes
		}

		lines := strings.Split(string(src), "\n")
		fixes := 0

		// Pre-pass: fix stray commas on statement lines (Rule H)
		// These don't have their own parser error — they cause cascading errors
		// on subsequent lines. Must run before the error-driven loop.
		prePassFixes := fixStrayCommaStatements(lines, path)
		fixes += prePassFixes

		// Process errors in reverse line order so line numbers stay valid
		applied := make(map[int]bool) // track lines already modified this pass
		for i := len(errList) - 1; i >= 0; i-- {
			e := errList[i]
			lineIdx := e.Pos.Line - 1
			if lineIdx < 0 || lineIdx >= len(lines) || applied[lineIdx] {
				continue
			}

			fixed := false
			rule := ""
			switch {
			case strings.Contains(e.Msg, "missing ',' before newline in argument list"):
				fixed = fixMissingComma(lines, lineIdx)
				rule = "missing-trailing-comma"

			case strings.Contains(e.Msg, "expected statement, found ')'"):
				fixed = fixUnexpectedCloseParen(lines, lineIdx)
				rule = "unexpected-close-paren"

			case strings.Contains(e.Msg, "expected declaration, found ')'"):
				fixed = fixUnexpectedCloseParenTopLevel(lines, lineIdx)
				rule = "stray-top-level-paren"

			case strings.Contains(e.Msg, "expected '}', found 'EOF'"):
				fixed = fixMissingCloseBrace(lines)
				rule = "missing-close-brace-eof"

			case strings.Contains(e.Msg, "expected 1 expression"):
				fixed = fixExpectedOneExpression(lines, lineIdx)
				rule = "expected-one-expression"

			case strings.Contains(e.Msg, "expected operand, found"):
				fixed = fixExpectedOperand(lines, lineIdx, e.Msg)
				rule = "expected-operand"

			case strings.Contains(e.Msg, "expected ';', found ','"):
				fixed = fixSemicolonExpectedCommaFound(lines, lineIdx)
				rule = "semicolon-expected-comma-found"
			}

			if fixed {
				fixes++
				applied[lineIdx] = true
				addRecord(path, e.Pos.Line, rule, e.Msg)
			}
		}

		if fixes == 0 {
			return totalFixes // no more auto-fixable errors
		}

		totalFixes += fixes

		if dryRun {
			// Don't write changes in dry-run mode; stop after first pass
			return totalFixes
		}

		result := strings.Join(lines, "\n")
		if err := os.WriteFile(path, []byte(result), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "  ✗ %s: write error: %v\n", path, err)
			return totalFixes
		}
	}
	return totalFixes
}

// fixMissingComma adds a trailing comma to the line that needs it.
// The error may point to the line itself (e.g. a closing '}') or the line after.
func fixMissingComma(lines []string, errLine int) bool {
	// First, check if the error line itself needs the comma (e.g., "}" closing a map literal in an arg list)
	if errLine >= 0 && errLine < len(lines) {
		selfTrimmed := strings.TrimRight(lines[errLine], " \t\r")
		selfClean := strings.TrimSpace(selfTrimmed)
		if selfClean == "}" {
			lines[errLine] = selfTrimmed + ","
			return true
		}
	}

	// Otherwise, the missing comma is on the previous non-empty line
	target := -1
	for i := errLine - 1; i >= 0; i-- {
		trimmed := strings.TrimSpace(lines[i])
		if trimmed == "" || strings.HasPrefix(trimmed, "//") {
			continue
		}
		target = i
		break
	}
	if target < 0 {
		return false
	}

	line := lines[target]
	trimmed := strings.TrimRight(line, " \t\r")

	// Don't add comma if line already ends with comma, opening brace/paren, or is a comment
	if trimmed == "" {
		return false
	}
	lastChar := trimmed[len(trimmed)-1]
	if lastChar == ',' || lastChar == '{' || lastChar == '(' || lastChar == '[' {
		return false
	}
	// Don't add comma after comment lines
	if strings.HasPrefix(strings.TrimSpace(trimmed), "//") {
		return false
	}

	lines[target] = trimmed + ","
	return true
}

// rxOnlyCloseParen matches lines that are only whitespace + ")" or "),"
//
//nolint:unused // retained for upcoming close-paren autofix variant
var rxOnlyCloseParen = regexp.MustCompile(`^\s*\)\s*,?\s*$`)

// fixUnexpectedCloseParen handles "expected statement, found ')'"
// Common cause: extra ')' inside a closure, or a }) that should be just }
func fixUnexpectedCloseParen(lines []string, errLine int) bool {
	if errLine < 0 || errLine >= len(lines) {
		return false
	}
	trimmed := strings.TrimSpace(lines[errLine])

	// Case 1: Line is just ")" — likely extra paren, remove it
	if trimmed == ")" {
		// Check if previous non-empty line ends with "}" — this is a "})" split across lines
		prev := findPrevNonEmpty(lines, errLine)
		if prev >= 0 && strings.TrimSpace(lines[prev]) == "}" {
			// Merge: change prev to "})" and remove current line
			indent := leadingWhitespace(lines[prev])
			lines[prev] = indent + "})"
			lines = removeLineInPlace(lines, errLine)
			return true
		}
		// Otherwise just remove the stray ")"
		lines[errLine] = ""
		return true
	}

	// Case 2: Line is "})" but parser says ')' is unexpected as a statement.
	// This means '}' closed a struct/map literal (not a func body), so ')' is wrong.
	// Fix: remove the ')' → change "})" to "}"
	if trimmed == "})" {
		indent := leadingWhitespace(lines[errLine])
		lines[errLine] = indent + "}"
		return true
	}

	// Case 3: Line has "})}" or similar — normalize safeTest closure
	if trimmed == "})}" {
		indent := leadingWhitespace(lines[errLine])
		lines[errLine] = indent + "})"
		return true
	}

	return false
}

// fixUnexpectedCloseParenTopLevel handles "expected declaration, found ')'"
// Usually a stray ')' at top level from a mismatched safeTest closure.
func fixUnexpectedCloseParenTopLevel(lines []string, errLine int) bool {
	if errLine < 0 || errLine >= len(lines) {
		return false
	}
	trimmed := strings.TrimSpace(lines[errLine])

	// If the line is just ")" at top level, remove it
	if trimmed == ")" {
		lines[errLine] = ""
		return true
	}

	// If the line is "})" at top level, it's a safeTest closure that ended up
	// at package scope due to earlier mismatches. Check if previous non-empty
	// line is "}" — if so, merge into "})" on prev and blank this line.
	if trimmed == "})" {
		prev := findPrevNonEmpty(lines, errLine)
		if prev >= 0 {
			prevTrimmed := strings.TrimSpace(lines[prev])
			if prevTrimmed == "}" {
				indent := leadingWhitespace(lines[prev])
				lines[prev] = indent + "})"
				lines[errLine] = ""
				return true
			}
		}
		// Otherwise the '}' closed something, ')' is stray — keep just '}'
		indent := leadingWhitespace(lines[errLine])
		lines[errLine] = indent + "}"
		return true
	}

	return false
}

// fixMissingCloseBrace appends a closing "}" if the file ends without one.
func fixMissingCloseBrace(lines []string) bool {
	// Find last non-empty line
	for i := len(lines) - 1; i >= 0; i-- {
		trimmed := strings.TrimSpace(lines[i])
		if trimmed == "" {
			continue
		}
		if trimmed != "}" && trimmed != "})" {
			lines = append(lines, "}")
			return true
		}
		return false
	}
	return false
}

// fixExpectedOneExpression handles "expected 1 expression" errors.
// Common causes:
//   - Bare "return" with trailing comma: "return a,"  → remove trailing comma
//   - Multi-value return where only 1 expected: "return a, b" → keep first value
//   - Stray comma in single-expression context: "x," → remove comma
func fixExpectedOneExpression(lines []string, errLine int) bool {
	if errLine < 0 || errLine >= len(lines) {
		return false
	}

	line := lines[errLine]
	trimmed := strings.TrimSpace(line)

	// Case 1: "return something," — trailing comma after single return value
	// e.g. "return nil," or "return err,"
	if rxReturnTrailingComma.MatchString(trimmed) {
		// Remove the trailing comma
		indent := leadingWhitespace(line)
		cleaned := strings.TrimRight(trimmed, " \t")
		cleaned = strings.TrimRight(cleaned, ",")
		lines[errLine] = indent + cleaned
		return true
	}

	// Case 2: Line is a bare expression ending with comma (not a return)
	// e.g. inside a composite literal or call where "x," is unexpected
	if rxExprTrailingComma.MatchString(trimmed) && !strings.HasPrefix(trimmed, "return") {
		indent := leadingWhitespace(line)
		cleaned := strings.TrimRight(trimmed, " \t")
		cleaned = strings.TrimRight(cleaned, ",")
		lines[errLine] = indent + cleaned
		return true
	}

	// Case 3: "return a, b" where function expects 1 return value
	// The error column often points to the comma position.
	// We remove everything from the first comma to end of expression.
	if rxReturnMultiValue.MatchString(trimmed) {
		indent := leadingWhitespace(line)
		loc := rxReturnMultiValue.FindStringSubmatchIndex(trimmed)
		if loc != nil {
			// group 1 = the part before the comma
			firstVal := trimmed[loc[2]:loc[3]]
			lines[errLine] = indent + "return " + firstVal
			return true
		}
	}

	return false
}

// rxReturnTrailingComma matches "return <expr>," with a trailing comma
var rxReturnTrailingComma = regexp.MustCompile(`^return\s+.+,\s*$`)

// rxExprTrailingComma matches any expression ending with a trailing comma
var rxExprTrailingComma = regexp.MustCompile(`^[^,]+,\s*$`)

// rxReturnMultiValue matches "return <expr1>, <expr2>" (2+ values)
var rxReturnMultiValue = regexp.MustCompile(`^return\s+(\S+)\s*,\s*.+$`)

// fixExpectedOperand handles "expected operand, found <token>" errors.
// Common causes:
//   - Double commas: "a,, b" → "a, b"
//   - Trailing operator: "a +" on a line → remove the dangling operator
//   - Empty argument slot: "func(a, , b)" → "func(a, b)"
//   - Stray token like '}' or ')' where an expression is expected
func fixExpectedOperand(lines []string, errLine int, msg string) bool {
	if errLine < 0 || errLine >= len(lines) {
		return false
	}

	line := lines[errLine]
	trimmed := strings.TrimSpace(line)
	indent := leadingWhitespace(line)

	// Extract the unexpected token from the error message
	// Format: "expected operand, found '<token>'"
	foundToken := ""
	if idx := strings.Index(msg, "found '"); idx >= 0 {
		rest := msg[idx+7:]
		if end := strings.Index(rest, "'"); end >= 0 {
			foundToken = rest[:end]
		}
	}

	// Case 1: Double commas anywhere in the line: ",," → ","
	if strings.Contains(line, ",,") {
		for strings.Contains(line, ",,") {
			line = strings.ReplaceAll(line, ",,", ",")
		}
		lines[errLine] = line
		return true
	}

	// Case 2: Empty argument slot: "(a, , b)" or "(, b)" — remove empty slot
	if rxEmptyArgSlot.MatchString(trimmed) {
		cleaned := rxEmptyArgSlot.ReplaceAllString(trimmed, "$1$2")
		// Clean up leading comma after open paren: "(, " → "("
		cleaned = rxLeadingComma.ReplaceAllString(cleaned, "$1")
		lines[errLine] = indent + cleaned
		return true
	}

	// Case 3: Line ends with a dangling binary operator (+, -, *, /, |, &, etc.)
	if rxDanglingOperator.MatchString(trimmed) {
		// Check if next non-empty line starts with an operand — merge them
		next := findNextNonEmpty(lines, errLine)
		if next >= 0 {
			nextTrimmed := strings.TrimSpace(lines[next])
			// Merge: keep the operator on this line but append next line's content
			lines[errLine] = strings.TrimRight(line, " \t\r") + " " + nextTrimmed
			lines[next] = ""
			return true
		}
		// No next line to merge — remove the dangling operator
		cleaned := rxDanglingOperator.ReplaceAllString(trimmed, "")
		lines[errLine] = indent + cleaned
		return true
	}

	// Case 4: Found '}' or ')' where operand expected — likely a missing argument
	// before a closure end. Remove the offending line if it's just the token.
	if foundToken == "}" || foundToken == ")" {
		if trimmed == foundToken {
			// Check if previous line could absorb this (e.g., split "})") 
			prev := findPrevNonEmpty(lines, errLine)
			if prev >= 0 && foundToken == ")" && strings.TrimSpace(lines[prev]) == "}" {
				prevIndent := leadingWhitespace(lines[prev])
				lines[prev] = prevIndent + "})"
				lines[errLine] = ""
				return true
			}
		}
	}

	// Case 5: Found 'newline' — the line before is incomplete (missing operand after operator)
	if foundToken == "newline" {
		// Check if previous line ends with an operator
		prev := findPrevNonEmpty(lines, errLine)
		if prev >= 0 && rxDanglingOperator.MatchString(strings.TrimSpace(lines[prev])) {
			// Merge current line onto previous
			lines[prev] = strings.TrimRight(lines[prev], " \t\r") + " " + trimmed
			lines[errLine] = ""
			return true
		}
	}

	return false
}

// fixSemicolonExpectedCommaFound handles "expected ';', found ','".
// This occurs when '}' closes a code block (for/if/func body) but has a stray
// trailing comma, making the parser think it's a composite literal element.
// Fix: remove the trailing comma.
func fixSemicolonExpectedCommaFound(lines []string, errLine int) bool {
	if errLine < 0 || errLine >= len(lines) {
		return false
	}
	trimmed := strings.TrimSpace(lines[errLine])

	// Case 1: Line is "}," — remove the comma
	if trimmed == "}," {
		indent := leadingWhitespace(lines[errLine])
		lines[errLine] = indent + "}"
		return true
	}

	// Case 2: Line ends with "}," as part of a longer expression (e.g., "},")
	if strings.HasSuffix(trimmed, "},") {
		// Only fix if the line is JUST closing braces (possibly nested)
		// e.g., "},", "}},", etc. — not "foo},"
		stripped := strings.TrimRight(trimmed, ",")
		allBraces := true
		for _, c := range stripped {
			if c != '}' {
				allBraces = false
				break
			}
		}
		if allBraces {
			indent := leadingWhitespace(lines[errLine])
			lines[errLine] = indent + stripped
			return true
		}
	}

	// Case 3: Line has a trailing comma after any expression where ';' expected
	// e.g., "panicked = true," inside an if block
	line := lines[errLine]
	trimmedRight := strings.TrimRight(line, " \t\r")
	if len(trimmedRight) > 0 && trimmedRight[len(trimmedRight)-1] == ',' {
		lines[errLine] = trimmedRight[:len(trimmedRight)-1]
		return true
	}

	return false
}

// rxStrayCommaStatement matches statement lines that end with a stray comma.
// Covers:  x := expr,   |   var x Type,   |   x = expr,
var rxStrayCommaStatement = regexp.MustCompile(
	`^\s+` + // must be indented (inside a func body)
		`(?:` +
		`\w+\s*:=\s*.+` + // short var decl: x := expr
		`|` +
		`var\s+\w+\s+.+` + // var decl: var x Type
		`|` +
		`\w+\s*=\s*.+` + // assignment: x = expr
		`)` +
		`,\s*$`, // trailing comma
)

// fixStrayCommaStatements is a pre-pass that removes trailing commas from
// statement lines (`:=`, `var`, `=`) inside function bodies. These don't
// produce a unique parser error on the offending line — instead they cause
// cascading "missing ','" or "expected operand" errors on subsequent lines.
func fixStrayCommaStatements(lines []string, path string) int {
	fixes := 0
	for i, line := range lines {
		if !rxStrayCommaStatement.MatchString(line) {
			continue
		}
		trimmed := strings.TrimSpace(line)
		// Skip lines that are legitimately inside composite literals.
		// Heuristic: if the previous non-empty, non-comment line ends with '{',
		// we're likely inside a struct/map literal — don't touch it.
		prev := findPrevNonEmpty(lines, i)
		if prev >= 0 {
			prevTrimmed := strings.TrimSpace(lines[prev])
			if strings.HasSuffix(prevTrimmed, "{") {
				continue
			}
		}
		// Don't fix if line contains ":=" inside a struct literal value
		// (e.g., a map key with :=). Extra safety: skip if line starts with `"`
		if strings.HasPrefix(trimmed, "\"") || strings.HasPrefix(trimmed, "`") {
			continue
		}
		// Remove the trailing comma
		cleaned := strings.TrimRight(line, " \t\r")
		cleaned = cleaned[:len(cleaned)-1]
		lines[i] = cleaned
		fixes++
		addRecord(path, i+1, "stray-comma-on-statement", "removed trailing comma from statement line")
	}
	return fixes
}

// rxEmptyArgSlot matches ", ," or "(," patterns (empty argument slots)
var rxEmptyArgSlot = regexp.MustCompile(`(,)\s*,(\s*)`)

// rxLeadingComma matches "(, " at start of arg list
var rxLeadingComma = regexp.MustCompile(`(\()\s*,\s*`)

// rxDanglingOperator matches lines ending with a binary operator
var rxDanglingOperator = regexp.MustCompile(`[+\-*/|&^%<>]=?\s*$`)

func findNextNonEmpty(lines []string, from int) int {
	for i := from + 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) != "" {
			return i
		}
	}
	return -1
}

// --- helpers ---

func findPrevNonEmpty(lines []string, from int) int {
	for i := from - 1; i >= 0; i-- {
		if strings.TrimSpace(lines[i]) != "" {
			return i
		}
	}
	return -1
}

func leadingWhitespace(s string) string {
	for i, c := range s {
		if c != ' ' && c != '\t' {
			return s[:i]
		}
	}
	return s
}

func removeLineInPlace(lines []string, idx int) []string {
	// We can't change the slice header from the caller, so blank the line instead
	// This is simpler and avoids line-number drift in the same pass
	lines[idx] = ""
	return lines
}

// writeReport generates data/coverage/syntax-issues.txt with all fixes.
func writeReport(totalFixed, totalFiles, totalScanned int) {
	reportDir := "data/coverage"
	if err := os.MkdirAll(reportDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "  ⚠ Cannot create report dir %s: %v\n", reportDir, err)
		return
	}

	reportPath := filepath.Join(reportDir, "syntax-issues.txt")
	var b strings.Builder

	ts := time.Now().Format("2006-01-02 15:04:05")
	mode := "applied"
	if dryRun {
		mode = "dry-run (no files modified)"
	}

	b.WriteString("================================================================================\n")
	b.WriteString("  Syntax Issues Report — " + ts + "\n")
	b.WriteString("  Mode: " + mode + "\n")
	b.WriteString("================================================================================\n\n")

	b.WriteString(fmt.Sprintf("  Files scanned:    %d\n", totalScanned))
	b.WriteString(fmt.Sprintf("  Files with fixes: %d\n", totalFiles))
	b.WriteString(fmt.Sprintf("  Total fixes:      %d\n\n", totalFixed))

	// ── Embedded rules reference ──
	b.WriteString(rulesReference)

	if len(allRecords) == 0 {
		b.WriteString("  ✓ No syntax issues found.\n")
	} else {
		// Summary by rule
		ruleCounts := make(map[string]int)
		for _, r := range allRecords {
			ruleCounts[r.Rule]++
		}
		b.WriteString("────────────────────────────────────────────────────────────────────────────────\n")
		b.WriteString(" SUMMARY BY RULE\n")
		b.WriteString("────────────────────────────────────────────────────────────────────────────────\n\n")
		for rule, count := range ruleCounts {
			b.WriteString(fmt.Sprintf("  %-35s %d\n", rule, count))
		}

		// Per-file details
		b.WriteString("\n────────────────────────────────────────────────────────────────────────────────\n")
		b.WriteString(" DETAILS\n")
		b.WriteString("────────────────────────────────────────────────────────────────────────────────\n\n")

		currentFile := ""
		for _, r := range allRecords {
			if r.File != currentFile {
				currentFile = r.File
				b.WriteString(fmt.Sprintf("  %s\n", currentFile))
			}
			b.WriteString(fmt.Sprintf("    Line %-5d [%-30s] %s\n", r.Line, r.Rule, r.Message))
		}
	}

	b.WriteString("\n================================================================================\n")

	if err := os.WriteFile(reportPath, []byte(b.String()), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "  ⚠ Failed to write report: %v\n", err)
		return
	}
	fmt.Printf("  Report → %s\n", reportPath)
}
