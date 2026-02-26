package store

import (
	"testing"
)

func TestGetService(t *testing.T) {
	t.Run("valid name returns true", func(t *testing.T) {
		s, ok := GetService("mysql")
		if !ok {
			t.Fatal("expected ok=true for mysql")
		}
		if s.Name != "mysql" {
			t.Errorf("Name = %q, want %q", s.Name, "mysql")
		}
	})

	t.Run("invalid name returns false", func(t *testing.T) {
		_, ok := GetService("nonexistent")
		if ok {
			t.Fatal("expected ok=false for nonexistent service")
		}
	})

	t.Run("deep copy protects Env", func(t *testing.T) {
		s1, _ := GetService("mysql")
		original := s1.Env[0].Value
		s1.Env[0].Value = "mutated"

		s2, _ := GetService("mysql")
		if s2.Env[0].Value != original {
			t.Errorf("Env mutation leaked: got %q, want %q", s2.Env[0].Value, original)
		}
	})

	t.Run("deep copy protects PortMappings", func(t *testing.T) {
		s1, _ := GetService("mysql")
		original := s1.PortMappings[0].HostPort
		s1.PortMappings[0].HostPort = 9999

		s2, _ := GetService("mysql")
		if s2.PortMappings[0].HostPort != original {
			t.Errorf("PortMappings mutation leaked: got %d, want %d", s2.PortMappings[0].HostPort, original)
		}
	})

	t.Run("deep copy protects Volumes", func(t *testing.T) {
		s1, _ := GetService("mysql")
		if len(s1.Volumes) == 0 {
			t.Skip("mysql has no volumes")
		}
		original := s1.Volumes[0].Name
		s1.Volumes[0].Name = "mutated"

		s2, _ := GetService("mysql")
		if s2.Volumes[0].Name != original {
			t.Errorf("Volumes mutation leaked: got %q, want %q", s2.Volumes[0].Name, original)
		}
	})

	t.Run("deep copy protects Command", func(t *testing.T) {
		s1, _ := GetService("mongo")
		if len(s1.Command) == 0 {
			t.Skip("mongo has no command")
		}
		original := s1.Command[0]
		s1.Command[0] = "mutated"

		s2, _ := GetService("mongo")
		if s2.Command[0] != original {
			t.Errorf("Command mutation leaked: got %q, want %q", s2.Command[0], original)
		}
	})

	t.Run("nil InsecureEnv stays nil", func(t *testing.T) {
		s, ok := GetService("redis")
		if !ok {
			t.Fatal("expected ok=true for redis")
		}
		if s.InsecureEnv != nil {
			t.Error("expected InsecureEnv to be nil for redis")
		}
	})
}

func TestListServiceNames(t *testing.T) {
	names := ListServiceNames()

	t.Run("returns expected count", func(t *testing.T) {
		if len(names) != 24 {
			t.Errorf("expected 24 services, got %d: %v", len(names), names)
		}
	})

	t.Run("sorted", func(t *testing.T) {
		for i := 1; i < len(names); i++ {
			if names[i] < names[i-1] {
				t.Errorf("not sorted: %q comes after %q", names[i], names[i-1])
			}
		}
	})

	t.Run("no duplicates", func(t *testing.T) {
		seen := make(map[string]bool)
		for _, name := range names {
			if seen[name] {
				t.Errorf("duplicate service name: %q", name)
			}
			seen[name] = true
		}
	})

	t.Run("every name resolves via GetService", func(t *testing.T) {
		for _, name := range names {
			if _, ok := GetService(name); !ok {
				t.Errorf("ListServiceNames includes %q but GetService returns false", name)
			}
		}
	})
}

func TestServiceStructuralInvariants(t *testing.T) {
	names := ListServiceNames()
	for _, name := range names {
		s, ok := GetService(name)
		if !ok {
			t.Errorf("GetService(%q) returned false", name)
			continue
		}

		t.Run(name, func(t *testing.T) {
			if s.Name == "" {
				t.Error("Name is empty")
			}
			if s.Image == "" {
				t.Error("Image is empty")
			}
			if s.Tag == "" {
				t.Error("Tag is empty")
			}

			// Registry key must match Service.Name
			if s.Name != name {
				t.Errorf("registry key %q != Service.Name %q", name, s.Name)
			}

			// At least one PortMapping
			if len(s.PortMappings) == 0 {
				t.Error("no PortMappings defined")
			}

			// Non-zero ContainerPort
			for i, pm := range s.PortMappings {
				if pm.ContainerPort == 0 {
					t.Errorf("PortMappings[%d].ContainerPort is zero", i)
				}
			}

			// Env keys non-empty
			for i, env := range s.Env {
				if env.Key == "" {
					t.Errorf("Env[%d].Key is empty", i)
				}
			}

			// InsecureEnv/InsecureInfo consistency:
			// Use ApplyInsecure to test support, because the deep copy
			// turns empty InsecureEnv ([]EnvVar{}) into nil.
			testCopy := s
			_, insecureErr := testCopy.ApplyInsecure()
			supportsInsecure := insecureErr == nil
			if supportsInsecure && s.InsecureInfo == "" {
				t.Error("service supports insecure mode but InsecureInfo is empty")
			}
			if !supportsInsecure && s.InsecureInfo != "" {
				t.Error("InsecureInfo is set but service does not support insecure mode")
			}

			// ContainerName doesn't panic
			_ = s.ContainerName()
		})
	}
}
