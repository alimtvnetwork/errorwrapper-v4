#!/usr/bin/env python3
"""
AAA Comments Auto-Inserter
===========================
Adds missing // Arrange, // Act, // Assert comments to Go test functions.

Analyzes each test function body to detect section boundaries:
- Arrange: variable declarations, setup, input extraction
- Act: function-under-test calls, result capture (actual := args.Map{...})
- Assert: ShouldBeEqual, ShouldBeEqualMap, ShouldBeSafe calls

Usage:
  python3 aaa_comments.py --dry-run                     # preview all
  python3 aaa_comments.py --package ostypetests          # fix one package
  python3 aaa_comments.py                                # fix everything
"""

import re
import os
import sys
import argparse
from dataclasses import dataclass, field

TESTS_ROOT = "tests/integratedtests"

@dataclass
class Stats:
    files_scanned: int = 0
    files_modified: int = 0
    functions_processed: int = 0
    comments_added: int = 0
    functions_skipped: int = 0


def has_aaa_comments(lines: list[str]) -> bool:
    """Check if a function body already has AAA comments."""
    text = '\n'.join(lines)
    return '// Arrange' in text or '// Act' in text or '// Assert' in text


def is_arrange_line(line: str) -> bool:
    """Lines that belong to Arrange section."""
    s = line.strip()
    if not s or s.startswith('//'):
        return True  # blank/comment lines are neutral
    # Variable declarations and setup
    if s.startswith('var ') or s.startswith('const '):
        return True
    # Short variable declaration (not actual/expected)
    if ':=' in s and not s.startswith('actual') and not s.startswith('expected'):
        # But not if it's creating args.Map for actual
        if 'args.Map' not in s or 'actual' not in s.split(':=')[0]:
            return True
    # Input extraction from test case
    if 'testCase.' in s or 'tc.' in s or 'input[' in s or 'input.' in s:
        return True
    # Type assertions and setup
    if '.(' in s and ':=' in s and 'actual' not in s:
        return True
    return False


def is_act_line(line: str) -> bool:
    """Lines that belong to Act section."""
    s = line.strip()
    if not s:
        return False
    # actual := args.Map{...}
    if s.startswith('actual') and ':=' in s:
        return True
    # Function call that produces result
    if ':=' in s and 'actual' not in s and 'expected' not in s:
        # Could be the function-under-test call
        return True
    return False


def is_assert_line(line: str) -> bool:
    """Lines that belong to Assert section."""
    s = line.strip()
    if not s:
        return False
    if s.startswith('expected') and (':=' in s or '.ShouldBeEqual' in s):
        return True
    if '.ShouldBeEqual' in s or '.ShouldBeEqualMap' in s or '.ShouldBeSafe' in s:
        return True
    if 'testCase.ShouldBeEqual' in s or 'tc.ShouldBeEqual' in s:
        return True
    return False


def find_section_boundaries(body_lines: list[str]) -> tuple:
    """
    Find where Arrange ends, Act starts/ends, Assert starts.
    Returns (arrange_end, act_start, assert_start) as line indices within body_lines.
    Returns None if sections can't be determined.
    """
    n = len(body_lines)
    if n == 0:
        return None

    # Find the first 'actual := args.Map' or 'actual :=' line
    act_start = -1
    for i, line in enumerate(body_lines):
        s = line.strip()
        if s.startswith('actual') and ':=' in s:
            act_start = i
            break

    # Find the first 'expected' line or ShouldBeEqual
    assert_start = -1
    for i, line in enumerate(body_lines):
        s = line.strip()
        if (s.startswith('expected') and ':=' in s) or '.ShouldBeEqual' in s or '.ShouldBeEqualMap' in s:
            assert_start = i
            break

    # For loop-based tests with testCase.ShouldBeEqualMap
    if assert_start == -1:
        for i, line in enumerate(body_lines):
            s = line.strip()
            if 'ShouldBeEqualMap' in s or 'ShouldBeSafe' in s:
                assert_start = i
                break

    return act_start, assert_start


def classify_function_type(body_lines: list[str]) -> str:
    """Classify test function type."""
    text = '\n'.join(body_lines)
    if 'for caseIndex' in text or 'for _, testCase' in text or 'for _, tc' in text:
        return 'loop'
    if 'safeTest(' in text:
        return 'safetest'
    return 'simple'


def insert_aaa_comments(body_lines: list[str], func_indent: str) -> list[str]:
    """Insert AAA comments into a function body. Returns modified lines."""
    if has_aaa_comments(body_lines):
        return body_lines

    func_type = classify_function_type(body_lines)

    # For safeTest wrapper, find the inner closure body
    if func_type == 'safetest':
        return insert_aaa_in_safetest(body_lines, func_indent)

    # For loop tests, insert inside the loop body
    if func_type == 'loop':
        return insert_aaa_in_loop(body_lines, func_indent)

    # Simple test: insert at top level
    return insert_aaa_simple(body_lines, func_indent)


