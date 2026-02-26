package cmd

import (
	"fmt"

	"github.com/fhsinchy/tent/store"
	"github.com/spf13/cobra"
)

// servicesCmd represents the services command
var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Lists all available services",
	Run: func(cmd *cobra.Command, args []string) {
		for _, name := range store.ListServiceNames() {
			fmt.Println(name)
		}
	},
}

func init() {
	rootCmd.AddCommand(servicesCmd)
}
