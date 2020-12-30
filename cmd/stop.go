package cmd

import (
	"strings"

	"github.com/fhsinchy/tent/utils"
	"github.com/spf13/cobra"
)

var isAll bool

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops a running service",
	Long:  `The stop command stops a runnig service. The service container gets removed automatically once stopped.`,
	Run: func(cmd *cobra.Command, args []string) {
		connText := utils.GetContext()

		if isAll {
			for _, container := range utils.ListTentContainers(connText) {
				utils.StopContainer(connText, container.ID)
			}
		} else {
			for _, service := range args {
				tentContainers := utils.ListTentContainers(connText)

				for _, tentContainer := range tentContainers {
					if service == strings.Split(tentContainer.Names[0], "-")[1] {
						utils.StopContainer(connText, tentContainer.ID)
					}
				}
			}
		}
	},
}

func init() {
	stopCmd.Flags().BoolVarP(&isAll, "all", "a", false, "stops all running services")

	rootCmd.AddCommand(stopCmd)
}
