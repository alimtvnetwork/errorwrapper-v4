#!/usr/bin/env python3
"""
AAA Format Auto-Transformer
============================
Transforms Go test files from t.Error/t.Errorf patterns into
args.Map + ShouldBeEqual AAA format.

Handles these patterns:
  Pattern 1: if !expr { t.Error("msg") }
  Pattern 2: if expr { t.Error("msg") }
  Pattern 3: if val != expected { t.Errorf("fmt", args...) }
  Pattern 4: if val == expected { t.Errorf("fmt", args...) }

Usage:
  python3 aaa_transform.py [--dry-run] [--package <name>] [--file <path>]
  python3 aaa_transform.py --dry-run                      # preview all changes
  python3 aaa_transform.py --package ostypetests           # fix one package
  python3 aaa_transform.py --file tests/.../Coverage.go    # fix one file
  python3 aaa_transform.py                                 # fix everything
"""

import re
import os
import sys
import argparse
from dataclasses import dataclass, field
from typing import Optional

TESTS_ROOT = "tests/integratedtests"
ARGS_IMPORT = '"github.com/alimtvnetwork/core-v8/coretests/args"'

@dataclass
class TransformStats:
    files_processed: int = 0
    files_modified: int = 0
    patterns_transformed: int = 0
    patterns_skipped: int = 0
    skipped_details: list = field(default_factory=list)


def ensure_args_import(content: str) -> str:
    """Add args import if missing."""
    if ARGS_IMPORT in content:
        return content

    # Find import block
    # Pattern: import ( ... )
    import_block = re.search(r'(import\s*\()(.*?)(\))', content, re.DOTALL)
    if import_block:
        imports_text = import_block.group(2)
        # Add args import after the last import line
        new_imports = imports_text.rstrip() + f'\n\t{ARGS_IMPORT}\n'
        content = content[:import_block.start(2)] + new_imports + content[import_block.end(2):]
        return content

    # Single import: import "testing"
    single_import = re.search(r'import\s+"testing"', content)
    if single_import:
        replacement = f'import (\n\t"testing"\n\n\t{ARGS_IMPORT}\n)'
        content = content[:single_import.start()] + replacement + content[single_import.end():]
        return content

    return content


def extract_description(error_msg: str) -> str:
    """Clean up error message for use as ShouldBeEqual description."""
    # Remove format verbs and args
    msg = error_msg.strip().strip('"')
    # Remove trailing format args like %d, %s, %v, %c
    msg = re.sub(r'\s*%[dsvfcqxXtTp]', '', msg)
    # Remove "got" suffixes
    msg = re.sub(r'\s+got\s*$', '', msg)
    # Clean up
    msg = msg.strip().rstrip(',').strip()
    if not msg:
        msg = "assertion"
    return msg


