package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/podman/v2/pkg/bindings"
	"github.com/containers/podman/v2/pkg/bindings/images"
	"github.com/containers/podman/v2/pkg/domain/entities"
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

	fmt.Println("pulling hello-world image")
	_, err = images.Pull(connText, "docker.io/hello-world", entities.ImagePullOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
