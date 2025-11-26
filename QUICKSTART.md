# PLZ CLI - Quick Start Guide

## What is PLZ?

`plz` is an extensible command-line tool built with Go and Cobra. It follows the GitHub CLI architecture and allows you to extend its functionality by installing executable files written in any programming language.

## Quick Setup

### 1. Build the CLI

```bash
# Using make (recommended)
make build

# Or using go directly
go build -o bin/plz cmd/plz/main.go
```

### 2. Verify Installation

```bash
./bin/plz --help
```

You should see the help output showing available commands.

## Basic Usage

### Extension Management

```bash
# List installed extensions
plz extension list

# Install an extension
plz extension install <path-to-executable> <extension-name>

# Uninstall an extension
plz extension uninstall <extension-name>

# Run an extension
plz <extension-name> [args...]
```

### Quick Examples

#### 1. Install and Run the Bash Test Extension

```bash
chmod +x test-extension.sh
./bin/plz extension install test-extension.sh hello
./bin/plz hello
./bin/plz hello world
```

#### 2. Install and Run the Python Extension

```bash
chmod +x python-extension.py
./bin/plz extension install python-extension.py pyinfo
./bin/plz pyinfo
```

#### 3. Install System Info Extension

```bash
chmod +x examples/sysinfo.sh
./bin/plz extension install examples/sysinfo.sh sysinfo
./bin/plz sysinfo
```

## Creating Your First Extension

### Simple Bash Extension

1. Create a file `my-extension.sh`:

```bash
#!/bin/bash
echo "Hello from my extension!"
echo "Arguments: $@"
```

2. Make it executable:

```bash
chmod +x my-extension.sh
```

3. Install it:

```bash
plz extension install my-extension.sh myext
```

4. Run it:

```bash
plz myext arg1 arg2
```

## Extensions Storage

All extensions are stored in: `~/.plz/extensions/`

You can manually inspect or manage files there if needed.

## Project Structure

```
plz/
├── cmd/plz/main.go              # Entry point
├── pkg/cmd/
│   ├── root/root.go            # Root command & dispatch logic
│   └── extension/              # Extension management
│       ├── extension.go        # Extension command group
│       ├── manager.go          # Core extension logic
│       ├── install/            # Install command
│       ├── uninstall/          # Uninstall command
│       └── list/               # List command
├── examples/                    # Example extensions
├── Makefile                     # Build automation
└── README.md                    # Full documentation
```

## Make Commands

```bash
make build      # Build the binary
make clean      # Remove build artifacts
make install    # Install to /usr/local/bin
make build-all  # Build for multiple platforms
make help       # Show all make targets
```

## Tips

- Extensions can be written in **any language** (Bash, Python, Ruby, Node.js, Go, etc.)
- Extensions receive arguments passed after the extension name
- Extensions should return proper exit codes (0 = success, non-zero = error)
- Use descriptive names for your extensions
- Check the `examples/` directory for more inspiration

## Troubleshooting

### Extension not found
- Make sure you installed it: `plz extension list`
- Check the file is in `~/.plz/extensions/`
- Verify the file is executable: `ls -la ~/.plz/extensions/`

### Extension fails to run
- Check the shebang line is correct (e.g., `#!/bin/bash`)
- Verify required interpreters are installed (python3, node, ruby, etc.)
- Test the extension directly: `~/.plz/extensions/<extension-name>`

### Permission denied
- Make sure the original file was executable: `chmod +x <file>`
- The install command copies the executable bit

## Next Steps

1. ✅ Build the CLI: `make build`
2. ✅ Install test extension: `./bin/plz extension install test-extension.sh hello`
3. ✅ Run it: `./bin/plz hello`
4. ✅ Try more examples from `examples/` directory
5. ✅ Create your own extension!
6. 🚀 Optional: Install globally: `make install` (requires sudo)

## Resources

- Full documentation: [README.md](README.md)
- Example extensions: [examples/](examples/)
- GitHub CLI (inspiration): https://github.com/cli/cli

Happy extending! 🎉
