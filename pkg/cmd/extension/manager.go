package extension

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type Manager struct {
	extensionDir string
}

func NewManager() *Manager {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}
	return &Manager{
		extensionDir: filepath.Join(homeDir, ".plz", "extensions"),
	}
}

// ExtensionDir returns the directory where extensions are installed
func (m *Manager) ExtensionDir() string {
	return m.extensionDir
}

// List returns all installed extensions
func (m *Manager) List() ([]string, error) {
	if _, err := os.Stat(m.extensionDir); os.IsNotExist(err) {
		return []string{}, nil
	}

	entries, err := os.ReadDir(m.extensionDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read extensions directory: %w", err)
	}

	var extensions []string
	for _, entry := range entries {
		if !entry.IsDir() {
			// Check if file is executable
			info, err := entry.Info()
			if err != nil {
				continue
			}
			if info.Mode()&0111 != 0 { // Check if any execute bit is set
				extensions = append(extensions, entry.Name())
			}
		}
	}

	return extensions, nil
}

// Install copies an executable file to the extensions directory
func (m *Manager) Install(sourcePath, name string) error {
	// Ensure extension directory exists
	if err := os.MkdirAll(m.extensionDir, 0755); err != nil {
		return fmt.Errorf("failed to create extensions directory: %w", err)
	}

	// Check if source file exists and is executable
	sourceInfo, err := os.Stat(sourcePath)
	if err != nil {
		return fmt.Errorf("source file not found: %w", err)
	}

	if sourceInfo.IsDir() {
		return fmt.Errorf("source path is a directory, not a file")
	}

	// Read source file
	sourceData, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read source file: %w", err)
	}

	// Write to extension directory
	targetPath := filepath.Join(m.extensionDir, name)
	if err := os.WriteFile(targetPath, sourceData, 0755); err != nil {
		return fmt.Errorf("failed to write extension: %w", err)
	}

	return nil
}

// Uninstall removes an extension
func (m *Manager) Uninstall(name string) error {
	targetPath := filepath.Join(m.extensionDir, name)

	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		return fmt.Errorf("extension '%s' not found", name)
	}

	if err := os.Remove(targetPath); err != nil {
		return fmt.Errorf("failed to remove extension: %w", err)
	}

	return nil
}

// Dispatch checks if the command is an extension and executes it
func (m *Manager) Dispatch(rootCmd *cobra.Command, args []string) (bool, error) {
	if len(args) == 0 {
		return false, nil
	}

	extName := args[0]

	// Check if this is a built-in command
	if cmd, _, err := rootCmd.Find(args); err == nil && cmd != rootCmd {
		return false, nil
	}

	// Check if extension exists
	extPath := filepath.Join(m.extensionDir, extName)
	if _, err := os.Stat(extPath); os.IsNotExist(err) {
		return false, nil
	}

	// Execute the extension
	extArgs := []string{}
	if len(args) > 1 {
		extArgs = args[1:]
	}

	return true, m.executeExtension(extPath, extArgs)
}

// executeExtension runs an extension with the given arguments
func (m *Manager) executeExtension(extPath string, args []string) error {
	cmd := exec.Command(extPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		return fmt.Errorf("failed to execute extension: %w", err)
	}

	return nil
}

// FindExtension searches for an extension by name
func (m *Manager) FindExtension(name string) (string, error) {
	extPath := filepath.Join(m.extensionDir, name)
	if _, err := os.Stat(extPath); os.IsNotExist(err) {
		return "", fmt.Errorf("extension '%s' not found", name)
	}
	return extPath, nil
}

// PrintExtensions prints all installed extensions
func (m *Manager) PrintExtensions(w io.Writer) error {
	extensions, err := m.List()
	if err != nil {
		return err
	}

	if len(extensions) == 0 {
		fmt.Fprintln(w, "No extensions installed.")
		return nil
	}

	fmt.Fprintln(w, "Installed extensions:")
	for _, ext := range extensions {
		fmt.Fprintf(w, "  - %s\n", ext)
	}

	return nil
}

// IsExtension checks if a command name is an extension
func (m *Manager) IsExtension(name string) bool {
	extPath := filepath.Join(m.extensionDir, name)
	info, err := os.Stat(extPath)
	if err != nil {
		return false
	}
	return !info.IsDir() && info.Mode()&0111 != 0
}

// GetExtensionPath returns the full path to an extension
func (m *Manager) GetExtensionPath(name string) string {
	return filepath.Join(m.extensionDir, name)
}

// ValidateExtensionName checks if an extension name is valid
func ValidateExtensionName(name string) error {
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
