package runtime

import (
	"testing"
)

func TestFilterContainers(t *testing.T) {
	containers := []ContainerInfo{
		{ID: "aaa", Name: "tent-mysql-latest-3306", Service: "mysql"},
		{ID: "bbb", Name: "tent-redis-latest-6379", Service: "redis"},
		{ID: "ccc", Name: "tent-mysql-8.0-3307", Service: "mysql"},
	}

	t.Run("single match", func(t *testing.T) {
		got := FilterContainers(containers, "redis")
		if len(got) != 1 {
			t.Fatalf("expected 1 match, got %d", len(got))
		}
		if got[0].ID != "bbb" {
			t.Errorf("ID = %q, want %q", got[0].ID, "bbb")
		}
	})

	t.Run("multiple matches", func(t *testing.T) {
		got := FilterContainers(containers, "mysql")
		if len(got) != 2 {
			t.Fatalf("expected 2 matches, got %d", len(got))
		}
	})

	t.Run("no matches", func(t *testing.T) {
		got := FilterContainers(containers, "postgres")
		if len(got) != 0 {
			t.Errorf("expected 0 matches, got %d", len(got))
		}
	})

	t.Run("empty service name", func(t *testing.T) {
		got := FilterContainers(containers, "")
		if len(got) != 0 {
			t.Errorf("expected 0 matches for empty name, got %d", len(got))
		}
	})

	t.Run("nil input", func(t *testing.T) {
		got := FilterContainers(nil, "mysql")
		if len(got) != 0 {
			t.Errorf("expected 0 matches for nil input, got %d", len(got))
		}
	})

	t.Run("does not modify input slice", func(t *testing.T) {
		input := []ContainerInfo{
			{ID: "aaa", Service: "mysql"},
			{ID: "bbb", Service: "redis"},
		}

		_ = FilterContainers(input, "mysql")

		if input[0].ID != "aaa" || input[0].Service != "mysql" {
			t.Error("input[0] was modified")
		}
		if input[1].ID != "bbb" || input[1].Service != "redis" {
			t.Error("input[1] was modified")
		}
	})
}
