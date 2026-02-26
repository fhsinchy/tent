package cmd

import (
	"fmt"
	"testing"
)

func TestStartServices(t *testing.T) {
	t.Run("valid service creates and starts", func(t *testing.T) {
		m := &mockEngine{createID: "abc123"}
		startServices(m, []string{"mysql"}, true, false, "")

		if len(m.created) != 1 {
			t.Fatalf("expected 1 create call, got %d", len(m.created))
		}
		if m.created[0].service.Name != "mysql" {
			t.Errorf("created service = %q, want %q", m.created[0].service.Name, "mysql")
		}
		if len(m.started) != 1 {
			t.Fatalf("expected 1 start call, got %d", len(m.started))
		}
		if m.started[0] != "abc123" {
			t.Errorf("started ID = %q, want %q", m.started[0], "abc123")
		}
	})

	t.Run("invalid service skips engine calls", func(t *testing.T) {
		m := &mockEngine{createID: "abc123"}
		startServices(m, []string{"nonexistent"}, true, false, "")

		if len(m.created) != 0 {
			t.Errorf("expected 0 create calls, got %d", len(m.created))
		}
		if len(m.started) != 0 {
			t.Errorf("expected 0 start calls, got %d", len(m.started))
		}
	})

	t.Run("already running container skips start", func(t *testing.T) {
		m := &mockEngine{createID: ""}
		startServices(m, []string{"mysql"}, true, false, "")

		if len(m.created) != 1 {
			t.Fatalf("expected 1 create call, got %d", len(m.created))
		}
		if len(m.started) != 0 {
			t.Errorf("expected 0 start calls for already running, got %d", len(m.started))
		}
	})

	t.Run("create error skips start", func(t *testing.T) {
		m := &mockEngine{createErr: fmt.Errorf("pull failed")}
		startServices(m, []string{"mysql"}, true, false, "")

		if len(m.created) != 1 {
			t.Fatalf("expected 1 create call, got %d", len(m.created))
		}
		if len(m.started) != 0 {
			t.Errorf("expected 0 start calls after create error, got %d", len(m.started))
		}
	})

	t.Run("restart policy forwarded", func(t *testing.T) {
		m := &mockEngine{createID: "abc123"}
		startServices(m, []string{"redis"}, true, false, "always")

		if len(m.created) != 1 {
			t.Fatalf("expected 1 create call, got %d", len(m.created))
		}
		if m.created[0].restartPolicy != "always" {
			t.Errorf("restartPolicy = %q, want %q", m.created[0].restartPolicy, "always")
		}
	})

	t.Run("insecure mode applied", func(t *testing.T) {
		m := &mockEngine{createID: "abc123"}
		startServices(m, []string{"mysql"}, true, true, "")

		if len(m.created) != 1 {
			t.Fatalf("expected 1 create call, got %d", len(m.created))
		}
		// MySQL with insecure should have MYSQL_ALLOW_EMPTY_PASSWORD
		svc := m.created[0].service
		found := false
		for _, env := range svc.Env {
			if env.Key == "MYSQL_ALLOW_EMPTY_PASSWORD" {
				found = true
				break
			}
		}
		if !found {
			t.Error("insecure mode not applied: MYSQL_ALLOW_EMPTY_PASSWORD not in Env")
		}
	})

	t.Run("insecure on unsupported service skips", func(t *testing.T) {
		m := &mockEngine{createID: "abc123"}
		// Redis does not support insecure mode (InsecureEnv is nil)
		startServices(m, []string{"redis"}, true, true, "")

		if len(m.created) != 0 {
			t.Errorf("expected 0 create calls for unsupported insecure, got %d", len(m.created))
		}
	})

	t.Run("multiple services processed independently", func(t *testing.T) {
		m := &mockEngine{createID: "id1"}
		startServices(m, []string{"mysql", "redis"}, true, false, "")

		if len(m.created) != 2 {
			t.Fatalf("expected 2 create calls, got %d", len(m.created))
		}
		if len(m.started) != 2 {
			t.Fatalf("expected 2 start calls, got %d", len(m.started))
		}
	})
}
