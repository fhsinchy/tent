package utils

import (
	"context"
	"log"

	"github.com/containers/podman/v2/pkg/bindings/containers"
	"github.com/containers/podman/v2/pkg/specgen"
)

func CreateContainer(connText *context.Context, rawImage string, env map[string]string, containerName string) {
	s := specgen.NewSpecGenerator(rawImage, false)
	s.Name = containerName
	s.Remove = true
	s.Env = env
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
