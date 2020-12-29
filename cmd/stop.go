package cmd

import (
	"github.com/fhsinchy/tent/services"
	"github.com/fhsinchy/tent/types"

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

		var services = map[string]types.Service{
			"mysql":      services.MySQL,
			"mariadb":    services.MariaDB,
			"phpmyadmin": services.PHPMyAdmin,
			"redis":      services.Redis,
		}

		if isAll {
			for _, instance := range services {
				instance.StopContainer(connText)
			}
		} else {
			for i := 0; i < len(args); i++ {
				service := args[i]

				services[service].StopContainer(connText)
			}
		}
	},
}

func init() {
	stopCmd.Flags().BoolVarP(&isAll, "all", "a", false, "stops all running services")

	rootCmd.AddCommand(stopCmd)
}
