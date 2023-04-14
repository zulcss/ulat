package cmd

import (
	"github.com/spf13/cobra"
	"log"
	// "os"
)

var createRootfs = &cobra.Command{
	Use: "rootfs",
	Short: "Process a configuration file, bootstap, install packages and commit the " +
		"result to a compressed tarball",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("in create rootfs")
	},
}

func init() {
	cmdCreate.AddCommand(createRootfs)
}
