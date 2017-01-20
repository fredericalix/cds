package admin

import (
	"github.com/spf13/cobra"

	"github.com/ovh/cds/cli/cds/admin/maintenance"
	"github.com/ovh/cds/cli/cds/admin/plugin"
	"github.com/ovh/cds/cli/cds/admin/repositoriesmanager"
	"github.com/ovh/cds/cli/cds/admin/template"
	"github.com/ovh/cds/cli/cds/admin/warning"
)

var (
	rootCmd = &cobra.Command{
		Use:   "admin",
		Short: "CDS Admin Management",
	}
)

func init() {
	rootCmd.AddCommand(maintenance.Cmd())
	rootCmd.AddCommand(plugin.Cmd())
	rootCmd.AddCommand(repositoriesmanager.Cmd())
	rootCmd.AddCommand(template.Cmd())
	rootCmd.AddCommand(warning.Cmd())
}

//Cmd returns the root command
func Cmd() *cobra.Command {
	return rootCmd
}