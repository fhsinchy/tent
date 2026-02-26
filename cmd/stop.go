package cmd

import (
	"fmt"
	"log"

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
		rt, err := runtime.Connect()
		if err != nil {
			log.Fatalln(err)
		}

		if isAll && len(args) == 0 {
			tentContainers, err := rt.ListTentContainers()
			if err != nil {
				log.Fatalln(err)
			}

			if len(tentContainers) > 0 {
				for _, container := range tentContainers {
					fmt.Printf("Stopping %s container...\n", container.Name)
					if err := rt.StopContainer(container.ID); err != nil {
						fmt.Printf("error stopping container %s: %s\n", container.Name, err)
						continue
					}
					fmt.Printf("Removing %s container...\n", container.Name)
					if err := rt.RemoveContainer(container.ID); err != nil {
						fmt.Printf("error removing container %s: %s\n", container.Name, err)
					}
				}
			} else {
				fmt.Println("no running containers found")
			}
		} else {
			for _, service := range args {
				if _, ok := store.GetService(service); !ok {
					fmt.Printf("%s is not a valid service name. Run 'tent services' to see available services.\n", service)
					continue
				}

				tentContainers, err := rt.ListTentContainers()
				if err != nil {
					fmt.Printf("error listing containers: %s\n", err)
					continue
				}

				filteredTentContainers := runtime.FilterContainers(tentContainers, service)

				containerCount := len(filteredTentContainers)

				if containerCount == 1 {
					fmt.Printf("Stopping %s container...\n", filteredTentContainers[0].Name)
					if err := rt.StopContainer(filteredTentContainers[0].ID); err != nil {
						fmt.Printf("error stopping container: %s\n", err)
						continue
					}
					fmt.Printf("Removing %s container...\n", filteredTentContainers[0].Name)
					if err := rt.RemoveContainer(filteredTentContainers[0].ID); err != nil {
						fmt.Printf("error removing container: %s\n", err)
					}
				} else if containerCount > 1 {
					if isAll {
						for _, tentContainer := range filteredTentContainers {
							fmt.Printf("Stopping %s container...\n", tentContainer.Name)
							if err := rt.StopContainer(tentContainer.ID); err != nil {
								fmt.Printf("error stopping container %s: %s\n", tentContainer.Name, err)
								continue
							}
							fmt.Printf("Removing %s container...\n", tentContainer.Name)
							if err := rt.RemoveContainer(tentContainer.ID); err != nil {
								fmt.Printf("error removing container %s: %s\n", tentContainer.Name, err)
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
						if choice >= 0 && choice < containerCount {
							fmt.Printf("Stopping %s container...\n", filteredTentContainers[choice].Name)
							if err := rt.StopContainer(filteredTentContainers[choice].ID); err != nil {
								fmt.Printf("error stopping container: %s\n", err)
								continue
							}
							fmt.Printf("Removing %s container...\n", filteredTentContainers[choice].Name)
							if err := rt.RemoveContainer(filteredTentContainers[choice].ID); err != nil {
								fmt.Printf("error removing container: %s\n", err)
							}
						}
					}
				} else {
					fmt.Printf("no running %s container found\n", service)
				}
			}
		}
	},
}

func init() {
	stopCmd.Flags().BoolVarP(&isAll, "all", "a", false, "stops all running services")
	stopCmd.ValidArgs = store.ListServiceNames()

	rootCmd.AddCommand(stopCmd)
}
