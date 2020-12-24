package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/containers/podman/v2/pkg/bindings/images"
	"github.com/containers/podman/v2/pkg/domain/entities"
)

func PullImage(connText *context.Context, rawImage string) {
	fmt.Println("pulling mysql image")
	_, err := images.Pull(*connText, rawImage, entities.ImagePullOptions{})
	if err != nil {
		log.Fatalln(err)
	}
}
