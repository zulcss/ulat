package cmd

import (
	"github.com/spf13/cobra"
)

var (
	ConfigFile string
	Workspace  string
)

var cmdCreate = &cobra.Command{
	Use:   "create",
	Short: "commands to create an image artifact",
}

func init() {
	rootCmd.AddCommand(cmdCreate)

	cmdCreate.PersistentFlags().StringVarP(&ConfigFile, "config", "c", "", "Path to configuration file.")
	cmdCreate.PersistentFlags().StringVarP(&Workspace, "workspace", "w", "", "Path to workspace directory.")

}
