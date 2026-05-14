#!/usr/bin/env python3
"""
args.Map Multi-Line Reformatter
================================
Reformats single-line args.Map{...} with 2+ key-value pairs to multi-line format.

Before:
    ArrangeInput: args.Map{"isZero": false, "value": 5},

After:
    ArrangeInput: args.Map{
        "isZero": false,
        "value":  5,
    },

Handles nested structs (args.TwoAny{...}), function calls, and trailing commas.

Usage:
  python3 args_map_multiline.py --dry-run                     # preview
  python3 args_map_multiline.py --package argstests            # one package
  python3 args_map_multiline.py                                # fix everything
"""

import re
import os
import sys
import argparse

TESTS_ROOT = "tests/integratedtests"


def count_unescaped(s, open_ch, close_ch):
    """Count nesting depth delta, ignoring characters inside string literals."""
    depth = 0
    in_str = False
    escape = False
    for c in s:
        if escape:
            escape = False
            continue
        if c == '\\':
            escape = True
            continue
        if c == '"':
            in_str = not in_str
            continue
        if in_str:
            continue
        if c == open_ch:
            depth += 1
        elif c == close_ch:
            depth -= 1
    return depth


def find_map_end(lines, start_line, start_col):
    """Find the closing } of args.Map{ starting at (start_line, start_col).
    Returns (end_line, end_col) of the closing brace."""
    depth = 0
    for li in range(start_line, len(lines)):
        line = lines[li]
        col_start = start_col if li == start_line else 0
        in_str = False
        escape = False
        for ci in range(col_start, len(line)):
            c = line[ci]
            if escape:
                escape = False
                continue
            if c == '\\':
                escape = True
                continue
            if c == '"':
                in_str = not in_str
                continue
            if in_str:
                continue
            if c == '{':
                depth += 1
            elif c == '}':
                depth -= 1
                if depth == 0:
                    return (li, ci)
    return (-1, -1)


def split_top_level_entries(content):
    """Split map content at top-level commas (not inside nested {}, (), or strings)."""
    entries = []
    current = []
    depth_brace = 0
    depth_paren = 0
    in_str = False
    escape = False

    for c in content:
        if escape:
            current.append(c)
            escape = False
            continue
        if c == '\\':
            current.append(c)
            escape = True
            continue
        if c == '"':
            in_str = not in_str
            current.append(c)
            continue
        if in_str:
            current.append(c)
            continue
        if c == '{':
            depth_brace += 1
        elif c == '}':
            depth_brace -= 1
        elif c == '(':
            depth_paren += 1
        elif c == ')':
            depth_paren -= 1
        elif c == ',' and depth_brace == 0 and depth_paren == 0:
            entries.append(''.join(current).strip())
            current = []
            continue
        current.append(c)

    remainder = ''.join(current).strip()
    if remainder:
        entries.append(remainder)

    return entries


def needs_reformatting(map_content):
    """Check if the map content has 2+ top-level key-value entries."""
    entries = split_top_level_entries(map_content)
    # Filter out empty entries
    entries = [e for e in entries if e]
    return len(entries) >= 2


def reformat_map(prefix, map_content, suffix, indent):
    """Reformat a single-line args.Map to multi-line."""
    entries = split_top_level_entries(map_content)
    entries = [e for e in entries if e]

    if len(entries) < 2:
        return None  # Don't reformat single-entry maps

    inner_indent = indent + '\t'
    lines = [f"{prefix}args.Map{{"]
    for entry in entries:
        lines.append(f"{inner_indent}{entry},")
    lines.append(f"{indent}}}{suffix}")

    return lines


def process_file(filepath, dry_run):
    """Process a single file. Returns number of reformats applied."""
    with open(filepath, 'r') as f:
        content = f.read()

    if 'args.Map{' not in content:
        return 0

    lines = content.split('\n')
    modifications = 0
    i = 0
    new_lines = []

    while i < len(lines):
        line = lines[i]

        # Find args.Map{ on this line
        match = re.search(r'args\.Map\{', line)
        if not match:
            new_lines.append(line)
            i += 1
            continue

        map_start_col = match.start() + len('args.Map')  # position of {

        # Check if the map closes on the same line
        end_line, end_col = find_map_end(lines, i, map_start_col)

        if end_line != i:
            # Multi-line already or spans lines — skip
            new_lines.append(line)
            i += 1
            continue

        # Extract map content (between { and })
        map_content = line[map_start_col + 1:end_col]

        if not needs_reformatting(map_content):
            new_lines.append(line)
            i += 1
            continue

        # Determine indentation
        indent = ''
        for c in line:
            if c == '\t':
                indent += '\t'
            elif c == ' ':
                indent += ' '
            else:
                break

        # Extract prefix (everything before args.Map{) and suffix (everything after })
        prefix = line[:match.start()]
        suffix = line[end_col + 1:]  # after closing }

        reformatted = reformat_map(prefix, map_content, suffix, indent)
        if reformatted:
            new_lines.extend(reformatted)
            modifications += 1
        else:
            new_lines.append(line)

        i += 1

    if modifications == 0:
        return 0

    new_content = '\n'.join(new_lines)

    if dry_run:
        rel = os.path.relpath(filepath)
        print(f"  [DRY-RUN] {rel}: {modifications} maps reformatted")
    else:
        with open(filepath, 'w') as f:
            f.write(new_content)

    return modifications


def main():
    parser = argparse.ArgumentParser(description="args.Map Multi-Line Reformatter")
    parser.add_argument('--dry-run', action='store_true', help='Preview without writing')
    parser.add_argument('--package', type=str, help='Process only this package')
    parser.add_argument('--file', type=str, help='Process only this file')
    parser.add_argument('--root', type=str, default=TESTS_ROOT, help='Tests root directory')
    parsed = parser.parse_args()

    total_files = 0
    total_modified = 0
    total_reformats = 0

    if parsed.file:
        total_files = 1
        n = process_file(parsed.file, parsed.dry_run)
        if n > 0:
            total_modified = 1
            total_reformats = n
    else:
        root = parsed.root
        for dirpath, dirnames, filenames in sorted(os.walk(root)):
            if parsed.package:
                rel = os.path.relpath(dirpath, root)
                parts = rel.split(os.sep)
                if parsed.package not in parts:
                    continue
            for fname in sorted(filenames):
                if not fname.endswith('.go'):
                    continue
                fpath = os.path.join(dirpath, fname)
                total_files += 1
                n = process_file(fpath, parsed.dry_run)
                if n > 0:
                    total_modified += 1
                    total_reformats += n

    print()
    print("=" * 60)
    print("args.Map Multi-Line Reformatter Report")
    print("=" * 60)
    print(f"Files scanned:           {total_files}")
    print(f"Files modified:          {total_modified}")
    print(f"Maps reformatted:        {total_reformats}")
    print()

    if parsed.dry_run:
        print(">>> DRY RUN — no files were modified. Remove --dry-run to apply.")
    else:
        print(f">>> Done. {total_modified} files updated, {total_reformats} maps reformatted.")


if __name__ == '__main__':
    main()
