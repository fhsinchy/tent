package cmd

import (
	"fmt"
	"log"

	"github.com/fhsinchy/tent/runtime"
	"github.com/fhsinchy/tent/store"
	"github.com/fhsinchy/tent/types"

	"github.com/spf13/cobra"
)

var isDefault bool
var insecure bool
var restartPolicy string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start service",
	Short: "Starts a new service",
	Long: `
The start command can start new containers. This command can be used in following configurations:

  1. tent start mysql --default ## starts a new mysql container with default configuration
  2. tent start mysql ## starts a new mysql container but prompts you for configuration

The start command will take care of pulling images if not found in local registries.
It also sets up necessary named volumes for persisting data.
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rt, err := runtime.Connect()
		if err != nil {
			log.Fatalln(err)
		}

		for _, service := range args {
			s, ok := store.GetService(service)
			if !ok {
				fmt.Printf("%s is not a valid service name\n", service)
				continue
			}

			if insecure {
				info, err := s.ApplyInsecure()
				if err != nil {
					fmt.Println(err)
					continue
				}
				if info != "" {
					fmt.Printf("insecure mode: %s\n", info)
				}
			}

			if !isDefault {
				promptForService(&s)
			}

			fmt.Printf("Creating %s container using %s image...\n", s.ContainerName(), s.ImageName())
			containerID, err := rt.CreateContainer(&s, restartPolicy)
			if err != nil {
				fmt.Printf("error creating %s container: %s\n", service, err)
				continue
			}

			if containerID == "" {
				fmt.Printf("%s container already running\n", s.ContainerName())
				continue
			}

			if err := rt.StartContainer(containerID); err != nil {
				fmt.Printf("error starting %s container: %s\n", service, err)
				continue
			}
		}
	},
}

func promptForService(s *types.Service) {
	var tag string
	fmt.Printf("Which tag do you want to use? (default: %s): ", s.Tag)
	fmt.Scanln(&tag)
	if tag != "" {
		s.Tag = tag
	}

	for index, mapping := range s.PortMappings {
		var port uint16
		fmt.Printf("%s? (default: %d): ", mapping.Text, mapping.HostPort)
		fmt.Scanln(&port)
		if port != 0 {
			s.PortMappings[index].HostPort = port
		}
	}

	for index, env := range s.Env {
		if env.Mutable {
			var value string
			fmt.Printf("%s? (default: %s): ", env.Text, env.Value)
			fmt.Scanln(&value)
			if value != "" {
				s.Env[index].Value = value
			}
		}
	}

	for index, volume := range s.Volumes {
		var name string
		fmt.Printf("%s? (default: %s): ", volume.Text, volume.Name)
		fmt.Scanln(&name)
		if name != "" {
			s.Volumes[index].Name = name
		}
	}
}

func init() {
	startCmd.Flags().BoolVarP(&isDefault, "default", "d", false, "starts the service with default options")
	startCmd.Flags().BoolVar(&insecure, "insecure", false, "start the service without authentication")
	startCmd.Flags().StringVarP(&restartPolicy, "restart", "r", "", "restart policy (no, always, on-failure[:max-retries], unless-stopped)")

	rootCmd.AddCommand(startCmd)
}
