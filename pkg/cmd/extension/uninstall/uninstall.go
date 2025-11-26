package uninstall

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type UninstallOptions struct {
	Name string
}

func NewCmdUninstall() *cobra.Command {
	opts := &UninstallOptions{}

	cmd := &cobra.Command{
		Use:   "uninstall <extension-name>",
		Short: "Uninstall an extension",
		Long:  "Remove an installed extension from the system.",
		Example: `  # Uninstall an extension
  plz extension uninstall my-cmd`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return runUninstall(opts)
		},
	}

	return cmd
}

func runUninstall(opts *UninstallOptions) error {
	if err := validateExtensionName(opts.Name); err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}
	extensionDir := filepath.Join(homeDir, ".plz", "extensions")
	targetPath := filepath.Join(extensionDir, opts.Name)

	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		return fmt.Errorf("extension '%s' not found", opts.Name)
	}

	if err := os.Remove(targetPath); err != nil {
		return fmt.Errorf("failed to remove extension: %w", err)
	}

	fmt.Printf("✓ Extension '%s' uninstalled successfully!\n", opts.Name)

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
