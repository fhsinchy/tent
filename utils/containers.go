package utils

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/containers/podman/v5/libpod/define"
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/containers/podman/v5/pkg/domain/entities"
)

// StartContainer function starts a given container created by the CreateContainer function.
func StartContainer(connText *context.Context, containerID string) {
	exists, err := containers.Exists(*connText, containerID, new(containers.ExistsOptions).WithExternal(false))
	if err != nil {
		log.Fatalln(err)
	}
	if exists {
		err := containers.Start(*connText, containerID, nil)
		if err != nil {
			log.Fatalln(err)
		}

		_, err = containers.Wait(*connText, containerID, new(containers.WaitOptions).WithCondition([]define.ContainerStatus{define.ContainerStateRunning}))
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// StopContainer function stops a running container by dispatching a SIGTERM signal.
func StopContainer(connText *context.Context, containerID string) {
	exists, err := containers.Exists(*connText, containerID, new(containers.ExistsOptions).WithExternal(false))
	if err != nil {
		log.Fatalln(err)
	}

	if exists {
		ins, err := containers.Inspect(*connText, containerID, new(containers.InspectOptions).WithSize(false))
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
	exists, err := containers.Exists(*connText, containerID, new(containers.ExistsOptions).WithExternal(false))
	if err != nil {
		log.Fatalln(err)
	}

	if exists {
		ins, err := containers.Inspect(*connText, containerID, new(containers.InspectOptions).WithSize(false))
		if err != nil {
			log.Fatalln(err)
		}

		if !ins.State.Running {
			fmt.Printf("Removing %s container...\n", containerID)
			_, err := containers.Remove(*connText, containerID, new(containers.RemoveOptions).WithForce(false).WithVolumes(false))
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

// ListTentContainers function lists all containers started by tent.
func ListTentContainers(connText *context.Context) (containerList []entities.ListContainer) {
	filters := map[string][]string{
		"name": {"tent-"},
	}

	containerList, err := containers.List(*connText, new(containers.ListOptions).WithFilters(filters))
	if err != nil {
		log.Fatalln(err)
	}

	return
}

// FilterContainers function filters a list of entities.ListContainer type by running a given callback.
func FilterContainers(containers []entities.ListContainer, serviceName string) (filteredContainers []entities.ListContainer) {
	for _, container := range containers {
		if strings.Split(container.Names[0], "-")[1] == serviceName {
			filteredContainers = append(filteredContainers, container)
		}
	}

	return
}
