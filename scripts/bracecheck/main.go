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

// Package main provides a Go-based brace balance and syntax pre-checker for test files.
// Usage: go run ./scripts/bracecheck/ [files or dirs...]
// If no args, defaults to tests/integratedtests/corestrtests/
package main

import (
	"fmt"
	"go/parser"
	"go/scanner"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// errorCategory classifies a parser error for the summary.
type errorCategory struct {
	Pattern     string
	Description string
	Suggestion  string
}

var categories = []errorCategory{
	{
		Pattern:     "missing ',' before newline",
		Description: "Missing trailing comma",
		Suggestion:  "Add a trailing comma after the last argument in a multi-line function call or composite literal.",
	},
	{
		Pattern:     "expected statement, found ')'",
		Description: "Unexpected closing paren",
		Suggestion:  "Extra ')' or empty block inside a closure. Check safeTest boundaries — correct syntax is '})' not '})'.",
	},
	{
		Pattern:     "expected declaration, found ')'",
		Description: "Unexpected ')' at top level",
		Suggestion:  "A ')' appeared where a func/var/type declaration was expected. Likely a mismatched safeTest '})' closure.",
	},
	{
		Pattern:     "expected 1 expression",
		Description: "Expression count mismatch",
		Suggestion:  "A return or assignment has the wrong number of values on one side.",
	},
	{
		Pattern:     "expected '}', found 'EOF'",
		Description: "Unclosed brace at EOF",
		Suggestion:  "A '{' block was never closed. Check for missing '}' in the last function or closure.",
	},
	{
		Pattern:     "expected ';', found ','",
		Description: "Stray comma after block brace",
		Suggestion:  "A '}' closing a code block (for/if/func) has a trailing ',' — remove it. This causes massive cascading errors.",
	},
}

func classifyError(msg string) *errorCategory {
	for i := range categories {
		if strings.Contains(msg, categories[i].Pattern) {
			return &categories[i]
		}
	}
	return nil
}

func main() {
	args := os.Args[1:]
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

	fset := token.NewFileSet()
	failed := 0
	totalErrors := 0
	categoryCounts := make(map[string]int)

	for _, f := range files {
		src, readErr := os.ReadFile(f)
		if readErr != nil {
			fmt.Fprintf(os.Stderr, "✗ %s: cannot read: %v\n", f, readErr)
			failed++
			continue
		}
		lines := strings.Split(string(src), "\n")

		_, err := parser.ParseFile(fset, f, src, parser.AllErrors)
		if err == nil {
			continue
		}

		failed++
		errList, ok := err.(scanner.ErrorList)
		if !ok {
			fmt.Fprintf(os.Stderr, "✗ %s:\n  %v\n", f, err)
			totalErrors++
			continue
		}

		fmt.Fprintf(os.Stderr, "✗ %s: %d error(s)\n", f, len(errList))
		totalErrors += len(errList)

		// Show up to 8 errors with context
		limit := len(errList)
		if limit > 8 {
			limit = 8
		}

		for i := 0; i < limit; i++ {
			e := errList[i]
			cat := classifyError(e.Msg)
			catLabel := "unknown"
			if cat != nil {
				catLabel = cat.Description
				categoryCounts[cat.Description]++
			} else {
				categoryCounts["other"]++
			}

			fmt.Fprintf(os.Stderr, "\n  [%s] Line %d, Col %d: %s\n",
				catLabel, e.Pos.Line, e.Pos.Column, e.Msg)

			// Print surrounding source lines
			lineIdx := e.Pos.Line - 1
			start := lineIdx - 2
			if start < 0 {
				start = 0
			}
			end := lineIdx + 3
			if end > len(lines) {
				end = len(lines)
			}
			for li := start; li < end; li++ {
				marker := "  "
				if li == lineIdx {
					marker = ">>"
				}
				fmt.Fprintf(os.Stderr, "    %s %4d | %s\n", marker, li+1, lines[li])
			}

			// Show pointer to column
			if lineIdx >= 0 && lineIdx < len(lines) {
				col := e.Pos.Column - 1
				if col < 0 {
					col = 0
				}
				padding := strings.Repeat(" ", col)
				fmt.Fprintf(os.Stderr, "    %s %4s   %s^\n", "  ", "", padding)
			}

			if cat != nil {
				fmt.Fprintf(os.Stderr, "    💡 %s\n", cat.Suggestion)
			}
		}

		if len(errList) > limit {
			fmt.Fprintf(os.Stderr, "\n  ... and %d more error(s)\n", len(errList)-limit)
		}
		fmt.Fprintln(os.Stderr)
	}

	if failed > 0 {
		fmt.Fprintf(os.Stderr, "─── Summary ───\n")
		fmt.Fprintf(os.Stderr, "✗ %d file(s), %d total error(s)\n", failed, totalErrors)
		if len(categoryCounts) > 0 {
			fmt.Fprintf(os.Stderr, "\nBreakdown by error type:\n")
			for cat, count := range categoryCounts {
				fmt.Fprintf(os.Stderr, "  • %-30s %d\n", cat, count)
			}
		}
		fmt.Fprintln(os.Stderr)
		os.Exit(1)
	}
	fmt.Printf("✓ %d file(s) parsed OK\n", len(files))
}
