#!/usr/bin/env bash
set -euo pipefail

DIR="tests/integratedtests/corestrtests"

if [ ! -d "$DIR" ]; then
  echo "⚠ Directory $DIR not found, skipping check."
  exit 0
fi

python3 - "$DIR" <<'PY'
import re
import sys
from pathlib import Path

root = Path(sys.argv[1])
issues: list[str] = []

control_re = re.compile(r'^\t\t(?:if|for|switch|select|else)\b')
func_open_re = re.compile(r'^\t\t[^\t].*\(func\(.*\{\s*$')
block_open_re = re.compile(r'^\t\t[^\t].*\{\s*$')
if_re = re.compile(r'^(\t+)if\b.*\{\s*$')
assign_open_re = re.compile(r'^\s*(?:var\s+\w+\s*=\s*func\s*\(|\w+\s*(?::=|=)\s*func\s*\()')

for path in sorted(root.glob("*_test.go")):
    lines = path.read_text(encoding="utf-8").splitlines()
    rel = str(path).replace("\\", "/")

    # Check 1: malformed safeTest boundary (missing inner `}` before `\t})`)
    for i in range(1, len(lines)):
        curr = lines[i]
        if curr == "\t})" or re.match(r"^\t\)\}$", curr):
            prev = lines[i - 1]
            if (
                re.match(r'^\t\t\t', prev)
                and not re.match(r'^\t\t\t\s*\}\s*$', prev)
                and not re.match(r'^\t\t\t\s*//', prev)
            ):
                issues.append(f"  {rel}:{i+1}: missing closing }} before %}})")

    # Check 2: closure arg `}` missing `)` (avoid false-positives on if/for/switch/select blocks)
    for i, line in enumerate(lines):
        if line.rstrip() != "\t\t}":
            continue

        depth = 0
        for j in range(i - 1, max(i - 40, -1), -1):
            l = lines[j].rstrip()
            if not l:
                continue

            if l == "\t\t})":
                depth += 1
                continue

            if l.startswith("func Test_") or l.startswith("\tsafeTest("):
                break

            if l == "\t\t}":
                if depth == 0:
                    break
                continue

            if block_open_re.search(l):
                if control_re.search(l):
                    break

                if func_open_re.search(l):
                    if depth > 0:
                        depth -= 1
                    else:
                        issues.append(f"  {rel}:{i+1}: closure }} missing )")
                break

    # Check 3: empty if blocks (comment-only body with no actual statement)
    for i, line in enumerate(lines):
        m = if_re.match(line)
        if not m:
            continue

        indent = m.group(1)
        close = indent + "}"
        has_stmt = False
        for j in range(i + 1, min(i + 20, len(lines))):
            body = lines[j]
            stripped = body.strip()
            if body.rstrip() == close or body.rstrip() == close + ")":
                break
            if stripped == "" or stripped.startswith("//"):
                continue
            has_stmt = True
            break

        if not has_stmt:
            issues.append(f"  {rel}:{i+1}: empty if block (no statements)")

    # Check 4: func assignment closure must end with `}` (not `})`)
    for i, line in enumerate(lines):
        if not assign_open_re.search(line):
            continue

        depth = 0
        started = False
        for j in range(i, min(i + 300, len(lines))):
            l = lines[j]
            open_count = l.count("{")
            close_count = l.count("}")
            if open_count > 0:
                started = True

            depth += open_count
            depth -= close_count

            if started and depth == 0:
                if lines[j].rstrip().endswith("})"):
                    issues.append(
                        f"  {rel}:{j+1}: func assignment closes with }}) (should close with }} only)"
                    )
                break

    # Check 5: placeholder `...` lines are forbidden
    for i, line in enumerate(lines):
        if line.strip() == "...":
            issues.append(f"  {rel}:{i+1}: placeholder line ... is not allowed")

    # Check 6: unclosed if/for blocks (missing closing } before }) or })
    # Detects: if cond { ... }) where the if's own } is absent
    if_for_re = re.compile(r'^(\t+)(?:if |for |select \{).*\{\s*$')
    for i, line in enumerate(lines):
        m = if_for_re.match(line)
        if not m:
            continue
        block_indent = m.group(1)
        block_tabs = len(block_indent)
        # Scan forward for the matching close
        found_close = False
        for j in range(i + 1, min(i + 50, len(lines))):
            inner = lines[j]
            inner_stripped = inner.strip()
            inner_tabs = len(inner) - len(inner.lstrip('\t'))
            if inner_stripped == '':
                continue
            # Proper close: } at same indent as the if/for
            if inner_stripped.startswith('}') and inner_tabs == block_tabs:
                found_close = True
                break
            # If we hit a line at LOWER indent (e.g. }), }) at parent level),
            # the if/for was never closed
            if inner_tabs < block_tabs and inner_stripped != '':
                issues.append(
                    f"  {rel}:{i+1}: unclosed if/for block (no matching }} before line {j+1})"
                )
                found_close = True  # reported, stop scanning
                break

    # Check 7: brace balance per file (robust scanner over full content)
    depth = 0
    in_str = False; in_raw = False; in_lc = False; in_rune = False
    prev_ch = '\0'
    content = '\n'.join(lines)

    ci = 0
    while ci < len(content):
        ch = content[ci]

        if ch == '\n':
            in_lc = False
            prev_ch = ch
            ci += 1
            continue

        if in_lc:
            prev_ch = ch
            ci += 1
            continue

        if in_str:
            if ch == '"' and prev_ch != '\\':
                in_str = False
            prev_ch = ch
            ci += 1
            continue

        if in_raw:
            if ch == '`':
                in_raw = False
            prev_ch = ch
            ci += 1
            continue

        if in_rune:
            if ch == "'" and prev_ch != '\\':
                in_rune = False
            prev_ch = ch
            ci += 1
            continue

        if ch == '/' and prev_ch == '/':
            in_lc = True
            prev_ch = ch
            ci += 1
            continue

        if ch == '"':
            in_str = True
            prev_ch = ch
            ci += 1
            continue

        if ch == '`':
            in_raw = True
            prev_ch = ch
            ci += 1
            continue

        if ch == "'":
            in_rune = True
            prev_ch = ch
            ci += 1
            continue

        if ch == '{':
            depth += 1
        elif ch == '}':
            depth -= 1

        prev_ch = ch
        ci += 1

    if depth != 0:
        issues.append(f"  {rel}: unbalanced braces (depth={depth})")

if issues:
    print("\n".join(issues))
    print("")
    print("✗ Lint failures detected.")
    sys.exit(1)

print("✓ All safeTest boundaries, if blocks, assignment closures, and placeholder lines are clean.")
sys.exit(0)
PY