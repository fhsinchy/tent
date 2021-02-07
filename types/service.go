package types

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/containers/podman/v2/pkg/bindings"
	"github.com/containers/podman/v2/pkg/bindings/containers"
	"github.com/containers/podman/v2/pkg/bindings/images"
	"github.com/containers/podman/v2/pkg/domain/entities"
	"github.com/containers/podman/v2/pkg/specgen"
)

// Service describes the properties and methods for a service like MySQL or Redis. All the available services in tent uses this struct as their type.
type Service struct {
	Tag          string
	Name         string
	Image        string
	Volumes      []VolumeMount
	PortMappings []PortMapping
	Env          []EnvVar
	Command      []string
}

// CreateContainer method creates a new container with using a given image pulled by PullImage method.
func (service *Service) CreateContainer(connText *context.Context) (containerID string) {
	containerExists, err := containers.Exists(*connText, service.GetContainerName(), false)
	if err != nil {
		log.Fatalln(err)
	}

	if containerExists {
		ins, err := containers.Inspect(*connText, service.GetContainerName(), bindings.PFalse)
		if err != nil {
			log.Fatalln(err)
		}

		if ins.State.Running {
			fmt.Printf("%s container already running", service.GetContainerName())
		} else {
			containerID = ins.ID
		}
	} else {
		imageExists, err := images.Exists(*connText, service.GetImageName())
		if err != nil {
			log.Fatalln(err)
		}

		if !imageExists {
			_, err := images.Pull(*connText, service.GetImageName(), entities.ImagePullOptions{})
			if err != nil {
				log.Fatalln(err)
			}
		}

		fmt.Printf("Creating %s container using %s image...\n", service.GetContainerName(), service.GetImageName())
		s := specgen.NewSpecGenerator(service.GetImageName(), false)
		s.Name = service.GetContainerName()

		for _, mapping := range service.PortMappings {
			s.PortMappings = append(s.PortMappings, specgen.PortMapping{
				ContainerPort: mapping.ContainerPort,
				HostPort:      mapping.HostPort,
			})
		}

		if len(service.Env) > 0 {
			e := make(map[string]string)
			for _, env := range service.Env {
				e[env.Key] = env.Value
				s.Env = e
			}
		}

		if len(service.Volumes) > 0 {
			for _, volume := range service.Volumes {
				vol := specgen.NamedVolume{
					Name: volume.Name,
					Dest: volume.Dest,
				}
				s.Volumes = append(s.Volumes, &vol)
			}
		}

		if len(service.Command) > 0 {
			s.Command = service.Command
		}

		createResponse, err := containers.CreateWithSpec(*connText, s)
		if err != nil {
			log.Fatalln(err)
		}

		containerID = createResponse.ID
	}

	return
}

// ShowPrompt method presents user with user friendly prompts.
func (service *Service) ShowPrompt() {
	var tag string
	fmt.Printf("Which tag do you want to use? (default: %s): ", service.Tag)
	fmt.Scanln(&tag)
	if tag != "" {
		service.Tag = tag
	}

	for index, mapping := range service.PortMappings {
		var port uint16
		fmt.Printf("%s? (default: %d): ", mapping.Text, mapping.HostPort)
		fmt.Scanln(&port)
		if port != 0 {
			service.PortMappings[index].HostPort = port
		}
	}

	for index, env := range service.Env {
		if env.Mutable {
			var value string
			fmt.Printf("%s? (default: %s): ", env.Text, env.Value)
			fmt.Scanln(&value)
			if value != "" {
				service.Env[index].Value = value
			}
		}
	}

	for index, volume := range service.Volumes {
		var name string
		fmt.Printf("%s? (default: %s): ", volume.Text, volume.Name)
		fmt.Scanln(&name)
		if name != "" {
			service.Volumes[index].Name = name
		}
	}
}

// GetContainerName method generates unique name for each container by combining their image tag and exposed port number.
func (service *Service) GetContainerName() (containerName string) {
	containerName = "tent" + "-" + service.Name + "-" + service.Tag + "-" + strconv.Itoa(int(service.PortMappings[0].HostPort))

	return
}

// GetVolumeName method generates unique name for each volume used by different containers by using their container name.
func (service *Service) GetVolumeName() (volumeName string) {
	volumeName = service.GetContainerName() + "-" + "data"

	return
}

// GetImageName method generates full image name for services by combining their image name and tag.
func (service *Service) GetImageName() (imageName string) {
	imageName = service.Image + ":" + service.Tag

	return
}
