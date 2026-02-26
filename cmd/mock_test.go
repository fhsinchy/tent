package cmd

import (
	"github.com/fhsinchy/tent/runtime"
	"github.com/fhsinchy/tent/types"
)

// mockEngine is a test double for runtime.ContainerEngine.
// Configure return values and inspect call records after tests.
type mockEngine struct {
	// Return values
	createID   string
	createErr  error
	startErr   error
	stopErr    error
	removeErr  error
	listResult []runtime.ContainerInfo
	listErr    error

	// Call tracking
	created []createCall
	started []string
	stopped []string
	removed []string
	listed  int
}

type createCall struct {
	service       *types.Service
	restartPolicy string
}

func (m *mockEngine) CreateContainer(service *types.Service, restartPolicy string) (string, error) {
	m.created = append(m.created, createCall{service: service, restartPolicy: restartPolicy})
	return m.createID, m.createErr
}

func (m *mockEngine) StartContainer(containerID string) error {
	m.started = append(m.started, containerID)
	return m.startErr
}

func (m *mockEngine) StopContainer(containerID string) error {
	m.stopped = append(m.stopped, containerID)
	return m.stopErr
}

func (m *mockEngine) RemoveContainer(containerID string) error {
	m.removed = append(m.removed, containerID)
	return m.removeErr
}

func (m *mockEngine) ListTentContainers() ([]runtime.ContainerInfo, error) {
	m.listed++
	return m.listResult, m.listErr
}
