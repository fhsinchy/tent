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

		service := args[0]

		switch service {
		case "mysql":
			if !isDefault {
				var tag string
				var password string
				var volume string
				var port uint16

				fmt.Print("Which tag you want to use? (default: latest): ")
				fmt.Scanln(&tag)

				fmt.Print("Password for the root user? (default: secret): ")
				fmt.Scanln(&password)

				fmt.Print("Volume name for persisting data? (default: tent-mysql-data): ")
				fmt.Scanln(&volume)

				fmt.Print("Host system port? (default: 3306): ")
				fmt.Scanln(&port)

				if tag != "" {
					services.MySQL.Tag = tag
				}

				if password != "" {
					services.MySQL.Env = map[string]string{
						"MYSQL_ROOT_PASSWORD": password,
					}
				}

				if port != 0 {
					services.MySQL.PortMapping.HostPort = port
				}
			}

			services.MySQL.PullImage(connText)
			services.MySQL.CreateContainer(connText)
			services.MySQL.StartContainer(connText)
		case "mariadb":
			if !isDefault {
				var tag string
				var password string
				var volume string
				var port uint16

				fmt.Print("Which tag you want to use? (default: latest): ")
				fmt.Scanln(&tag)

				fmt.Print("Password for the root user? (default: secret): ")
				fmt.Scanln(&password)

				fmt.Print("Volume name for persisting data? (default: tent-mariadb-data): ")
				fmt.Scanln(&volume)

				fmt.Print("Host system port? (default: 3306): ")
				fmt.Scanln(&port)

				if tag != "" {
					services.MySQL.Tag = tag
				}

				if password != "" {
					services.MariaDB.Env = map[string]string{
						"MYSQL_ROOT_PASSWORD": password,
					}
				}

				if port != 0 {
					services.MySQL.PortMapping.HostPort = port
				}
			}

			services.MariaDB.PullImage(connText)
			services.MariaDB.CreateContainer(connText)
			services.MariaDB.StartContainer(connText)
		case "phpmyadmin":
			if !isDefault {
				var tag string
				var port uint16

				fmt.Print("Which tag you want to use? (default: latest): ")
				fmt.Scanln(&tag)

				fmt.Print("Host system port? (default: 8080): ")
				fmt.Scanln(&port)

				if tag != "" {
					services.PHPMyAdmin.Tag = tag
				}

				if port != 0 {
					services.PHPMyAdmin.PortMapping.HostPort = port
				}
			}

			services.PHPMyAdmin.PullImage(connText)
			services.PHPMyAdmin.CreateContainer(connText)
			services.PHPMyAdmin.StartContainer(connText)
		case "redis":
			if !isDefault {
				var tag string
				var volume string
				var port uint16

				fmt.Print("Which tag you want to use? (default: latest): ")
				fmt.Scanln(&tag)

				fmt.Print("Volume name for persisting data? (default: tent-redis-data): ")
				fmt.Scanln(&volume)

				fmt.Print("Host system port? (default: 6379): ")
				fmt.Scanln(&port)

				if tag != "" {
					services.Redis.Tag = tag
				}

				if port != 0 {
					services.Redis.PortMapping.HostPort = port
				}
			}

			services.Redis.PullImage(connText)
			services.Redis.CreateContainer(connText)
			services.Redis.StartContainer(connText)
		default:
			fmt.Println("invalid service name given")
		}
	},
}

func init() {
	startCmd.Flags().BoolVarP(&isDefault, "default", "d", false, "starts the service with default options")

	rootCmd.AddCommand(startCmd)
}
