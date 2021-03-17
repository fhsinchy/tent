package utils

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/containers/podman/v3/libpod/define"
	"github.com/containers/podman/v3/pkg/bindings/containers"
	"github.com/containers/podman/v3/pkg/domain/entities"
)

// StartContainer function starts a given container created by the CreateContainer function.
func StartContainer(connText *context.Context, containerID string) {
	var containerExistsOptions containers.ExistsOptions
	var pFalse = false
	containerExistsOptions.External = &pFalse
	exists, err := containers.Exists(*connText, containerID, &containerExistsOptions)
	if err != nil {
		log.Fatalln(err)
	}
	if exists {
		err := containers.Start(*connText, containerID, nil)
		if err != nil {
			log.Fatalln(err)
		}

		var containerWaitOptions containers.WaitOptions
		containerWaitOptions.Condition = []define.ContainerStatus{
			define.ContainerStateRunning,
		}
		_, err = containers.Wait(*connText, containerID, &containerWaitOptions)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// StopContainer function stops a running container by dispatching a SIGTERM signal.
func StopContainer(connText *context.Context, containerID string) {
	var containerExistsOptions containers.ExistsOptions
	var pFalse = false
	containerExistsOptions.External = &pFalse
	exists, err := containers.Exists(*connText, containerID, &containerExistsOptions)
	if err != nil {
		log.Fatalln(err)
	}

	if exists {
		var containerInspectOptions containers.InspectOptions
		containerInspectOptions.Size = &pFalse
		ins, err := containers.Inspect(*connText, containerID, &containerInspectOptions)
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
	var containerExistsOptions containers.ExistsOptions
	var pFalse = false
	containerExistsOptions.External = &pFalse
	exists, err := containers.Exists(*connText, containerID, &containerExistsOptions)
	if err != nil {
		log.Fatalln(err)
	}

	if exists {
		var containerInspectOptions containers.InspectOptions
		containerInspectOptions.Size = &pFalse
		ins, err := containers.Inspect(*connText, containerID, &containerInspectOptions)
		if err != nil {
			log.Fatalln(err)
		}

		if !ins.State.Running {
			fmt.Printf("Removing %s container...\n", containerID)
			var containerRemoveOptions containers.RemoveOptions
			containerRemoveOptions.Force = &pFalse
			containerRemoveOptions.Volumes = &pFalse
			err := containers.Remove(*connText, containerID, &containerRemoveOptions)
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

	var containerListOptions containers.ListOptions
	containerListOptions.Filters = filters

	containerList, err := containers.List(*connText, &containerListOptions)
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