def insert_aaa_simple(body_lines: list[str], base_indent: str) -> list[str]:
    """Insert AAA comments for simple (non-loop) test functions."""
    boundaries = find_section_boundaries(body_lines)
    if boundaries is None:
        return body_lines

    act_start, assert_start = boundaries

    if act_start == -1 and assert_start == -1:
        return body_lines

    result = []
    indent = base_indent + '\t'

    # Determine arrange_end: everything before act_start that isn't blank
    arrange_end = -1
    if act_start > 0:
        for i in range(act_start - 1, -1, -1):
            if body_lines[i].strip():
                arrange_end = i
                break

    # Insert comments
    arrange_inserted = False
    act_inserted = False
    assert_inserted = False

    for i, line in enumerate(body_lines):
        # Insert // Arrange before the first non-blank line if there's arrange content
        if not arrange_inserted and i == 0 and act_start > 0:
            # Check if first line is non-blank
            first_content = -1
            for j in range(len(body_lines)):
                if body_lines[j].strip() and not body_lines[j].strip().startswith('//'):
                    first_content = j
                    break
            if first_content >= 0 and first_content < act_start:
                result.append(f'{indent}// Arrange')
                arrange_inserted = True

        # Insert // Act before act_start
        if not act_inserted and i == act_start:
            if arrange_inserted:
                # Add blank line before Act if previous line isn't blank
                if result and result[-1].strip():
                    result.append('')
            result.append(f'{indent}// Act')
            act_inserted = True

        # Insert // Assert before assert_start
        if not assert_inserted and i == assert_start and assert_start != act_start:
            if result and result[-1].strip():
                result.append('')
            result.append(f'{indent}// Assert')
            assert_inserted = True

        result.append(line)

    # If no arrange section (act is at line 0), just add Act and Assert
    if not arrange_inserted and act_start == 0 and not act_inserted:
        result.insert(0, f'{indent}// Act')

    return result


def insert_aaa_in_loop(body_lines: list[str], base_indent: str) -> list[str]:
    """Insert AAA comments inside a for-loop body."""
    result = []
    loop_indent = base_indent + '\t\t'

    # Find the loop start
    in_loop = False
    loop_body_start = -1

    for i, line in enumerate(body_lines):
        s = line.strip()
        if ('for caseIndex' in s or 'for _, testCase' in s or 'for _, tc' in s) and '{' in s:
            in_loop = True
            result.append(line)
            loop_body_start = i + 1
            continue

        if in_loop and i == loop_body_start:
            # Analyze the loop body
            loop_end = find_matching_brace(body_lines, i - 1)
            if loop_end == -1:
                loop_end = len(body_lines)

            loop_body = body_lines[i:loop_end]
            if not has_aaa_comments(loop_body):
                boundaries = find_section_boundaries(loop_body)
                if boundaries and (boundaries[0] != -1 or boundaries[1] != -1):
                    act_start, assert_start = boundaries
                    commented = insert_comments_at_boundaries(
                        loop_body, act_start, assert_start, loop_indent
                    )
                    result.extend(commented)
                    # Skip original loop body lines
                    for j in range(loop_end, len(body_lines)):
                        result.append(body_lines[j])
                    return result

            in_loop = False

        result.append(line)

    return result


def insert_aaa_in_safetest(body_lines: list[str], base_indent: str) -> list[str]:
    """Insert AAA comments inside a safeTest closure."""
    result = []
    inner_indent = base_indent + '\t\t'

    # Find safeTest( line and the closure body
    safe_start = -1
    for i, line in enumerate(body_lines):
        if 'safeTest(' in line.strip():
            safe_start = i
            break

    if safe_start == -1:
        return body_lines

    # Add lines up to and including safeTest line
    for i in range(safe_start + 1):
        result.append(body_lines[i])

    # Analyze closure body (everything between safeTest and closing })
    closure_body = body_lines[safe_start + 1:]
    # Remove trailing }) lines
    closing_tokens = {'', '})', '})'}
    while closure_body and closure_body[-1].strip() in closing_tokens:
        closure_body = closure_body[:-1]

    if not has_aaa_comments(closure_body):
        boundaries = find_section_boundaries(closure_body)
        if boundaries and (boundaries[0] != -1 or boundaries[1] != -1):
            act_start, assert_start = boundaries
            commented = insert_comments_at_boundaries(
                closure_body, act_start, assert_start, inner_indent
            )
            result.extend(commented)
            # Add back trailing lines
            for i in range(safe_start + 1 + len(closure_body), len(body_lines)):
                result.append(body_lines[i])
            return result

    # If we can't analyze, return original
    return body_lines


def insert_comments_at_boundaries(
    lines: list[str], act_start: int, assert_start: int, indent: str
) -> list[str]:
    """Insert AAA comments at detected boundaries."""
    result = []
    has_arrange = act_start > 0
    arrange_inserted = False
    act_inserted = False
    assert_inserted = False

    for i, line in enumerate(lines):
        # Arrange at beginning if there's content before act
        if not arrange_inserted and i == 0 and has_arrange:
            # Find first non-blank line
            for j in range(len(lines)):
                if lines[j].strip() and not lines[j].strip().startswith('//'):
                    if j < act_start:
                        result.append(f'{indent}// Arrange')
                        arrange_inserted = True
                    break

        # Act
        if not act_inserted and i == act_start:
            if result and result[-1].strip():
                result.append('')
            result.append(f'{indent}// Act')
            act_inserted = True

        # Assert
        if not assert_inserted and assert_start != -1 and i == assert_start and i != act_start:
            if result and result[-1].strip():
                result.append('')
            result.append(f'{indent}// Assert')
            assert_inserted = True

        result.append(line)

    return result


