package cmd

import (
	"fmt"

	"github.com/fhsinchy/tent/services"
	"github.com/fhsinchy/tent/utils"

	"github.com/spf13/cobra"
)

var isDefault bool

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
		connText := utils.GetContext()

		for _, service := range args {
			switch service {
			case "mysql":
				if !isDefault {
					services.MySQL.ShowPrompt()
				}

				utils.StartContainer(connText, services.MySQL.CreateContainer(connText))
			case "mariadb":
				if !isDefault {
					services.MariaDB.ShowPrompt()
				}

				utils.StartContainer(connText, services.MariaDB.CreateContainer(connText))
			case "phpmyadmin":
				if !isDefault {
					services.PHPMyAdmin.ShowPrompt()
				}

				utils.StartContainer(connText, services.PHPMyAdmin.CreateContainer(connText))
			case "postgres":
				if !isDefault {
					services.Postgres.ShowPrompt()
				}

				utils.StartContainer(connText, services.Postgres.CreateContainer(connText))
			case "redis":
				if !isDefault {
					services.Redis.ShowPrompt()
				}

				utils.StartContainer(connText, services.Redis.CreateContainer(connText))
			default:
				fmt.Printf("%s is not a valid service name\n", service)
			}
		}

	},
}

func init() {
	startCmd.Flags().BoolVarP(&isDefault, "default", "d", false, "starts the service with default options")

	rootCmd.AddCommand(startCmd)
}
