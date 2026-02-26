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

// ContainerName generates unique name for each container by combining their image tag and exposed port number.
func (s *Service) ContainerName() string {
	return "tent" + "-" + s.Name + "-" + s.Tag + "-" + strconv.Itoa(int(s.PortMappings[0].HostPort))
}

// ApplyInsecure replaces the service's env vars with InsecureEnv for passwordless operation.
func (s *Service) ApplyInsecure() (string, error) {
	if s.InsecureEnv == nil {
		return "", fmt.Errorf("%s does not support insecure mode", s.Name)
	}

	s.Env = s.InsecureEnv
	return s.InsecureInfo, nil
}

// ImageName generates full image name for services by combining their image name and tag.
func (s *Service) ImageName() string {
	return s.Image + ":" + s.Tag
}
