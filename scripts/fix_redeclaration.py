#!/usr/bin/env python3
"""
Fix ':=' redeclaration and type-mismatch errors in Go test files.

Usage:
    python3 scripts/fix_redeclaration.py [--dry-run] [blocked-packages-file]

If no blocked-packages file is given, searches data/blocked-packages*.txt
and data/precommit/api-check.json automatically.

Fixes:
  1. 'no new variables on left side of :=' → replaces ':=' with '='
  2. Type-mismatch where a []string variable is later assigned args.Map →
     splits into separate variable names to avoid type conflict.
"""

import os
import re
import sys
import glob
import json


def find_errors_from_blocked(path):
    """Parse a blocked-packages or api-check text/json file for errors."""
    redecl = {}   # {filepath: set(line_numbers)}
    typemm = {}   # {filepath: {lineno: error_detail}}

    content = open(path, 'r', errors='replace').read()

    # Try JSON (api-check.json)
    lines_to_parse = []
    try:
        data = json.loads(content)
        if isinstance(data, dict):
            for pkg, info in data.items():
                if isinstance(info, dict):
                    for err_line in info.get('errors', []):
                        lines_to_parse.append(str(err_line))
                elif isinstance(info, list):
                    for err_line in info:
                        lines_to_parse.append(str(err_line))
        elif isinstance(data, list):
            for item in data:
                lines_to_parse.append(str(item))
    except (json.JSONDecodeError, ValueError):
        lines_to_parse = content.splitlines()

    redecl_pat = re.compile(r'(\S+?\.go):(\d+).*no new variables on left side of :=')
    typemm_pat = re.compile(r'(\S+?\.go):(\d+).*\[type-mismatch\]\s*(.*)')

    for line in lines_to_parse:
        line = line.strip()
        m = redecl_pat.search(line)
        if m:
            fp = m.group(1)
            ln = int(m.group(2))
            redecl.setdefault(fp, set()).add(ln)
            continue
        m2 = typemm_pat.search(line)
        if m2:
            fp = m2.group(1)
            ln = int(m2.group(2))
            detail = m2.group(3)
            typemm.setdefault(fp, {})[ln] = detail

    return redecl, typemm


def scan_test_files(root_dir):
    """Fallback: scan all _test.go files for := redeclarations."""
    fixes = {}
    for dirpath, _, filenames in os.walk(root_dir):
        for fname in filenames:
            if not fname.endswith('_test.go'):
                continue
            filepath = os.path.join(dirpath, fname)
            try:
                lines = open(filepath, 'r').readlines()
            except Exception:
                continue

            func_vars = set()
            for i, line in enumerate(lines, 1):
                stripped = line.strip()
                if stripped.startswith('func ') or stripped.startswith('func('):
                    func_vars = set()
                m = re.match(r'^\s*(\w+)\s*:=\s*', stripped)
                if m:
                    var = m.group(1)
                    if var in func_vars:
                        fixes.setdefault(filepath, set()).add(i)
                    else:
                        func_vars.add(var)
                m2 = re.match(r'^\s*(\w+)\s*=\s*', stripped)
                if m2:
                    func_vars.add(m2.group(1))
    return fixes


def fix_redeclarations(fixes, dry_run=False):
    """Replace ':=' with '=' on specified lines."""
    total = 0
    for filepath, line_numbers in sorted(fixes.items()):
        if not os.path.exists(filepath):
            print(f"  SKIP (not found): {filepath}")
            continue
        lines = open(filepath, 'r').readlines()
        changed = False
        for ln in sorted(line_numbers):
            idx = ln - 1
            if 0 <= idx < len(lines):
                old = lines[idx]
                new = old.replace(':=', '=', 1)
                if new != old:
                    lines[idx] = new
                    changed = True
                    total += 1
        if changed:
            if dry_run:
                print(f"  WOULD FIX: {filepath} ({len(line_numbers)} := errors)")
            else:
                open(filepath, 'w').writelines(lines)
                print(f"  FIXED: {filepath} ({len(line_numbers)} := errors)")
    return total


