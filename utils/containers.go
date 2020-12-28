package utils

import (
	"context"
	"log"

	"github.com/containers/podman/v2/pkg/bindings/containers"
	"github.com/containers/podman/v2/pkg/specgen"
)

func CreateContainer(connText *context.Context, rawImage string, env map[string]string, containerName string, portMapping specgen.PortMapping, volumeName string) {
	s := specgen.NewSpecGenerator(rawImage, false)
	s.Env = env
	s.Remove = true
	s.Name = containerName
	s.Volumes = append(s.Volumes, &specgen.NamedVolume{
		Name: volumeName,
		Dest: "/var/lib/mysql",
	})
	s.PortMappings = append(s.PortMappings, portMapping)
	_, err := containers.CreateWithSpec(*connText, s)
	if err != nil {
		log.Fatalln(err)
	}
}

func StartContainer(connText *context.Context, containerName string) {
	err := containers.Start(*connText, containerName, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func StopContainer(connText *context.Context, containerName string) {
	err := containers.Stop(*connText, containerName, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
