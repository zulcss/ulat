package cmd

import (
	"github.com/spf13/cobra"
)

var cmdCreate = &cobra.Command{
	Use:   "create",
	Short: "commands to create an image artifact",
}

func init() {
	rootCmd.AddCommand(cmdCreate)
}
