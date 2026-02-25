package runtime

import (
	"fmt"
	"log"
	"strings"

	nettypes "github.com/containers/common/libnetwork/types"
	"github.com/containers/podman/v5/libpod/define"
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/containers/podman/v5/pkg/bindings/images"
	"github.com/containers/podman/v5/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// CreateContainer creates a new container for the given service, pulling the image if needed.
func (r *Runtime) CreateContainer(service *types.Service) string {
	containerExists, err := containers.Exists(r.conn, service.GetContainerName(), new(containers.ExistsOptions).WithExternal(false))
	if err != nil {
		log.Fatalln(err)
	}

	if containerExists {
		ins, err := containers.Inspect(r.conn, service.GetContainerName(), new(containers.InspectOptions).WithSize(false))
		if err != nil {
			log.Fatalln(err)
		}

		if ins.State.Running {
			fmt.Printf("%s container already running", service.GetContainerName())
			return ""
		}
		return ins.ID
	}

	imageExists, err := images.Exists(r.conn, service.GetImageName(), nil)
	if err != nil {
		log.Fatalln(err)
	}

	if !imageExists {
		_, err := images.Pull(r.conn, service.GetImageName(), nil)
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Printf("Creating %s container using %s image...\n", service.GetContainerName(), service.GetImageName())
	s := specgen.NewSpecGenerator(service.GetImageName(), false)
	s.Name = service.GetContainerName()

	for _, mapping := range service.PortMappings {
		s.PortMappings = append(s.PortMappings, nettypes.PortMapping{
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

	createResponse, err := containers.CreateWithSpec(r.conn, s, nil)
	if err != nil {
		log.Fatalln(err)
	}

	return createResponse.ID
}

// StartContainer starts a container by ID, waiting until it reaches the running state.
func (r *Runtime) StartContainer(containerID string) {
	exists, err := containers.Exists(r.conn, containerID, new(containers.ExistsOptions).WithExternal(false))
	if err != nil {
		log.Fatalln(err)
	}
	if exists {
		err := containers.Start(r.conn, containerID, nil)
		if err != nil {
			log.Fatalln(err)
		}

		_, err = containers.Wait(r.conn, containerID, new(containers.WaitOptions).WithCondition([]define.ContainerStatus{define.ContainerStateRunning}))
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// StopContainer stops a running container by dispatching a SIGTERM signal.
func (r *Runtime) StopContainer(containerID string) {
	exists, err := containers.Exists(r.conn, containerID, new(containers.ExistsOptions).WithExternal(false))
	if err != nil {
		log.Fatalln(err)
	}

	if exists {
		ins, err := containers.Inspect(r.conn, containerID, new(containers.InspectOptions).WithSize(false))
		if err != nil {
			log.Fatalln(err)
		}

		if ins.State.Running {
			fmt.Printf("Stopping %s container...\n", containerID)
			err := containers.Stop(r.conn, containerID, nil)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

// RemoveContainer removes a stopped container.
func (r *Runtime) RemoveContainer(containerID string) {
	exists, err := containers.Exists(r.conn, containerID, new(containers.ExistsOptions).WithExternal(false))
	if err != nil {
		log.Fatalln(err)
	}

	if exists {
		ins, err := containers.Inspect(r.conn, containerID, new(containers.InspectOptions).WithSize(false))
		if err != nil {
			log.Fatalln(err)
		}

		if !ins.State.Running {
			fmt.Printf("Removing %s container...\n", containerID)
			_, err := containers.Remove(r.conn, containerID, new(containers.RemoveOptions).WithForce(false).WithVolumes(false))
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

// ListTentContainers lists all containers started by tent, converting Podman types to ContainerInfo.
func (r *Runtime) ListTentContainers() []ContainerInfo {
	filters := map[string][]string{
		"name": {"tent-"},
	}

	podmanContainers, err := containers.List(r.conn, new(containers.ListOptions).WithFilters(filters))
	if err != nil {
		log.Fatalln(err)
	}

	result := make([]ContainerInfo, 0, len(podmanContainers))
	for _, c := range podmanContainers {
		var name string
		if len(c.Names) > 0 {
			name = c.Names[0]
		}

		ports := make([]PortInfo, 0, len(c.Ports))
		for _, p := range c.Ports {
			ports = append(ports, PortInfo{
				HostPort:      p.HostPort,
				ContainerPort: p.ContainerPort,
				Protocol:      p.Protocol,
			})
		}

		result = append(result, ContainerInfo{
			ID:    c.ID,
			Name:  name,
			Image: c.Image,
			Ports: ports,
		})
	}

	return result
}

// FilterContainers filters a list of ContainerInfo by service name.
func FilterContainers(containers []ContainerInfo, serviceName string) []ContainerInfo {
	var filtered []ContainerInfo
	for _, container := range containers {
		if strings.Split(container.Name, "-")[1] == serviceName {
			filtered = append(filtered, container)
		}
	}
	return filtered
}
