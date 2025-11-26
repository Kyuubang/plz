# Extension Examples

This directory contains example extensions in various languages to demonstrate the extensibility of `plz`.

## Available Examples

### 1. Node.js Extension (`node-extension.js`)
A simple extension written in Node.js that displays system information.

**Install:**
```bash
chmod +x examples/node-extension.js
plz extension install examples/node-extension.js nodeinfo
```

**Usage:**
```bash
plz nodeinfo
plz nodeinfo arg1 arg2
```

### 2. Ruby Extension (`ruby-extension.rb`)
A simple extension written in Ruby.

**Install:**
```bash
chmod +x examples/ruby-extension.rb
plz extension install examples/ruby-extension.rb rubyinfo
```

**Usage:**
```bash
plz rubyinfo
```

### 3. System Info (`sysinfo.sh`)
A practical Bash extension that displays comprehensive system information.

**Install:**
```bash
chmod +x examples/sysinfo.sh
plz extension install examples/sysinfo.sh sysinfo
```

**Usage:**
```bash
plz sysinfo
```

### 4. Git Status (`git-status.py`)
A useful Python extension that shows git repository status with enhanced information.

**Install:**
```bash
chmod +x examples/git-status.py
plz extension install examples/git-status.py gs
```

**Usage:**
```bash
cd /path/to/git/repo
plz gs
```

## Creating Your Own Extensions

Extensions can be written in any language! Here are the requirements:

1. **Executable**: The file must have execute permissions (`chmod +x`)
2. **Shebang**: For scripts, include a shebang line (e.g., `#!/bin/bash`, `#!/usr/bin/env python3`)
3. **Arguments**: Access command-line arguments using your language's standard method
4. **Exit Codes**: Return appropriate exit codes (0 for success, non-zero for errors)

### Template for New Extensions

#### Bash Template
```bash
#!/bin/bash
set -e

# Your extension logic here
echo "Extension name: $0"
echo "Arguments: $@"

# Exit with appropriate code
exit 0
```

#### Python Template
```python
#!/usr/bin/env python3
import sys

def main():
    print(f"Extension: {sys.argv[0]}")
    print(f"Arguments: {sys.argv[1:]}")
    
    # Your logic here
    
    return 0

if __name__ == "__main__":
    sys.exit(main())
```

#### Node.js Template
```javascript
#!/usr/bin/env node

const args = process.argv.slice(2);

console.log(`Extension: ${process.argv[1]}`);
console.log(`Arguments: ${args}`);

// Your logic here

process.exit(0);
```

## Tips

- Use descriptive names for your extensions
- Add help text when called with `--help` or `-h`
- Handle errors gracefully
- Document dependencies required by your extension
- Consider making extensions configurable via environment variables or config files
