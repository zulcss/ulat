package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/zulcss/ulat/pkg/compose"
	"github.com/zulcss/ulat/pkg/constants"
)

var createRootfs = &cobra.Command{
	Use: "rootfs",
	Short: "Process a configuration file, bootstap, install packages and commit the " +
		"result to a compressed tarball",
	Run: func(cmd *cobra.Command, args []string) {
		if ConfigFile == "" {
			log.Fatal("You did not specify a configuration file.")
		}
		if Workspace == "" {
			Workspace, err := os.MkdirTemp(constants.CacheDir, "_workspace")
			if err != nil {
				log.Printf("Failed to create workspace %s: %v", Workspace, err)
				os.Exit(1)
			}

			log.Printf("Workspace not configured, using default workspace: %s", Workspace)
		}

		c := compose.NewComposeContext(ConfigFile, Workspace, Verbose)
		cfg := compose.LoadConfig(c)

		c.Config = cfg

		c.Bootstrap()
	},
}

func init() {
	cmdCreate.AddCommand(createRootfs)
}
