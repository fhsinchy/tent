package main

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
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Action: func(c *cli.Context) error {
			fmt.Println("the podman thingy")

			if c.Args().Get(0) == "mysql" {
				// Get Podman socket location
				sockDir := os.Getenv("XDG_RUNTIME_DIR")
				socket := "unix:" + sockDir + "/podman/podman.sock"

				// Connect to Podman socket
				connText, err := bindings.NewConnection(context.Background(), socket)
				if err != nil {
					log.Fatalln(err)
				}

				rawImage := "docker.io/mysql:5.7"
				fmt.Println("pulling mysql image")
				_, err = images.Pull(connText, rawImage, entities.ImagePullOptions{})
				if err != nil {
					log.Fatalln(err)
				}

				env := make(map[string]string)
				env["MYSQL_ROOT_PASSWORD"] = "secret"

				// Container create
				s := specgen.NewSpecGenerator(rawImage, false)
				s.Env = env
				r, err := containers.CreateWithSpec(connText, s)
				if err != nil {
					log.Fatalln(err)
				}

				// Container start
				fmt.Println("starting mysql container")
				err = containers.Start(connText, r.ID, nil)
				if err != nil {
					log.Fatalln(err)
				}

				running := define.ContainerStateRunning
				_, err = containers.Wait(connText, r.ID, &running)
				if err != nil {
					log.Fatalln(err)
				}
			} else {
				fmt.Println("service name is required")
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
