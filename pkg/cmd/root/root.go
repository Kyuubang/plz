package root

import (
	"fmt"
	"os"

	"github.com/Kyuubang/plz/pkg/cmd/extension"
	"github.com/spf13/cobra"
)

type ExitCode int

const (
	ExitOK    ExitCode = 0
	ExitError ExitCode = 1
)

func NewCmdRoot() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "plz",
		Short: "An extensible CLI tool",
		Long: `plz is a super CLI app that can be extended with external executable files.
Extensions can be written in any language as long as they are executable.`,
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	// Add extension management commands
	cmd.AddCommand(extension.NewCmdExtension())

	return cmd, nil
}

func Execute() ExitCode {
	rootCmd, err := NewCmdRoot()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create root command: %s\n", err)
		return ExitError
	}

	// Try to execute extension if it exists
	extManager := extension.NewManager()
	if handled, err := extManager.Dispatch(rootCmd, os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return ExitError
	} else if handled {
		return ExitOK
	}

	// Execute normal command if no extension was found
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return ExitError
	}

	return ExitOK
}
