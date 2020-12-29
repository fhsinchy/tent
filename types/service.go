package types

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/containers/podman/v2/libpod/define"
	"github.com/containers/podman/v2/pkg/bindings/containers"
	"github.com/containers/podman/v2/pkg/bindings/images"
	"github.com/containers/podman/v2/pkg/domain/entities"
	"github.com/containers/podman/v2/pkg/specgen"
)

// Service describes the properties and methods for a service like MySQL or Redis. All the available services in tent are uses this struct as their type.
type Service struct {
	Tag         string
	Name        string
	Image       string
	Volume      specgen.NamedVolume
	PortMapping specgen.PortMapping
	Env         map[string]string
	HasVolumes  bool
}

// PullImage method pulls the image required for creating a service container from online registries if not found in local system.
func (service *Service) PullImage(connText *context.Context) {
	exists, err := images.Exists(*connText, service.Image)
	if err != nil {
		log.Fatalln(err)
	}

	if !exists {
		fmt.Printf("Pulling %s image from registry...\n", service.Image)
		_, err := images.Pull(*connText, service.Image, entities.ImagePullOptions{})
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// CreateContainer method creates a new container with using a given image pulled by PullImage method.
func (service *Service) CreateContainer(connText *context.Context) {
	exists, err := containers.Exists(*connText, service.GetContainerName(), false)
	if err != nil {
		log.Fatalln(err)
	}

	if !exists {
		fmt.Printf("Creating %s container using %s image...\n", service.GetContainerName(), service.Image+":"+service.Tag)
		s := specgen.NewSpecGenerator(service.Image+":"+service.Tag, false)
		s.Env = service.Env
		s.Remove = true
		s.Name = service.GetContainerName()
		s.PortMappings = append(s.PortMappings, service.PortMapping)

		if service.HasVolumes {
			service.Volume.Name = service.GetVolumeName()
			s.Volumes = append(s.Volumes, &service.Volume)
		}

		_, err := containers.CreateWithSpec(*connText, s)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// StartContainer method starts a given container created by the CreateContainer method.
func (service *Service) StartContainer(connText *context.Context) {
	exists, err := containers.Exists(*connText, service.GetContainerName(), false)
	if err != nil {
		log.Fatalln(err)
	}
	if exists {
		fmt.Printf("Starting %s container...\n", service.GetContainerName())
		err := containers.Start(*connText, service.GetContainerName(), nil)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// StopContainer method stops a running container by dispatching a SIGTERM signal.
func (service Service) StopContainer(connText *context.Context) {
	running := define.ContainerStateRunning
	_, err := containers.Wait(*connText, service.GetContainerName(), &running)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Stopping %s container...\n", service.GetContainerName())
	err = containers.Stop(*connText, service.GetContainerName(), nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// GetContainerName method generates unique name for each container by combining their image tag and exposed port number.
func (service *Service) GetContainerName() string {
	container := "tent" + "-" + service.Name + "-" + service.Tag + "-" + strconv.Itoa(int(service.PortMapping.HostPort))

	return container
}

// GetVolumeName method generates unique name for each volume used by different containers by using their container name.
func (service *Service) GetVolumeName() string {
	volume := service.GetContainerName() + "-" + "data"

	return volume
}
