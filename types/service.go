package types

import (
	"context"
	"log"

	"github.com/containers/podman/v2/pkg/bindings/containers"
	"github.com/containers/podman/v2/pkg/specgen"
)

type Service struct {
	Tag         string
	Name        string
	Image       string
	Volume      specgen.NamedVolume
	Container   string
	PortMapping specgen.PortMapping
	Env         map[string]string
}

func (service Service) CreateContainer(connText *context.Context) {
	s := specgen.NewSpecGenerator(service.Image+":"+service.Tag, false)
	s.Env = service.Env
	s.Remove = true
	s.Name = service.Container
	s.Volumes = append(s.Volumes, &service.Volume)
	s.PortMappings = append(s.PortMappings, service.PortMapping)
	_, err := containers.CreateWithSpec(*connText, s)
	if err != nil {
		log.Fatalln(err)
	}
}
