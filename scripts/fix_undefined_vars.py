#!/usr/bin/env python3
"""
Fix RC2: Reverse over-corrected ':=' → '=' conversions.

The fix_redeclaration.py script changed some first-occurrence ':=' to '='
in scopes where the variable hadn't been declared yet, causing 'undefined' errors.

This script properly tracks brace-level scoping: each {} block is a new scope,
so variables declared inside one block aren't visible in sibling blocks.
It also respects named return values (already declared in function signature).
"""

import re
import sys
import os

# Variables commonly over-corrected
TARGET_VARS = {'actual', 'expected', 'result', 'err', 'ok'}

def find_test_files(root):
    """Find all _test.go files under root."""
    files = []
    for dirpath, _, filenames in os.walk(root):
        for f in filenames:
            if f.endswith('_test.go'):
                files.append(os.path.join(dirpath, f))
    return files

def extract_named_returns(sig):
    """Extract named return variable names from a full function signature string."""
    # Match the return params: ") (name1 type1, name2 type2)" or single return
    m = re.search(r'\)\s*\(([^)]+)\)\s*\{', sig)
    if not m:
        return set()
    names = set()
    for part in m.group(1).split(','):
        part = part.strip()
        tokens = part.split()
        if len(tokens) >= 2:
            names.add(tokens[0])
    return names

def fix_file(filepath):
    """Fix undefined variable errors using brace-level scope tracking."""
    with open(filepath, 'r', errors='replace') as f:
        lines = f.readlines()

    fixed_count = 0
    # Stack of sets: each entry = variables declared at that brace depth
    scope_stack = [set()]  # depth 0 = file level
    
    # For multi-line function signatures
    in_func_sig = False
    func_sig_accum = ""
    
    i = 0
    while i < len(lines):
        line = lines[i]
        stripped = line.strip()
        
        pending_named_returns = set()

        # Detect function start - may span multiple lines
        if re.match(r'^func\s', stripped):
            if '{' in line:
                # Single-line signature
                pending_named_returns = extract_named_returns(line)
            else:
                # Multi-line signature - start accumulating
                in_func_sig = True
                func_sig_accum = line
        elif in_func_sig:
            func_sig_accum += line
            if '{' in line:
                # End of multi-line signature
                in_func_sig = False
                pending_named_returns = extract_named_returns(func_sig_accum)
                func_sig_accum = ""

        # Track := declarations - add to current scope
        decl_match = re.match(r'\s*([\w,\s]+)\s*:=', line)
        if decl_match:
            var_names = decl_match.group(1)
            for v in re.split(r'[,\s]+', var_names):
                v = v.strip()
                if v and v != '_':
                    scope_stack[-1].add(v)

        # Track var declarations
        var_match = re.match(r'\s*var\s+(\w+)', line)
        if var_match:
            scope_stack[-1].add(var_match.group(1))

        # Check for plain assignments to TARGET_VARS that need := 
        for var in TARGET_VARS:
            pattern = rf'^(\s+){var}\s*=\s*(?!=)'
            m = re.match(pattern, line)
            if m:
                # Check it's not == or !=
                if re.match(rf'^\s+{var}\s*[!=]=', line):
                    continue
                # Check if var is declared in ANY enclosing scope
                declared = any(var in s for s in scope_stack)
                if not declared:
                    old = line
                    line = re.sub(rf'^(\s+){var}\s*=\s*', rf'\g<1>{var} := ', line, count=1)
                    if line != old:
                        lines[i] = line
                        fixed_count += 1
                        scope_stack[-1].add(var)

        # Adjust scope stack for braces on this line
        opens = line.count('{')
        closes = line.count('}')
        
        for _ in range(opens):
            new_scope = set()
            if pending_named_returns:
                new_scope = pending_named_returns
                pending_named_returns = set()
            scope_stack.append(new_scope)
        
        for _ in range(closes):
            if len(scope_stack) > 1:
                scope_stack.pop()

        i += 1

    if fixed_count > 0:
        with open(filepath, 'w') as f:
            f.writelines(lines)

    return fixed_count

def main():
    root = sys.argv[1] if len(sys.argv) > 1 else './tests/integratedtests'

    if not os.path.isdir(root):
        print(f"Directory not found: {root}")
        sys.exit(1)

    files = find_test_files(root)
    total_fixed = 0
    fixed_files = 0

    for fp in sorted(files):
        count = fix_file(fp)
        if count > 0:
            print(f"  FIXED: {fp} ({count} := restored)")
            total_fixed += count
            fixed_files += 1

    print(f"\nRestored {total_fixed} assignments across {fixed_files} files")

if __name__ == '__main__':
    main()
