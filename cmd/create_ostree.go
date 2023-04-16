package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/zulcss/ulat/pkg/compose"
	"github.com/zulcss/ulat/pkg/constants"
)

var createOstree = &cobra.Command{
	Use:   "ostree",
	Short: "Create ostree branch",
	Long:  "Process a configuraiton file, untar rootfs, and create ostree branch",
	Run: func(cmd *cobra.Command, args []string) {
		if ConfigFile == "" {
			log.Fatal("You did not specify a configuration file.")
		}
		if Workspace == "" {
			Workspace, err := os.MkdirTemp(constants.CacheDir, "_workspace")
			if err != nil {
				os.Exit(1)
			}
			log.Printf("Workspace not configured, using default workspace: %s", Workspace)
		}
		c := compose.NewComposeContext(ConfigFile, Workspace, Verbose)
		c.CreateOstree()
	},
}

func init() {
	cmdCreate.AddCommand(createOstree)
}
