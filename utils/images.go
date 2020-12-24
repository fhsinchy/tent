package utils

import (
	"context"
	"log"

	"github.com/containers/podman/v2/pkg/bindings/images"
	"github.com/containers/podman/v2/pkg/domain/entities"
)

func PullImage(connText *context.Context, rawImage string) {
	exists, err := images.Exists(*connText, rawImage)
	if err != nil {
		log.Fatalln(err)
	}

	if !exists {
		_, err := images.Pull(*connText, rawImage, entities.ImagePullOptions{})
		if err != nil {
			log.Fatalln(err)
		}
	}
}
