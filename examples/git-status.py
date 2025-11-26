#!/usr/bin/env python3

"""
Git status extension - Shows git repository status
"""

import subprocess
import sys
import os

def run_command(cmd):
    """Run a shell command and return output"""
    try:
        result = subprocess.run(
            cmd,
            shell=True,
            capture_output=True,
            text=True,
            check=True
        )
        return result.stdout.strip()
    except subprocess.CalledProcessError as e:
        return None

def main():
    print("🔧 Git Status Extension")
    print("=" * 50)
    
    # Check if we're in a git repository
    if not run_command("git rev-parse --git-dir 2>/dev/null"):
        print("❌ Not a git repository")
        sys.exit(1)
    
    # Get current branch
    branch = run_command("git branch --show-current")
    print(f"Branch: {branch}")
    
    # Get remote URL
    remote = run_command("git remote get-url origin 2>/dev/null")
    if remote:
        print(f"Remote: {remote}")
    
    print("\nStatus:")
    status = run_command("git status --short")
    if status:
        print(status)
    else:
        print("  ✓ Working tree clean")
    
    # Count commits ahead/behind
    ahead_behind = run_command(f"git rev-list --left-right --count origin/{branch}...{branch} 2>/dev/null")
    if ahead_behind:
        behind, ahead = ahead_behind.split()
        if ahead != "0" or behind != "0":
            print(f"\n↑ {ahead} ahead, ↓ {behind} behind origin/{branch}")
    
    print("=" * 50)

if __name__ == "__main__":
    main()
