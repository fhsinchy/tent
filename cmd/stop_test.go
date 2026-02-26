package cmd

import (
	"fmt"
	"testing"

	"github.com/fhsinchy/tent/runtime"
)

func TestStopServices(t *testing.T) {
	t.Run("all flag with no args stops all containers", func(t *testing.T) {
		m := &mockEngine{
			listResult: []runtime.ContainerInfo{
				{ID: "aaa", Name: "tent-mysql-latest-3306", Service: "mysql"},
				{ID: "bbb", Name: "tent-redis-latest-6379", Service: "redis"},
			},
		}
		stopServices(m, []string{}, true)

		if len(m.stopped) != 2 {
			t.Errorf("expected 2 stop calls, got %d", len(m.stopped))
		}
		if len(m.removed) != 2 {
			t.Errorf("expected 2 remove calls, got %d", len(m.removed))
		}
	})

	t.Run("all flag with no containers found", func(t *testing.T) {
		m := &mockEngine{listResult: []runtime.ContainerInfo{}}
		stopServices(m, []string{}, true)

		if len(m.stopped) != 0 {
			t.Errorf("expected 0 stop calls, got %d", len(m.stopped))
		}
	})

	t.Run("specific service with one match stops and removes", func(t *testing.T) {
		m := &mockEngine{
			listResult: []runtime.ContainerInfo{
				{ID: "aaa", Name: "tent-mysql-latest-3306", Service: "mysql"},
			},
		}
		stopServices(m, []string{"mysql"}, false)

		if len(m.stopped) != 1 {
			t.Fatalf("expected 1 stop call, got %d", len(m.stopped))
		}
		if m.stopped[0] != "aaa" {
			t.Errorf("stopped ID = %q, want %q", m.stopped[0], "aaa")
		}
		if len(m.removed) != 1 {
			t.Fatalf("expected 1 remove call, got %d", len(m.removed))
		}
		if m.removed[0] != "aaa" {
			t.Errorf("removed ID = %q, want %q", m.removed[0], "aaa")
		}
	})

	t.Run("specific service with no match", func(t *testing.T) {
		m := &mockEngine{listResult: []runtime.ContainerInfo{}}
		stopServices(m, []string{"mysql"}, false)

		if len(m.stopped) != 0 {
			t.Errorf("expected 0 stop calls, got %d", len(m.stopped))
		}
	})

	t.Run("invalid service name skips engine", func(t *testing.T) {
		m := &mockEngine{}
		stopServices(m, []string{"nonexistent"}, false)

		if m.listed != 0 {
			t.Errorf("expected 0 list calls for invalid service, got %d", m.listed)
		}
		if len(m.stopped) != 0 {
			t.Errorf("expected 0 stop calls, got %d", len(m.stopped))
		}
	})

	t.Run("stop error does not prevent removing other containers", func(t *testing.T) {
		m := &mockEngine{
			listResult: []runtime.ContainerInfo{
				{ID: "aaa", Name: "tent-mysql-latest-3306", Service: "mysql"},
				{ID: "bbb", Name: "tent-redis-latest-6379", Service: "redis"},
			},
			stopErr: fmt.Errorf("stop failed"),
		}
		stopServices(m, []string{}, true)

		// Both containers should be attempted
		if len(m.stopped) != 2 {
			t.Errorf("expected 2 stop attempts, got %d", len(m.stopped))
		}
		// Remove should not be called since stop failed
		if len(m.removed) != 0 {
			t.Errorf("expected 0 remove calls after stop errors, got %d", len(m.removed))
		}
	})
}
