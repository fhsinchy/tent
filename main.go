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
	app := &cli.App{
		Action: func(c *cli.Context) error {
			fmt.Println("the podman thingy")

			// Get Podman socket location
			sockDir := os.Getenv("XDG_RUNTIME_DIR")
			socket := "unix:" + sockDir + "/podman/podman.sock"

			// Connect to Podman socket
			connText, err := bindings.NewConnection(context.Background(), socket)
			if err != nil {
				log.Fatalln(err)
			}

			command := c.Args().Get(0)
			service := c.Args().Get(1)

			if command == "start" {
				switch service {
				case "mysql":
					tag := "latest"
					password := "password"

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
			} else if command == "stop" {
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
					ctrData, err = containers.Inspect(connText, "tent-mysql", nil)
					if err != nil {
						log.Fatalln(err)
					}
					fmt.Printf("Container running status is now %s\n", ctrData.State.Status)
				default:
					fmt.Println("service name is required")
				}
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
