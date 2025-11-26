# PLZ - An Extensible CLI Tool

`plz` is a super CLI application built with Go and Cobra, structured similar to the GitHub CLI. It allows you to extend functionality by installing executable files from any programming language as extensions.

## Project Structure

```
plz/
├── cmd/
│   └── plz/
│       └── main.go           # Main entry point
├── pkg/
│   └── cmd/
│       ├── root/
│       │   └── root.go       # Root command
│       └── extension/
│           ├── extension.go   # Extension command group
│           ├── manager.go     # Extension manager logic
│           ├── install/
│           │   └── install.go
│           ├── uninstall/
│           │   └── uninstall.go
│           └── list/
│               └── list.go
├── go.mod
├── go.sum
└── README.md
```

This structure follows the GitHub CLI pattern:
- `cmd/` - Contains the main package and entry point
- `pkg/` - Contains all the application logic and commands
- Extensions are stored in `~/.plz/extensions/`

## Features

- ✅ **Extensible**: Add custom commands by installing executable files
- ✅ **Language Agnostic**: Extensions can be written in any language (Bash, Python, Node.js, Go, etc.)
- ✅ **Simple**: Easy to install and manage extensions
- ✅ **Organized**: Clean folder structure following gh CLI best practices

## Installation

### Build from source

```bash
# clone repository 
git clone https://github.com/Kyuubang/plz.git
cd plz

# Build the binary
go build -o bin/plz cmd/plz/main.go

# Optional: Move to PATH
# sudo mv bin/plz /usr/local/bin/
```

## Usage

### Basic Commands

```bash
# Show help
plz --help

# Show extension commands help
plz extension --help
```

### Extension Management

#### Install an Extension

```bash
# Install with automatic name (uses filename)
plz extension install /path/to/executable

# Install with custom name
plz extension install /path/to/script.sh my-command
```

#### List Extensions

```bash
plz extension list
```

#### Uninstall an Extension

```bash
plz extension uninstall my-command
```

### Running Extensions

Once installed, run extensions directly as subcommands:

```bash
plz <extension-name> [args...]
```

## Quick Test

### 1. Build the CLI

```bash
go build -o bin/plz cmd/plz/main.go
```

### 2. Make the test extension executable

```bash
chmod +x test-extension.sh
```

### 3. Install the test extension

```bash
./bin/plz extension install test-extension.sh hello
```

### 4. Run the extension

```bash
./bin/plz hello
./bin/plz hello arg1 arg2 arg3
```

### 5. List installed extensions

```bash
./bin/plz extension list
```

### 6. Uninstall the extension

```bash
./bin/plz extension uninstall hello
```

## Creating Extensions

Extensions can be any executable file. Here are examples:

### Bash Extension

```bash
#!/bin/bash
echo "Hello from bash extension!"
echo "Arguments: $@"
```

### Python Extension

```python
#!/usr/bin/env python3
import sys

print("Hello from Python extension!")
print(f"Arguments: {sys.argv[1:]}")
```

### Node.js Extension

```javascript
#!/usr/bin/env node

console.log("Hello from Node.js extension!");
console.log("Arguments:", process.argv.slice(2));
```

### Go Extension

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Hello from Go extension!")
    fmt.Printf("Arguments: %v\n", os.Args[1:])
}
```

Make sure to:
1. Add shebang line (for scripting languages)
2. Make the file executable: `chmod +x <file>`
3. Install it using `plz extension install`

## Extension Storage

Extensions are stored in: `~/.plz/extensions/`

Each extension is a standalone executable file that `plz` runs when called.

## Architecture

The `plz` CLI follows these principles:

1. **Command Dispatch**: When you run `plz <command>`, it first checks if it's a built-in command
2. **Extension Fallback**: If not found, it looks for an extension with that name in `~/.plz/extensions/`
3. **Direct Execution**: Extensions are executed directly with all arguments passed through
4. **Exit Codes**: Extension exit codes are preserved

## Development

### Project Layout

Following the GitHub CLI pattern:

- `cmd/plz/main.go` - Entry point, minimal code
- `pkg/cmd/root/root.go` - Root command setup, extension dispatch
- `pkg/cmd/extension/` - Extension management commands
  - `manager.go` - Core extension logic
  - `install/` - Install command
  - `uninstall/` - Uninstall command
  - `list/` - List command

### Adding Built-in Commands

To add new built-in commands:

1. Create a new package under `pkg/cmd/`
2. Implement the command using Cobra
3. Add it to the root command in `pkg/cmd/root/root.go`

Example:

```go
import "github.com/bayhaqi/plz/pkg/cmd/mycommand"

// In NewCmdRoot()
cmd.AddCommand(mycommand.NewCmdMyCommand())
```

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