def fix_type_mismatches(typemm, dry_run=False):
    """
    Fix type-mismatch where []string var is reassigned to args.Map.
    Pattern: line declares 'actual := someFunc()' returning []string,
    then later 'actual = args.Map{...}' causes type conflict.
    Fix: change the later assignment to use a new variable 'actualMap'.
    """
    total = 0
    for filepath, line_errors in sorted(typemm.items()):
        if not os.path.exists(filepath):
            print(f"  SKIP (not found): {filepath}")
            continue

        lines = open(filepath, 'r').readlines()
        changed = False

        # Group by pairs (assignment line + usage line)
        sorted_lines = sorted(line_errors.keys())

        for ln in sorted_lines:
            idx = ln - 1
            if idx < 0 or idx >= len(lines):
                continue
            detail = line_errors[ln]
            old = lines[idx]

            # Case 1: "cannot use args.Map{...} as []string value in assignment"
            # The line has: actual = args.Map{...}  or  actual := args.Map{...}
            # Fix: change to actualMap = args.Map{...} or actualMap := args.Map{...}
            if 'as []string value in assignment' in detail:
                # Replace 'actual' variable name with 'actualMap'
                new = re.sub(r'\bactual\b', 'actualMap', old, count=1)
                if new != old:
                    lines[idx] = new
                    changed = True
                    total += 1

            # Case 2: "cannot use actual (variable of type []string) as args.Map"
            # The line has: expected.ShouldBeEqual(t, ..., actual)
            # Fix: change 'actual' to 'actualMap'
            elif 'as args.Map value in argument' in detail:
                new = re.sub(r'\bactual\b', 'actualMap', old)
                if new != old:
                    lines[idx] = new
                    changed = True
                    total += 1

        if changed:
            if dry_run:
                print(f"  WOULD FIX TYPE: {filepath} ({len(line_errors)} type errors)")
            else:
                open(filepath, 'w').writelines(lines)
                print(f"  FIXED TYPE: {filepath} ({len(line_errors)} type errors)")

    return total


def find_input_files(root):
    """Find blocked-packages or api-check files."""
    candidates = []
    for pattern in [
        os.path.join(root, 'data', 'blocked-packages*.txt'),
        os.path.join(root, 'data', 'precommit', 'api-check.json'),
        'blocked-packages*.txt',
        'data/blocked-packages*.txt',
        'data/precommit/api-check.json',
    ]:
        candidates.extend(sorted(glob.glob(pattern)))
    return candidates


def main():
    dry_run = '--dry-run' in sys.argv
    args = [a for a in sys.argv[1:] if not a.startswith('--')]

    root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

    # Find input files
    if args:
        input_files = args
    else:
        input_files = find_input_files(root)

    all_redecl = {}
    all_typemm = {}

    if input_files:
        for f in input_files:
            print(f"Parsing: {f}")
            r, t = find_errors_from_blocked(f)
            for fp, lns in r.items():
                all_redecl.setdefault(fp, set()).update(lns)
            for fp, errs in t.items():
                all_typemm.setdefault(fp, {}).update(errs)
    else:
        print("No input files found. Scanning all test files for := errors...")
        all_redecl = scan_test_files(os.path.join(root, 'tests'))

    redecl_count = sum(len(v) for v in all_redecl.values())
    typemm_count = sum(len(v) for v in all_typemm.values())

    if redecl_count == 0 and typemm_count == 0:
        print("No errors found to fix.")
        return

    if dry_run:
        print("\n[DRY RUN — no changes]\n")

    print(f"\n=== Redeclaration errors: {redecl_count} across {len(all_redecl)} files ===")
    fixed1 = fix_redeclarations(all_redecl, dry_run)

    print(f"\n=== Type-mismatch errors: {typemm_count} across {len(all_typemm)} files ===")
    fixed2 = fix_type_mismatches(all_typemm, dry_run)

    mode = "Would fix" if dry_run else "Fixed"
    print(f"\n{mode} {fixed1} redeclarations + {fixed2} type mismatches = {fixed1+fixed2} total")


if __name__ == '__main__':
    main()