def transform_simple_bool_if(lines: list, i: int, indent: str, stats: TransformStats) -> Optional[tuple]:
    """
    Transform:
      if !expr { t.Error("msg") }      → result=expr, expected=true
      if expr { t.Error("msg") }       → result=expr, expected=false
      
    Multi-line:
      if !expr {
          t.Error("msg")
      }
    """
    line = lines[i].rstrip()
    stripped = line.strip()

    # --- Single-line: if !expr { t.Fatal() } (no message) ---
    m_noarg = re.match(
        r'^(\s*)if\s+(!?)(.+?)\s*\{\s*t\.(Fatal|Error)\(\)\s*\}',
        line,
    )
    if m_noarg:
        ws, neg, expr, _ = m_noarg.groups()
        expected_val = "true" if neg == "!" else "false"
        expr = expr.strip()
        new_lines = [
            f'{ws}actual := args.Map{{"result": {expr}}}',
            f'{ws}expected := args.Map{{"result": {expected_val}}}',
            f'{ws}expected.ShouldBeEqual(t, 0, "assertion", actual)',
        ]
        stats.patterns_transformed += 1
        return (i, i, new_lines)

    # --- Single-line: if !expr { t.Error("msg") } ---
    m = re.match(
        r'^(\s*)if\s+(!?)(.+?)\s*\{\s*t\.(Error|Errorf|Fatal|Fatalf)\((".*?")\s*\)\s*\}',
        line,
    )
    if m:
        ws, neg, expr, _, msg = m.groups()
        expected_val = "true" if neg == "!" else "false"
        expr = expr.strip()
        desc = extract_description(msg)
        
        new_lines = [
            f'{ws}actual := args.Map{{"result": {expr}}}',
            f'{ws}expected := args.Map{{"result": {expected_val}}}',
            f'{ws}expected.ShouldBeEqual(t, 0, "{desc}", actual)',
        ]
        stats.patterns_transformed += 1
        return (i, i, new_lines)

    # --- Multi-line: if !expr { \n t.Error("msg") \n } ---
    m = re.match(r'^(\s*)if\s+(!?)(.+?)\s*\{\s*$', line)
    if m and i + 2 < len(lines):
        ws, neg, expr = m.groups()
        next_line = lines[i + 1].strip()
        close_line = lines[i + 2].strip() if i + 2 < len(lines) else ""

        # Check for t.Fatal() / t.Error() with no args
        noarg_match = re.match(r't\.(Error|Fatal)\(\)\s*$', next_line)
        if noarg_match and close_line == "}":
            expected_val = "true" if neg == "!" else "false"
            expr = expr.strip()
            new_lines = [
                f'{ws}actual := args.Map{{"result": {expr}}}',
                f'{ws}expected := args.Map{{"result": {expected_val}}}',
                f'{ws}expected.ShouldBeEqual(t, 0, "assertion", actual)',
            ]
            stats.patterns_transformed += 1
            return (i, i + 2, new_lines)

        # Check for t.Error("msg") on next line and } on line after
        err_match = re.match(
            r't\.(Error|Errorf|Fatal|Fatalf)\((".*?")\s*\)',
            next_line,
        )
        if err_match and close_line == "}":
            _, msg = err_match.groups()
            expected_val = "true" if neg == "!" else "false"
            expr = expr.strip()
            desc = extract_description(msg)

            new_lines = [
                f'{ws}actual := args.Map{{"result": {expr}}}',
                f'{ws}expected := args.Map{{"result": {expected_val}}}',
                f'{ws}expected.ShouldBeEqual(t, 0, "{desc}", actual)',
            ]
            stats.patterns_transformed += 1
            return (i, i + 2, new_lines)

        # t.Errorf with format args: t.Errorf("expected %d got %d", expected, val)
        err_match_fmt = re.match(
            r't\.(Errorf|Fatalf)\((".*?"),\s*(.*)\)',
            next_line,
        )
        if err_match_fmt and close_line == "}":
            _, msg, fmt_args = err_match_fmt.groups()
            expected_val = "true" if neg == "!" else "false"
            expr = expr.strip()
            desc = extract_description(msg)

            new_lines = [
                f'{ws}actual := args.Map{{"result": {expr}}}',
                f'{ws}expected := args.Map{{"result": {expected_val}}}',
                f'{ws}expected.ShouldBeEqual(t, 0, "{desc}", actual)',
            ]
            stats.patterns_transformed += 1
            return (i, i + 2, new_lines)

    return None


def transform_comparison_if(lines: list, i: int, indent: str, stats: TransformStats) -> Optional[tuple]:
    """
    Transform:
      if val != expected { t.Errorf("msg", args) }
      if val == bad { t.Errorf("msg") }
    
    Multi-line variants too.
    """
    line = lines[i].rstrip()

    # Single-line: if val != expected { t.Errorf("msg %d", val) }
    m = re.match(
        r'^(\s*)if\s+(.+?)\s*(!=|==)\s*(.+?)\s*\{\s*t\.(Error|Errorf|Fatal|Fatalf)\((.*?)\)\s*\}',
        line,
    )
    if m:
        ws, left, op, right, _, msg_and_args = m.groups()
        left = left.strip()
        right = right.strip()
        # Extract just the message part
        msg_parts = msg_and_args.split(',', 1)
        desc = extract_description(msg_parts[0])

        if op == "!=":
            # if val != expected → actual should equal expected
            new_lines = [
                f'{ws}actual := args.Map{{"result": {left}}}',
                f'{ws}expected := args.Map{{"result": {right}}}',
                f'{ws}expected.ShouldBeEqual(t, 0, "{desc}", actual)',
            ]
        else:
            # if val == bad → val should not equal bad (use bool check)
            new_lines = [
                f'{ws}actual := args.Map{{"result": {left} != {right}}}',
                f'{ws}expected := args.Map{{"result": true}}',
                f'{ws}expected.ShouldBeEqual(t, 0, "{desc}", actual)',
            ]
        stats.patterns_transformed += 1
        return (i, i, new_lines)

    # Multi-line: if val != expected {
    m = re.match(r'^(\s*)if\s+(.+?)\s*(!=|==)\s*(.+?)\s*\{\s*$', line)
    if m and i + 2 < len(lines):
        ws, left, op, right = m.groups()
        next_line = lines[i + 1].strip()
        close_line = lines[i + 2].strip() if i + 2 < len(lines) else ""

        err_match = re.match(r't\.(Error|Errorf|Fatal|Fatalf)\((.*)\)', next_line)
        if err_match and close_line == "}":
            _, msg_and_args = err_match.groups()
            msg_parts = msg_and_args.split(',', 1)
            desc = extract_description(msg_parts[0])
            left = left.strip()
            right = right.strip()

            if op == "!=":
                new_lines = [
                    f'{ws}actual := args.Map{{"result": {left}}}',
                    f'{ws}expected := args.Map{{"result": {right}}}',
                    f'{ws}expected.ShouldBeEqual(t, 0, "{desc}", actual)',
                ]
            else:
                new_lines = [
                    f'{ws}actual := args.Map{{"result": {left} != {right}}}',
                    f'{ws}expected := args.Map{{"result": true}}',
                    f'{ws}expected.ShouldBeEqual(t, 0, "{desc}", actual)',
                ]
            stats.patterns_transformed += 1
            return (i, i + 2, new_lines)

    return None


