package cmd

import (
	"fmt"
	"strings"

	"github.com/fhsinchy/tent/runtime"
	"github.com/fhsinchy/tent/store"
	"github.com/spf13/cobra"
)

var isAll bool

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop [service]",
	Short: "Stops a running service",
	Long: `
The stop command can stop running containers. This command can be used in following configurations:

  1. tent stop mysql ## stops a running mysql container
  2. tent stop mysql --all ## stops all running mysql containers
  3. tent stop --all ## stops all running containers

Stopped containers will be automatically removed from your system.
Volumes used for persisting data however, will be kept for later usage.
`,
	Run: func(cmd *cobra.Command, args []string) {
		rt := runtime.Connect()

		if isAll && len(args) == 0 {
			tentContainers := rt.ListTentContainers()

			if len(tentContainers) > 0 {
				for _, container := range rt.ListTentContainers() {
					rt.StopContainer(container.ID)
					rt.RemoveContainer(container.ID)
				}
			} else {
				fmt.Println("no running containers found")
			}
		} else {
			for _, service := range args {
				if _, ok := store.Services[service]; ok {
					tentContainers := rt.ListTentContainers()

					filteredTentContainers := runtime.FilterContainers(tentContainers, service)

					containerCount := len(filteredTentContainers)

					if containerCount == 1 {
						rt.StopContainer(filteredTentContainers[0].ID)
						rt.RemoveContainer(filteredTentContainers[0].ID)
					} else if containerCount > 1 {
						if isAll {
							for _, tentContainer := range filteredTentContainers {
								if service == strings.Split(tentContainer.Name, "-")[1] {
									rt.StopContainer(tentContainer.ID)
									rt.RemoveContainer(tentContainer.ID)
								}
							}
						} else {
							var choice int
							fmt.Printf("multiple %s containers found:\n", service)
							for index, tentContainer := range filteredTentContainers {
								fmt.Printf("  %d --> %s\n", index, tentContainer.Name)
							}
							fmt.Println("you can execute 'tent stop --all' to stop all running containers")
							fmt.Printf("pick the container you want to stop (0 - %d): ", containerCount-1)
							fmt.Scanln(&choice)
							if choice < containerCount {
								rt.StopContainer(filteredTentContainers[choice].ID)
								rt.RemoveContainer(filteredTentContainers[choice].ID)
							}
						}
					} else {
						fmt.Printf("no running %s container found", service)
					}
				} else {
					fmt.Printf("%s is not a valid service name\n", service)
				}
			}
		}
	},
}

func init() {
	stopCmd.Flags().BoolVarP(&isAll, "all", "a", false, "stops all running services")

	rootCmd.AddCommand(stopCmd)
}