def find_matching_brace(lines: list[str], open_line: int) -> int:
    """Find line index of matching closing brace."""
    depth = 0
    for i in range(open_line, len(lines)):
        depth += lines[i].count('{') - lines[i].count('}')
        if depth <= 0:
            return i
    return -1


def extract_functions(content: str) -> list[tuple]:
    """
    Extract test function positions: (func_start, body_start, body_end, indent).
    """
    lines = content.split('\n')
    functions = []
    i = 0

    while i < len(lines):
        # Match func Test...(t *testing.T) { (with or without underscore)
        m = re.match(r'^(func Test\w*\(t \*testing\.T\)\s*\{)', lines[i])
        if m:
            func_start = i
            body_start = i + 1
            # Find matching closing brace
            depth = 1
            j = body_start
            while j < len(lines) and depth > 0:
                depth += lines[j].count('{') - lines[j].count('}')
                if depth <= 0:
                    break
                j += 1
            body_end = j  # line with closing }
            functions.append((func_start, body_start, body_end, ''))
            i = body_end + 1
        else:
            i += 1

    return functions


def process_file(filepath: str, dry_run: bool, stats: Stats) -> bool:
    """Process a single test file. Returns True if modified."""
    with open(filepath, 'r') as f:
        content = f.read()

    # Skip files with no test functions
    if 'func Test_' not in content:
        return False

    stats.files_scanned += 1
    lines = content.split('\n')
    functions = extract_functions(content)

    if not functions:
        return False

    modified = False
    # Process functions in reverse order to preserve line numbers
    for func_start, body_start, body_end, indent in reversed(functions):
        body_lines = lines[body_start:body_end]

        if has_aaa_comments(body_lines):
            continue

        stats.functions_processed += 1
        new_body = insert_aaa_comments(body_lines, indent)

        if new_body != body_lines:
            lines[body_start:body_end] = new_body
            modified = True
            added = sum(1 for l in new_body if l.strip() in [
                '// Arrange', '// Act', '// Assert',
                '// Arrange', '// Act', '// Assert',  # with tabs
            ]) + sum(1 for l in new_body if l.strip().startswith('// Arrange')
                     or l.strip().startswith('// Act')
                     or l.strip().startswith('// Assert'))
            # Count unique AAA comments added (not already in body_lines)
            old_comments = sum(1 for l in body_lines if '// Arrange' in l or '// Act' in l or '// Assert' in l)
            new_comments = sum(1 for l in new_body if '// Arrange' in l or '// Act' in l or '// Assert' in l)
            stats.comments_added += (new_comments - old_comments)
        else:
            stats.functions_skipped += 1

    if not modified:
        return False

    new_content = '\n'.join(lines)

    if dry_run:
        rel = os.path.relpath(filepath)
        print(f"  [DRY-RUN] Would modify: {rel}")
    else:
        with open(filepath, 'w') as f:
            f.write(new_content)

    stats.files_modified += 1
    return True


def main():
    parser = argparse.ArgumentParser(description="AAA Comments Auto-Inserter")
    parser.add_argument('--dry-run', action='store_true', help='Preview without writing')
    parser.add_argument('--package', type=str, help='Process only this package')
    parser.add_argument('--file', type=str, help='Process only this file')
    parser.add_argument('--root', type=str, default=TESTS_ROOT, help='Tests root directory')
    parsed = parser.parse_args()

    stats = Stats()

    if parsed.file:
        process_file(parsed.file, parsed.dry_run, stats)
    else:
        root = parsed.root
        for dirpath, dirnames, filenames in sorted(os.walk(root)):
            if parsed.package:
                rel = os.path.relpath(dirpath, root)
                parts = rel.split(os.sep)
                if parsed.package not in parts:
                    continue
            for fname in sorted(filenames):
                if not fname.endswith("_test.go"):
                    continue
                fpath = os.path.join(dirpath, fname)
                process_file(fpath, parsed.dry_run, stats)

    # Report
    print()
    print("=" * 60)
    print("AAA Comments Auto-Insert Report")
    print("=" * 60)
    print(f"Files scanned:           {stats.files_scanned}")
    print(f"Files modified:          {stats.files_modified}")
    print(f"Functions processed:     {stats.functions_processed}")
    print(f"Comments added:          {stats.comments_added}")
    print(f"Functions skipped:       {stats.functions_skipped}")
    print()

    if parsed.dry_run:
        print(">>> DRY RUN — no files were modified. Remove --dry-run to apply.")
    else:
        print(f">>> Done. {stats.files_modified} files updated.")


if __name__ == '__main__':
    main()
