package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/podman/v2/libpod/define"
	"github.com/containers/podman/v2/pkg/bindings"
	"github.com/containers/podman/v2/pkg/bindings/containers"
	"github.com/containers/podman/v2/pkg/bindings/images"
	"github.com/containers/podman/v2/pkg/domain/entities"
	"github.com/containers/podman/v2/pkg/specgen"
)

func main() {
	fmt.Println("the podman thingy")

	// Get Podman socket location
	sockDir := os.Getenv("XDG_RUNTIME_DIR")
	socket := "unix:" + sockDir + "/podman/podman.sock"

	// Connect to Podman socket
	connText, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rawImage := "docker.io/mysql"
	fmt.Println("pulling mysql image")
	_, err = images.Pull(connText, rawImage, entities.ImagePullOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	env := make(map[string]string)
	env["MYSQL_ROOT_PASSWORD"] = "secret"

	// Container create
	s := specgen.NewSpecGenerator(rawImage, false)
	s.Env = env
	r, err := containers.CreateWithSpec(connText, s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Container start
	fmt.Println("starting mysql container")
	err = containers.Start(connText, r.ID, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	running := define.ContainerStateRunning
	_, err = containers.Wait(connText, r.ID, &running)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
