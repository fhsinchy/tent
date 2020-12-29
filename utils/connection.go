package utils

import (
	"context"
	"log"
	"os"

	"github.com/containers/podman/v2/pkg/bindings"
)

// GetContext function returns a context by making connection with the Podman System Service.
func GetContext() *context.Context {
	sockDir := os.Getenv("XDG_RUNTIME_DIR")
	socket := "unix:" + sockDir + "/podman/podman.sock"

	connText, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		log.Fatalln(err)
	}

	return &connText
}
