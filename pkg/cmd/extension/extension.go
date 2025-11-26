package extension

import (
	"github.com/Kyuubang/plz/pkg/cmd/extension/install"
	"github.com/Kyuubang/plz/pkg/cmd/extension/list"
	"github.com/Kyuubang/plz/pkg/cmd/extension/uninstall"
	"github.com/spf13/cobra"
)

func NewCmdExtension() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "extension",
		Short:   "Manage plz extensions",
		Long:    "Install, uninstall, and list plz extensions.",
		Aliases: []string{"ext"},
	}

	cmd.AddCommand(install.NewCmdInstall())
	cmd.AddCommand(uninstall.NewCmdUninstall())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
