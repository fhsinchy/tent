/*
Copyright © 2020 FARHAN HASIN CHOWDHURY <MAIL@FARHAN.INFO>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/fhsinchy/tent/services"
	"github.com/fhsinchy/tent/types"

	"github.com/fhsinchy/tent/utils"
	"github.com/spf13/cobra"
)

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

		for i := 0; i < len(args); i++ {
			service := args[i]

			services[service].StopContainer(connText)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
