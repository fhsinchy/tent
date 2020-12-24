package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/containers/podman/v2/libpod/define"
	"github.com/containers/podman/v2/pkg/bindings"
	"github.com/containers/podman/v2/pkg/bindings/containers"
	"github.com/containers/podman/v2/pkg/bindings/images"
	"github.com/containers/podman/v2/pkg/domain/entities"
	"github.com/containers/podman/v2/pkg/specgen"
)

func main() {
	// Get Podman socket location
	sockDir := os.Getenv("XDG_RUNTIME_DIR")
	socket := "unix:" + sockDir + "/podman/podman.sock"

	// Connect to Podman socket
	connText, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		log.Fatalln(err)
	}

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "starts a service",
				Action: func(c *cli.Context) error {
					service := c.Args().First()

					switch service {
					case "mysql":
						tag := "latest"
						password := "secret"

						fmt.Println(c.FlagNames())

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
					return nil
				},
			},
			{
				Name:    "stop",
				Aliases: []string{"x"},
				Usage:   "stops a service",
				Action: func(c *cli.Context) error {
					service := c.Args().First()
					switch service {
					case "mysql":
						running := define.ContainerStateRunning
						_, err = containers.Wait(connText, "tent-mysql", &running)
						if err != nil {
							log.Fatalln(err)
						}

						// Container inspect
						ctrData, err := containers.Inspect(connText, "tent-mysql", nil)
						if err != nil {
							log.Fatalln(err)
						}
						fmt.Printf("Container uses image %s\n", ctrData.ImageName)
						fmt.Printf("Container running status is %s\n", ctrData.State.Status)

						// Container stop
						fmt.Println("Stopping the container...")
						err = containers.Stop(connText, "tent-mysql", nil)
						if err != nil {
							log.Fatalln(err)
						}
					default:
						fmt.Println("service name is required")
					}
					return nil
				},
			},
		},
	}

	e := app.Run(os.Args)
	if e != nil {
		log.Fatal(e)
	}
}
