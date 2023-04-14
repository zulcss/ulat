package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	Verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "ulat",
	Short: "StarlingX image building system.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Print debug information, during command processing")
}
