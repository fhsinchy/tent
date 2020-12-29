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
	Use:   "start",
	Short: "Starts a new service",
	Long:  `The start command starts a new service inside a container. It sets-up all necessary environment variables and volumes as well.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		connText := utils.GetContext()

		for i := 0; i < len(args); i++ {
			service := args[i]

			switch service {
			case "mysql":
				if !isDefault {
					services.MySQL.ShowPrompt()
				}

				services.MySQL.PullImage(connText)
				services.MySQL.CreateContainer(connText)
				services.MySQL.StartContainer(connText)
			case "mariadb":
				if !isDefault {
					services.MariaDB.ShowPrompt()
				}

				services.MariaDB.PullImage(connText)
				services.MariaDB.CreateContainer(connText)
				services.MariaDB.StartContainer(connText)
			case "phpmyadmin":
				if !isDefault {
					services.PHPMyAdmin.ShowPrompt()
				}

				services.PHPMyAdmin.PullImage(connText)
				services.PHPMyAdmin.CreateContainer(connText)
				services.PHPMyAdmin.StartContainer(connText)
			case "postgres":
				if !isDefault {
					services.Postgres.ShowPrompt()
				}

				services.Postgres.PullImage(connText)
				services.Postgres.CreateContainer(connText)
				services.Postgres.StartContainer(connText)
			case "redis":
				if !isDefault {
					services.Redis.ShowPrompt()
				}

				services.Redis.PullImage(connText)
				services.Redis.CreateContainer(connText)
				services.Redis.StartContainer(connText)
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
