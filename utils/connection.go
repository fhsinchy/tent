package utils

import (
	"context"
	"log"
	"os"

	"github.com/containers/podman/v2/pkg/bindings"
)

func GetContext() *context.Context {
	// Get Podman socket location
	sockDir := os.Getenv("XDG_RUNTIME_DIR")
	socket := "unix:" + sockDir + "/podman/podman.sock"

	// Connect to Podman socket
	connText, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		log.Fatalln(err)
	}

	return &connText
}
