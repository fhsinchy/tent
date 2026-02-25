package types

import (
	"fmt"
	"strconv"
)

// Service describes the properties and methods for a service like MySQL or Redis. All the available services in tent uses this struct as their type.
type Service struct {
	Tag          string
	Name         string
	Image        string
	Volumes      []VolumeMount
	PortMappings []PortMapping
	Env          []EnvVar
	InsecureEnv  []EnvVar
	InsecureInfo string
	Command      []string
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

// ApplyInsecure replaces the service's env vars with InsecureEnv for passwordless operation.
func (service *Service) ApplyInsecure() (string, error) {
	if service.InsecureEnv == nil {
		return "", fmt.Errorf("%s does not support insecure mode", service.Name)
	}

	service.Env = service.InsecureEnv
	return service.InsecureInfo, nil
}

// GetImageName method generates full image name for services by combining their image name and tag.
func (service *Service) GetImageName() (imageName string) {
	imageName = service.Image + ":" + service.Tag

	return
}
