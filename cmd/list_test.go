package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/fhsinchy/tent/runtime"
)

func TestListContainers(t *testing.T) {
	t.Run("formats table with header and data", func(t *testing.T) {
		m := &mockEngine{
			listResult: []runtime.ContainerInfo{
				{
					Name:  "tent-mysql-latest-3306",
					Image: "docker.io/mysql:latest",
					Ports: []runtime.PortInfo{
						{HostPort: 3306, ContainerPort: 3306, Protocol: "tcp"},
					},
				},
			},
		}

		var buf bytes.Buffer
		err := listContainers(m, &buf)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		output := buf.String()
		if !strings.Contains(output, "CONTAINER") {
			t.Error("output missing CONTAINER header")
		}
		if !strings.Contains(output, "IMAGE") {
			t.Error("output missing IMAGE header")
		}
		if !strings.Contains(output, "PORTS") {
			t.Error("output missing PORTS header")
		}
		if !strings.Contains(output, "tent-mysql-latest-3306") {
			t.Error("output missing container name")
		}
		if !strings.Contains(output, "3306->3306/tcp") {
			t.Error("output missing port mapping")
		}
	})

	t.Run("multiple ports joined with comma", func(t *testing.T) {
		m := &mockEngine{
			listResult: []runtime.ContainerInfo{
				{
					Name:  "tent-rabbitmq-latest-5672",
					Image: "docker.io/rabbitmq:latest",
					Ports: []runtime.PortInfo{
						{HostPort: 5672, ContainerPort: 5672, Protocol: "tcp"},
						{HostPort: 15672, ContainerPort: 15672, Protocol: "tcp"},
					},
				},
			},
		}

		var buf bytes.Buffer
		err := listContainers(m, &buf)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		output := buf.String()
		if !strings.Contains(output, "5672->5672/tcp, 15672->15672/tcp") {
			t.Errorf("expected comma-joined ports, got: %s", output)
		}
	})

	t.Run("empty container list still has header", func(t *testing.T) {
		m := &mockEngine{listResult: []runtime.ContainerInfo{}}

		var buf bytes.Buffer
		err := listContainers(m, &buf)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		output := buf.String()
		if !strings.Contains(output, "CONTAINER") {
			t.Error("output missing header even with no containers")
		}
	})

	t.Run("engine error returned", func(t *testing.T) {
		m := &mockEngine{listErr: fmt.Errorf("connection refused")}

		var buf bytes.Buffer
		err := listContainers(m, &buf)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !strings.Contains(err.Error(), "connection refused") {
			t.Errorf("error = %q, want it to contain 'connection refused'", err.Error())
		}
	})
}
