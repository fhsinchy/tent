package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/containers/podman/v2/libpod/define"
	"github.com/containers/podman/v2/pkg/bindings/containers"
	"github.com/containers/podman/v2/pkg/domain/entities"
)

// StartContainer function starts a given container created by the CreateContainer function.
func StartContainer(connText *context.Context, containerID string) {
	exists, err := containers.Exists(*connText, containerID, false)
	if err != nil {
		log.Fatalln(err)
	}
	if exists {
		err := containers.Start(*connText, containerID, nil)
		if err != nil {
			log.Fatalln(err)
		}

		running := define.ContainerStateRunning
		_, err = containers.Wait(*connText, containerID, &running)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// StopContainer function stops a running container by dispatching a SIGTERM signal.
func StopContainer(connText *context.Context, containerID string) {
	exists, err := containers.Exists(*connText, containerID, false)
	if err != nil {
		log.Fatalln(err)
	}

	if exists {
		size := false
		ins, err := containers.Inspect(*connText, containerID, &size)
		if err != nil {
			log.Fatalln(err)
		}

		if ins.State.Running {
			fmt.Printf("Stopping %s container...\n", containerID)
			err := containers.Stop(*connText, containerID, nil)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

// ListTentContainers function lists all containers started by tent.
func ListTentContainers(connText *context.Context) []entities.ListContainer {
	filters := map[string][]string{
		"name":   {"tent-"},
		"status": {"running"},
	}

	containerList, err := containers.List(*connText, filters, nil, nil, nil, nil, nil)
	if err != nil {
		log.Fatalln(err)
	}

	return containerList
}
