package runtime

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/podman/v5/pkg/bindings"
	"github.com/fhsinchy/tent/types"
)

// ContainerEngine abstracts container lifecycle operations so callers can be
// tested with a mock implementation.
type ContainerEngine interface {
	CreateContainer(service *types.Service, restartPolicy string) (string, error)
	StartContainer(containerID string) error
	StopContainer(containerID string) error
	RemoveContainer(containerID string) error
	ListTentContainers() ([]ContainerInfo, error)
}

// Compile-time check: *Runtime implements ContainerEngine.
var _ ContainerEngine = (*Runtime)(nil)

// Runtime holds a Podman connection and exposes methods for container operations.
type Runtime struct {
	conn context.Context
}

// Connect establishes a connection with the Podman System Service and returns a Runtime.
func Connect() (*Runtime, error) {
	sockDir := os.Getenv("XDG_RUNTIME_DIR")
	socket := "unix:" + sockDir + "/podman/podman.sock"

	conn, err := bindings.NewConnection(context.Background(), socket)
	if err != nil {
		return nil, fmt.Errorf("connecting to podman: %w", err)
	}

	return &Runtime{conn: conn}, nil
}
