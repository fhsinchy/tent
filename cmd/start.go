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
	"context"
	"fmt"
	"log"
	"os"

	"github.com/containers/podman/v2/libpod/define"
	"github.com/containers/podman/v2/pkg/bindings"
	"github.com/containers/podman/v2/pkg/bindings/containers"
	"github.com/containers/podman/v2/pkg/bindings/images"
	"github.com/containers/podman/v2/pkg/domain/entities"
	"github.com/containers/podman/v2/pkg/specgen"
	"github.com/spf13/cobra"
)

var isDefault bool

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get Podman socket location
		sockDir := os.Getenv("XDG_RUNTIME_DIR")
		socket := "unix:" + sockDir + "/podman/podman.sock"

		// Connect to Podman socket
		connText, err := bindings.NewConnection(context.Background(), socket)
		if err != nil {
			log.Fatalln(err)
		}

		service := args[0]

		switch service {
		case "mysql":
			tag := "latest"
			password := "secret"

			if !isDefault {
				var tagInput string
				var passwordInput string

				fmt.Print("Which tag you want to use? (default: latest): ")
				fmt.Scanln(&tagInput)

				fmt.Print("Password for the root user? (default: secret): ")
				fmt.Scanln(&passwordInput)

				if tagInput != "" {
					tag = tagInput
				}

				if passwordInput != "" {
					password = passwordInput
				}
			}

			rawImage := "docker.io/mysql:" + tag
			fmt.Println("pulling mysql image")
			_, err = images.Pull(connText, rawImage, entities.ImagePullOptions{})
			if err != nil {
				log.Fatalln(err)
			}

			env := make(map[string]string)
			env["MYSQL_ROOT_PASSWORD"] = password

			// Container create
			s := specgen.NewSpecGenerator(rawImage, false)
			s.Name = "tent-mysql"
			s.Remove = true
			s.Env = env
			_, err := containers.CreateWithSpec(connText, s)
			if err != nil {
				log.Fatalln(err)
			}

			// Container start
			fmt.Println("starting mysql container")
			err = containers.Start(connText, "tent-mysql", nil)
			if err != nil {
				log.Fatalln(err)
			}

			running := define.ContainerStateRunning
			_, err = containers.Wait(connText, "tent-mysql", &running)
			if err != nil {
				log.Fatalln(err)
			}
		default:
			fmt.Println("service name is required")
		}
	},
}

func init() {
	startCmd.Flags().BoolVarP(&isDefault, "default", "d", false, "starts the service with default options")

	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
