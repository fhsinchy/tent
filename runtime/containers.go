package runtime

import (
	"fmt"

	nettypes "github.com/containers/common/libnetwork/types"
	"github.com/containers/podman/v5/libpod/define"
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/containers/podman/v5/pkg/bindings/images"
	"github.com/containers/podman/v5/pkg/specgen"
	"github.com/fhsinchy/tent/types"
)

// CreateContainer creates a new container for the given service, pulling the image if needed.
// Returns the container ID and a nil error, or an empty string if the container is already running.
func (r *Runtime) CreateContainer(service *types.Service, restartPolicy string) (string, error) {
	containerExists, err := containers.Exists(r.conn, service.ContainerName(), new(containers.ExistsOptions).WithExternal(false))
	if err != nil {
		return "", fmt.Errorf("checking container existence: %w", err)
	}

	if containerExists {
		ins, err := containers.Inspect(r.conn, service.ContainerName(), new(containers.InspectOptions).WithSize(false))
		if err != nil {
			return "", fmt.Errorf("inspecting container: %w", err)
		}

		if ins.State.Running {
			return "", nil
		}
		return ins.ID, nil
	}

	imageExists, err := images.Exists(r.conn, service.ImageName(), nil)
	if err != nil {
		return "", fmt.Errorf("checking image existence: %w", err)
	}

	if !imageExists {
		_, err := images.Pull(r.conn, service.ImageName(), nil)
		if err != nil {
			return "", fmt.Errorf("pulling image %s: %w", service.ImageName(), err)
		}
	}

	s := specgen.NewSpecGenerator(service.ImageName(), false)
	s.Name = service.ContainerName()

	s.Labels = map[string]string{
		"tent.managed": "true",
		"tent.service": service.Name,
	}

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
		}
		s.Env = e
	}

	if len(service.Volumes) > 0 {
		for _, volume := range service.Volumes {
			vol := specgen.NamedVolume{
				Name: service.ContainerName() + "-" + volume.Name,
				Dest: volume.Dest,
			}
			s.Volumes = append(s.Volumes, &vol)
		}
	}

	if len(service.Command) > 0 {
		s.Command = service.Command
	}

	if restartPolicy != "" {
		s.RestartPolicy = restartPolicy
	}

	createResponse, err := containers.CreateWithSpec(r.conn, s, nil)
	if err != nil {
		return "", fmt.Errorf("creating container: %w", err)
	}

	return createResponse.ID, nil
}

// StartContainer starts a container by ID, waiting until it reaches the running state.
func (r *Runtime) StartContainer(containerID string) error {
	exists, err := containers.Exists(r.conn, containerID, new(containers.ExistsOptions).WithExternal(false))
	if err != nil {
		return fmt.Errorf("checking container existence: %w", err)
	}
	if exists {
		err := containers.Start(r.conn, containerID, nil)
		if err != nil {
			return fmt.Errorf("starting container: %w", err)
		}

		_, err = containers.Wait(r.conn, containerID, new(containers.WaitOptions).WithCondition([]define.ContainerStatus{define.ContainerStateRunning}))
		if err != nil {
			return fmt.Errorf("waiting for container: %w", err)
		}
	}
	return nil
}

// StopContainer stops a running container by dispatching a SIGTERM signal.
func (r *Runtime) StopContainer(containerID string) error {
	exists, err := containers.Exists(r.conn, containerID, new(containers.ExistsOptions).WithExternal(false))
	if err != nil {
		return fmt.Errorf("checking container existence: %w", err)
	}

	if exists {
		ins, err := containers.Inspect(r.conn, containerID, new(containers.InspectOptions).WithSize(false))
		if err != nil {
			return fmt.Errorf("inspecting container: %w", err)
		}

		if ins.State.Running {
			err := containers.Stop(r.conn, containerID, nil)
			if err != nil {
				return fmt.Errorf("stopping container: %w", err)
			}
		}
	}
	return nil
}

// RemoveContainer removes a stopped container.
func (r *Runtime) RemoveContainer(containerID string) error {
	exists, err := containers.Exists(r.conn, containerID, new(containers.ExistsOptions).WithExternal(false))
	if err != nil {
		return fmt.Errorf("checking container existence: %w", err)
	}

	if exists {
		ins, err := containers.Inspect(r.conn, containerID, new(containers.InspectOptions).WithSize(false))
		if err != nil {
			return fmt.Errorf("inspecting container: %w", err)
		}

		if !ins.State.Running {
			_, err := containers.Remove(r.conn, containerID, new(containers.RemoveOptions).WithForce(false).WithVolumes(false))
			if err != nil {
				return fmt.Errorf("removing container: %w", err)
			}
		}
	}
	return nil
}

// ListTentContainers lists all containers started by tent, converting Podman types to ContainerInfo.
func (r *Runtime) ListTentContainers() ([]ContainerInfo, error) {
	filters := map[string][]string{
		"label": {"tent.managed=true"},
	}

	podmanContainers, err := containers.List(r.conn, new(containers.ListOptions).WithFilters(filters))
	if err != nil {
		return nil, fmt.Errorf("listing containers: %w", err)
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
			ID:      c.ID,
			Name:    name,
			Image:   c.Image,
			Ports:   ports,
			Service: c.Labels["tent.service"],
		})
	}

	return result, nil
}

// FilterContainers filters a list of ContainerInfo by service name.
func FilterContainers(tentContainers []ContainerInfo, serviceName string) []ContainerInfo {
	var filtered []ContainerInfo
	for _, c := range tentContainers {
		if c.Service == serviceName {
			filtered = append(filtered, c)
		}
	}
	return filtered
}
