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
	Prompts     map[string]bool
}

// PullImage method pulls the image required for creating a service container from online registries if not found in local system.
func (service *Service) PullImage(connText *context.Context) {
	exists, err := images.Exists(*connText, service.GetImageName())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(exists)

	if !exists {
		fmt.Printf("Pulling %s image from registry...\n", service.GetImageName())
		_, err := images.Pull(*connText, service.GetImageName(), entities.ImagePullOptions{})
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
		fmt.Printf("Creating %s container using %s image...\n", service.GetContainerName(), service.GetImageName())
		s := specgen.NewSpecGenerator(service.GetImageName(), false)
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

		running := define.ContainerStateRunning
		_, err = containers.Wait(*connText, service.GetContainerName(), &running)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// StopContainer method stops a running container by dispatching a SIGTERM signal.
func (service Service) StopContainer(connText *context.Context) {
	exists, err := containers.Exists(*connText, service.GetContainerName(), false)
	if err != nil {
		log.Fatalln(err)
	}

	if exists {
		size := false
		ins, err := containers.Inspect(*connText, service.GetContainerName(), &size)
		if err != nil {
			log.Fatalln(err)
		}

		if ins.State.Running {
			fmt.Printf("Stopping %s container...\n", service.GetContainerName())
			err := containers.Stop(*connText, service.GetContainerName(), nil)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

// ShowPrompt method presents user with user friendly prompts.
func (service *Service) ShowPrompt() {
	if service.Prompts["tag"] {
		var tag string
		fmt.Print("Which tag you want to use? (default: latest): ")
		fmt.Scanln(&tag)
		if tag != "" {
			service.Tag = tag
		}
	}

	if service.Prompts["port"] {
		var port uint16
		fmt.Print("Host system port? (default: 3306): ")
		fmt.Scanln(&port)
		if port != 0 {
			service.PortMapping.HostPort = port
		}
	}

	if service.Prompts["password"] {
		keys := map[string]string{
			"mysql":    "MYSQL_ROOT_PASSWORD",
			"mariadb":  "MYSQL_ROOT_PASSWORD",
			"postgres": "POSTGRES_PASSWORD",
		}

		var password string
		fmt.Print("Password for the root user? (default: secret): ")
		fmt.Scanln(&password)
		if password != "" {
			service.Env[keys[service.Name]] = password
		}
	}

	if service.Prompts["volume"] {
		var volume string
		fmt.Printf("Volume name for persisting data? (default: %s): ", service.GetVolumeName())
		fmt.Scanln(&volume)
		if volume != "" {
			service.Volume.Name = volume
		}
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

// GetImageName method generates full image name for services by combining their image name and tag.
func (service *Service) GetImageName() string {
	image := service.Image + ":" + service.Tag

	return image
}
