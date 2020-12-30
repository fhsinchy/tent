package cmd

import (
	"fmt"
	"strings"

	"github.com/containers/podman/v2/pkg/domain/entities"
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

		if isAll && len(args) == 0 {
			tentContainers := utils.ListTentContainers(connText)

			if len(tentContainers) > 0 {
				for _, container := range utils.ListTentContainers(connText) {
					utils.StopContainer(connText, container.ID)
				}
			} else {
				fmt.Println("no running containers found")
			}
		} else {
			for _, service := range args {
				tentContainers := utils.ListTentContainers(connText)

				filteredTentContainers := utils.FilterContainers(tentContainers, func(s entities.ListContainer) bool { return service == strings.Split(s.Names[0], "-")[1] })

				containerCount := len(filteredTentContainers)

				if containerCount == 1 {
					utils.StopContainer(connText, filteredTentContainers[0].ID)
				} else if containerCount > 1 {
					if isAll {
						for _, tentContainer := range filteredTentContainers {
							if service == strings.Split(tentContainer.Names[0], "-")[1] {
								utils.StopContainer(connText, tentContainer.ID)
							}
						}
					} else {
						var choice int
						fmt.Printf("multiple %s containers found:\n", service)
						for index, tentContainer := range filteredTentContainers {
							fmt.Printf("  %d --> %s\n", index, tentContainer.Names[0])
						}
						fmt.Println("you can execute 'tent stop --all' to stop all running containers")
						fmt.Printf("pick the container you want to stop (0 - %d): ", containerCount-1)
						fmt.Scanln(&choice)
						if choice < containerCount {
							utils.StopContainer(connText, filteredTentContainers[choice].ID)
						}
					}
				} else {
					fmt.Printf("no running %s container found", service)
				}
			}
		}
	},
}

func init() {
	stopCmd.Flags().BoolVarP(&isAll, "all", "a", false, "stops all running services")

	rootCmd.AddCommand(stopCmd)
}
