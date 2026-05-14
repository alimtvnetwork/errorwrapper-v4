#!/usr/bin/env bash
# Self-test for check-safetest-boundaries brace-balance logic (Check 7)
# using synthetic Go fixtures with tricky syntax.
set -euo pipefail

FIXTURES_DIR="/tmp/checker-fixtures"
rm -rf "$FIXTURES_DIR"
mkdir -p "$FIXTURES_DIR"

PASS=0
FAIL=0

run_fixture() {
    local name="$1"
    local expect="$2"  # "pass" or "fail"
    local file="$FIXTURES_DIR/${name}_test.go"

    # Run only the brace-balance check (Check 7) via inline Python
    result=$(python3 - "$file" <<'PY' 2>&1 || true
import sys
from pathlib import Path

path = Path(sys.argv[1])
lines = path.read_text(encoding="utf-8").splitlines()
content = '\n'.join(lines)

depth = 0
in_str = False; in_raw = False; in_lc = False; in_rune = False
prev_ch = '\0'

ci = 0
while ci < len(content):
    ch = content[ci]
    if ch == '\n':
        in_lc = False; prev_ch = ch; ci += 1; continue
    if in_lc:
        prev_ch = ch; ci += 1; continue
    if in_str:
        if ch == '"' and prev_ch != '\\': in_str = False
        prev_ch = ch; ci += 1; continue
    if in_raw:
        if ch == '`': in_raw = False
        prev_ch = ch; ci += 1; continue
    if in_rune:
        if ch == "'" and prev_ch != '\\': in_rune = False
        prev_ch = ch; ci += 1; continue
    if ch == '/' and prev_ch == '/':
        in_lc = True; prev_ch = ch; ci += 1; continue
    if ch == '"':
        in_str = True; prev_ch = ch; ci += 1; continue
    if ch == '`':
        in_raw = True; prev_ch = ch; ci += 1; continue
    if ch == "'":
        in_rune = True; prev_ch = ch; ci += 1; continue
    if ch == '{': depth += 1
    elif ch == '}': depth -= 1
    prev_ch = ch; ci += 1

if depth != 0:
    print(f"UNBALANCED depth={depth}")
    sys.exit(1)
print("BALANCED")
sys.exit(0)
PY
    )

    if [ "$expect" = "pass" ]; then
        if echo "$result" | grep -q "BALANCED"; then
            PASS=$((PASS + 1))
        else
            echo "FAIL: $name — expected pass but got: $result"
            FAIL=$((FAIL + 1))
        fi
    else
        if echo "$result" | grep -q "UNBALANCED"; then
            PASS=$((PASS + 1))
        else
            echo "FAIL: $name — expected fail but got: $result"
            FAIL=$((FAIL + 1))
        fi
    fi
}

# --- Fixture: braces inside double-quoted strings ---
cat > "$FIXTURES_DIR/strings_test.go" << 'GO'
package foo
func TestStrings() {
	x := "{ not a real brace }"
	y := "nested { { { braces"
}
GO
run_fixture "strings" "pass"

# --- Fixture: braces inside backtick (raw) strings ---
cat > "$FIXTURES_DIR/rawstrings_test.go" << 'GO'
package foo
func TestRaw() {
	x := `raw { brace
	still raw }}}
	`
}
GO
run_fixture "rawstrings" "pass"

# --- Fixture: braces inside single-line comments ---
cat > "$FIXTURES_DIR/comments_test.go" << 'GO'
package foo
func TestComments() {
	// this { should be ignored
	// so should this }}}
	x := 1
}
GO
run_fixture "comments" "pass"

# --- Fixture: rune literals with braces and quotes ---
cat > "$FIXTURES_DIR/runes_test.go" << 'GO'
package foo
func TestRunes() {
	a := '{'
	b := '}'
	c := '"'
	d := '\''
}
GO
run_fixture "runes" "pass"

# --- Fixture: escaped quote inside string ---
cat > "$FIXTURES_DIR/escaped_test.go" << 'GO'
package foo
func TestEscaped() {
	x := "escaped \" quote { inside"
	y := "another \\ backslash"
}
GO
run_fixture "escaped" "pass"

# --- Fixture: mixed edge cases ---
cat > "$FIXTURES_DIR/mixed_test.go" << 'GO'
package foo
func TestMixed() {
	a := '{'
	b := "}"
	// {{{
	c := `}{}{`
	if true {
		d := '"'
	}
}
GO
run_fixture "mixed" "pass"

# --- Fixture: actually unbalanced (should fail) ---
cat > "$FIXTURES_DIR/unbalanced_test.go" << 'GO'
package foo
func TestUnbalanced() {
	if true {
		x := 1
}
GO
run_fixture "unbalanced" "fail"

# --- Fixture: unbalanced hidden by string (should fail) ---
cat > "$FIXTURES_DIR/sneaky_test.go" << 'GO'
package foo
func TestSneaky() {
	x := "this string doesn't help"
	if true {
}
GO
run_fixture "sneaky" "fail"

echo ""
echo "Results: $PASS passed, $FAIL failed out of $((PASS + FAIL)) fixtures"

rm -rf "$FIXTURES_DIR"

if [ "$FAIL" -ne 0 ]; then
    exit 1
fi
echo "✓ All checker fixtures passed."
