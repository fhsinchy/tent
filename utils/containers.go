package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/containers/podman/v2/libpod/define"
	"github.com/containers/podman/v2/pkg/bindings"
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
		ins, err := containers.Inspect(*connText, containerID, bindings.PFalse)
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

// RemoveContainer function removes a stopped container.
func RemoveContainer(connText *context.Context, containerID string) {
	exists, err := containers.Exists(*connText, containerID, false)
	if err != nil {
		log.Fatalln(err)
	}

	if exists {
		ins, err := containers.Inspect(*connText, containerID, bindings.PFalse)
		if err != nil {
			log.Fatalln(err)
		}

		if !ins.State.Running {
			fmt.Printf("Removing %s container...\n", containerID)
			force := false
			volumes := false
			err := containers.Remove(*connText, containerID, &force, &volumes)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

// ListTentContainers function lists all containers started by tent.
func ListTentContainers(connText *context.Context) []entities.ListContainer {
	filters := map[string][]string{
		"name": {"tent-"},
	}

	containerList, err := containers.List(*connText, filters, nil, nil, nil, nil, nil)
	if err != nil {
		log.Fatalln(err)
	}

	return containerList
}

// FilterContainers function filters a list of entities.ListContainer type by running a given callback.
func FilterContainers(collection []entities.ListContainer, callback func(entities.ListContainer) bool) (ret []entities.ListContainer) {
	for _, item := range collection {
		if callback(item) {
			ret = append(ret, item)
		}
	}
	return
}
