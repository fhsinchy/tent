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
					var tag string
					var password string
					var volume string
					var port uint16

					fmt.Print("Which tag you want to use? (default: latest): ")
					fmt.Scanln(&tag)

					fmt.Print("Password for the root user? (default: secret): ")
					fmt.Scanln(&password)

					fmt.Printf("Volume name for persisting data? (default: %s): ", services.MySQL.GetVolumeName())
					fmt.Scanln(&volume)

					fmt.Print("Host system port? (default: 3306): ")
					fmt.Scanln(&port)

					if tag != "" {
						services.MySQL.Tag = tag
					}

					if password != "" {
						services.MySQL.Env["MYSQL_ROOT_PASSWORD"] = password
					}

					if port != 0 {
						services.MySQL.PortMapping.HostPort = port
					}

					if volume != "" {
						services.MySQL.Volume.Name = volume
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

					fmt.Printf("Volume name for persisting data? (default: %s): ", services.MariaDB.GetVolumeName())
					fmt.Scanln(&volume)

					fmt.Print("Host system port? (default: 3306): ")
					fmt.Scanln(&port)

					if tag != "" {
						services.MySQL.Tag = tag
					}

					if password != "" {
						services.MariaDB.Env["MYSQL_ROOT_PASSWORD"] = password
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
			case "postgres":
				if !isDefault {
					var tag string
					var password string
					var volume string
					var port uint16

					fmt.Print("Which tag you want to use? (default: latest): ")
					fmt.Scanln(&tag)

					fmt.Print("Password for the root user? (default: secret): ")
					fmt.Scanln(&password)

					fmt.Printf("Volume name for persisting data? (default: %s): ", services.Postgres.GetVolumeName())
					fmt.Scanln(&volume)

					fmt.Print("Host system port? (default: 3306): ")
					fmt.Scanln(&port)

					if tag != "" {
						services.Postgres.Tag = tag
					}

					if password != "" {
						services.Postgres.Env["POSTGRES_PASSWORD"] = password
					}

					if port != 0 {
						services.Postgres.PortMapping.HostPort = port
					}

					if volume != "" {
						services.Postgres.Volume.Name = volume
					}
				}

				services.Postgres.PullImage(connText)
				services.Postgres.CreateContainer(connText)
				services.Postgres.StartContainer(connText)
			case "redis":
				if !isDefault {
					var tag string
					var volume string
					var port uint16

					fmt.Print("Which tag you want to use? (default: latest): ")
					fmt.Scanln(&tag)

					fmt.Printf("Volume name for persisting data? (default: %s): ", services.Redis.GetVolumeName())
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
		}

	},
}

func init() {
	startCmd.Flags().BoolVarP(&isDefault, "default", "d", false, "starts the service with default options")

	rootCmd.AddCommand(startCmd)
}
