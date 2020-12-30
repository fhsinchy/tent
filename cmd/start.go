package cmd

import (
	"fmt"

	"github.com/fhsinchy/tent/services"
	"github.com/fhsinchy/tent/types"
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

		ss := map[string]*types.Service{
			"mysql":         &services.MySQL,
			"mariadb":       &services.MariaDB,
			"phpmyadmin":    &services.PHPMyAdmin,
			"postgres":      &services.Postgres,
			"mongo":         &services.Mongo,
			"mongo-express": &services.MongoExpress,
			"redis":         &services.Redis,
		}

		for _, service := range args {
			if s, ok := ss[service]; ok {
				if !isDefault {
					s.ShowPrompt()
				}

				utils.StartContainer(connText, s.CreateContainer(connText))
			} else {
				fmt.Printf("%s is not a valid service name\n", service)
			}
		}
	},
}

func init() {
	startCmd.Flags().BoolVarP(&isDefault, "default", "d", false, "starts the service with default options")

	rootCmd.AddCommand(startCmd)
}
