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

type Service struct {
	Tag         string
	Name        string
	Image       string
	Volume      specgen.NamedVolume
	PortMapping specgen.PortMapping
	Env         map[string]string
	HasVolumes  bool
}

func (service *Service) PullImage(connText *context.Context) {
	exists, err := images.Exists(*connText, service.Image)
	if err != nil {
		log.Fatalln(err)
	}

	if !exists {
		_, err := images.Pull(*connText, service.Image, entities.ImagePullOptions{})
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (service *Service) CreateContainer(connText *context.Context) {
	exists, err := containers.Exists(*connText, service.GetContainerName(), false)
	if err != nil {
		log.Fatalln(err)
	}

	if !exists {
		fmt.Println("creating container...")
		s := specgen.NewSpecGenerator(service.Image+":"+service.Tag, false)
		s.Env = service.Env
		s.Remove = true
		s.Name = service.GetContainerName()
		s.PortMappings = append(s.PortMappings, service.PortMapping)

		if service.HasVolumes {
			s.Volumes = append(s.Volumes, &service.Volume)
		}

		_, err := containers.CreateWithSpec(*connText, s)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (service *Service) StartContainer(connText *context.Context) {
	exists, err := containers.Exists(*connText, service.GetContainerName(), false)
	if err != nil {
		log.Fatalln(err)
	}
	if exists {
		fmt.Println("starting container...")
		err := containers.Start(*connText, service.GetContainerName(), nil)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (service Service) StopContainer(connText *context.Context) {
	running := define.ContainerStateRunning
	_, err := containers.Wait(*connText, service.GetContainerName(), &running)
	if err != nil {
		log.Fatalln(err)
	}

	err = containers.Stop(*connText, service.GetContainerName(), nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func (service *Service) GetContainerName() string {
	container := "tent" + "-" + service.Name + "-" + service.Tag + "-" + strconv.Itoa(int(service.PortMapping.HostPort))

	return container
}