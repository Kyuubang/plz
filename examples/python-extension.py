#!/usr/bin/env python3

import sys
import datetime

print("🐍 Python Extension Demo")
print("=" * 50)
print(f"Current time: {datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
print(f"Python version: {sys.version.split()[0]}")
print(f"Arguments passed: {sys.argv[1:]}")
print("=" * 50)
print("\nThis extension is written in Python!")
print("Extensions can be in ANY language! 🚀")