def transform_standalone_error(lines: list, i: int, indent: str, stats: TransformStats) -> Optional[tuple]:
    """
    Transform standalone t.Error/t.Errorf (not inside if block).
    These typically indicate unconditional failure markers.
    """
    line = lines[i].rstrip()
    m = re.match(
        r'^(\s*)t\.(Error|Errorf|Fatal|Fatalf)\((".*?")\s*\)\s*$',
        line,
    )
    if m:
        ws, _, msg = m.groups()
        desc = extract_description(msg)
        # Standalone error = force fail
        new_lines = [
            f'{ws}actual := args.Map{{"result": false}}',
            f'{ws}expected := args.Map{{"result": true}}',
            f'{ws}expected.ShouldBeEqual(t, 0, "{desc}", actual)',
        ]
        stats.patterns_transformed += 1
        return (i, i, new_lines)

    return None


def is_error_line(line: str) -> bool:
    """Check if a line contains t.Error/t.Errorf/t.Fatal/t.Fatalf."""
    return bool(re.search(r'\bt\.(Error|Errorf|Fatal|Fatalf)\b', line))


def transform_file(filepath: str, dry_run: bool, stats: TransformStats) -> bool:
    """Transform a single test file. Returns True if modified."""
    with open(filepath, 'r') as f:
        content = f.read()

    if 't.Error' not in content and 't.Fatal' not in content:
        return False

    stats.files_processed += 1
    lines = content.split('\n')
    new_lines = []
    i = 0
    modified = False
    
    # Detect indentation (tabs vs spaces)
    indent = '\t'

    while i < len(lines):
        line = lines[i]

        # Skip lines that don't have error patterns
        if not is_error_line(line) and not (
            re.match(r'^\s*if\s+', line) and i + 1 < len(lines) and is_error_line(lines[i + 1])
        ):
            new_lines.append(line)
            i += 1
            continue

        # Try each transformer in order
        result = None
        for transformer in [transform_simple_bool_if, transform_comparison_if, transform_standalone_error]:
            result = transformer(lines, i, indent, stats)
            if result:
                break

        if result:
            start, end, replacement = result
            new_lines.extend(replacement)
            i = end + 1
            modified = True
        else:
            # Could not transform — keep original and record
            new_lines.append(line)
            stats.patterns_skipped += 1
            if len(stats.skipped_details) < 100:
                stats.skipped_details.append(f"  {filepath}:{i+1}: {line.strip()}")
            i += 1

    if not modified:
        return False

    new_content = '\n'.join(new_lines)

    # Add args import if needed
    new_content = ensure_args_import(new_content)

    if dry_run:
        rel = os.path.relpath(filepath)
        print(f"  [DRY-RUN] Would modify: {rel}")
    else:
        with open(filepath, 'w') as f:
            f.write(new_content)

    stats.files_modified += 1
    return True


def main():
    parser = argparse.ArgumentParser(description="AAA Format Auto-Transformer")
    parser.add_argument('--dry-run', action='store_true', help='Preview without writing')
    parser.add_argument('--package', type=str, help='Transform only this package (e.g., ostypetests)')
    parser.add_argument('--file', type=str, help='Transform only this file')
    parser.add_argument('--root', type=str, default=TESTS_ROOT, help='Tests root directory')
    args = parser.parse_args()

    stats = TransformStats()

    if args.file:
        transform_file(args.file, args.dry_run, stats)
    else:
        root = args.root
        for dirpath, dirnames, filenames in sorted(os.walk(root)):
            if args.package:
                # Match if the package name appears anywhere in the path
                rel = os.path.relpath(dirpath, root)
                parts = rel.split(os.sep)
                if args.package not in parts:
                    continue
            for fname in sorted(filenames):
                if not fname.endswith("_test.go"):
                    continue
                fpath = os.path.join(dirpath, fname)
                transform_file(fpath, args.dry_run, stats)

    # Report
    print()
    print("=" * 60)
    print("AAA Auto-Transform Report")
    print("=" * 60)
    print(f"Files processed:       {stats.files_processed}")
    print(f"Files modified:        {stats.files_modified}")
    print(f"Patterns transformed:  {stats.patterns_transformed}")
    print(f"Patterns skipped:      {stats.patterns_skipped}")
    print()

    if stats.skipped_details:
        print("Skipped patterns (first 100):")
        for detail in stats.skipped_details[:100]:
            print(detail)
        print()

    if args.dry_run:
        print(">>> DRY RUN — no files were modified. Remove --dry-run to apply.")
    else:
        print(f">>> Done. {stats.files_modified} files updated.")


if __name__ == '__main__':
    main()
