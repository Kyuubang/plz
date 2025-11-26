package install

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type InstallOptions struct {
	SourcePath string
	Name       string
}

func NewCmdInstall() *cobra.Command {
	opts := &InstallOptions{}

	cmd := &cobra.Command{
		Use:   "install <source-path> [name]",
		Short: "Install a new extension",
		Long: `Install a new extension from an executable file.
The extension can be written in any language as long as it's executable.`,
		Example: `  # Install extension with automatic name
  plz extension install /path/to/my-extension

  # Install extension with custom name
  plz extension install /path/to/script.sh my-cmd`,
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.SourcePath = args[0]

			if len(args) == 2 {
				opts.Name = args[1]
			} else {
				// Use the base filename as the extension name
				opts.Name = filepath.Base(opts.SourcePath)
			}

			return runInstall(opts)
		},
	}

	return cmd
}

func runInstall(opts *InstallOptions) error {
	if err := validateExtensionName(opts.Name); err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}
	extensionDir := filepath.Join(homeDir, ".plz", "extensions")

	// Ensure extension directory exists
	if err := os.MkdirAll(extensionDir, 0755); err != nil {
		return fmt.Errorf("failed to create extensions directory: %w", err)
	}

	// Check if source file exists and is executable
	sourceInfo, err := os.Stat(opts.SourcePath)
	if err != nil {
		return fmt.Errorf("source file not found: %w", err)
	}

	if sourceInfo.IsDir() {
		return fmt.Errorf("source path is a directory, not a file")
	}

	// Read source file
	sourceData, err := os.ReadFile(opts.SourcePath)
	if err != nil {
		return fmt.Errorf("failed to read source file: %w", err)
	}

	// Write to extension directory
	targetPath := filepath.Join(extensionDir, opts.Name)
	if err := os.WriteFile(targetPath, sourceData, 0755); err != nil {
		return fmt.Errorf("failed to write extension: %w", err)
	}

	fmt.Printf("✓ Extension '%s' installed successfully!\n", opts.Name)
	fmt.Printf("  You can now run it with: plz %s\n", opts.Name)

	return nil
}

func validateExtensionName(name string) error {
	if name == "" {
		return fmt.Errorf("extension name cannot be empty")
	}
	if strings.Contains(name, "/") || strings.Contains(name, "\\") {
		return fmt.Errorf("extension name cannot contain path separators")
	}
	if name == "." || name == ".." {
		return fmt.Errorf("invalid extension name")
	}
	return nil
}
