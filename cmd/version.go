package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "development"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the program version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
