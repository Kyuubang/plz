package list

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func NewCmdList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List installed extensions",
		Long:  "Display all installed extensions.",
		Example: `  # List all extensions
  plz extension list`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList()
		},
	}

	return cmd
}

func runList() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}
	extensionDir := filepath.Join(homeDir, ".plz", "extensions")

	if _, err := os.Stat(extensionDir); os.IsNotExist(err) {
		fmt.Println("No extensions installed.")
		fmt.Printf("\nTo install an extension, use:\n  plz extension install <source-path> [name]\n")
		return nil
	}

	entries, err := os.ReadDir(extensionDir)
	if err != nil {
		return fmt.Errorf("failed to read extensions directory: %w", err)
	}

	var extensions []string
	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			if info.Mode()&0111 != 0 {
				extensions = append(extensions, entry.Name())
			}
		}
	}

	if len(extensions) == 0 {
		fmt.Println("No extensions installed.")
		fmt.Printf("\nTo install an extension, use:\n  plz extension install <source-path> [name]\n")
		return nil
	}

	fmt.Println("Installed extensions:")
	for _, ext := range extensions {
		extPath := filepath.Join(extensionDir, ext)
		info, err := os.Stat(extPath)
		if err != nil {
			continue
		}

		fmt.Printf("  %-20s", ext)
		fmt.Printf(" (executable, %d bytes)\n", info.Size())
	}

	fmt.Printf("\nRun an extension with: plz <extension-name> [args...]\n")

	return nil
}
