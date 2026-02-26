package types

import (
	"testing"
)

func TestContainerName(t *testing.T) {
	tests := []struct {
		name    string
		service Service
		want    string
	}{
		{
			name: "standard service",
			service: Service{
				Name: "mysql",
				Tag:  "latest",
				PortMappings: []PortMapping{
					{HostPort: 3306, ContainerPort: 3306},
				},
			},
			want: "tent-mysql-latest-3306",
		},
		{
			name: "custom tag and port",
			service: Service{
				Name: "postgres",
				Tag:  "16-alpine",
				PortMappings: []PortMapping{
					{HostPort: 5433, ContainerPort: 5432},
				},
			},
			want: "tent-postgres-16-alpine-5433",
		},
		{
			name: "uses first port only",
			service: Service{
				Name: "rabbitmq",
				Tag:  "latest",
				PortMappings: []PortMapping{
					{HostPort: 5672, ContainerPort: 5672},
					{HostPort: 15672, ContainerPort: 15672},
				},
			},
			want: "tent-rabbitmq-latest-5672",
		},
		{
			name: "zero host port",
			service: Service{
				Name: "redis",
				Tag:  "latest",
				PortMappings: []PortMapping{
					{HostPort: 0, ContainerPort: 6379},
				},
			},
			want: "tent-redis-latest-0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.service.ContainerName()
			if got != tt.want {
				t.Errorf("ContainerName() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestImageName(t *testing.T) {
	tests := []struct {
		name    string
		service Service
		want    string
	}{
		{
			name:    "simple image",
			service: Service{Image: "docker.io/redis", Tag: "latest"},
			want:    "docker.io/redis:latest",
		},
		{
			name:    "image with org prefix",
			service: Service{Image: "docker.io/library/mysql", Tag: "8.0"},
			want:    "docker.io/library/mysql:8.0",
		},
		{
			name:    "specific tag",
			service: Service{Image: "docker.io/postgres", Tag: "16-alpine"},
			want:    "docker.io/postgres:16-alpine",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.service.ImageName()
			if got != tt.want {
				t.Errorf("ImageName() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestApplyInsecure(t *testing.T) {
	t.Run("error when InsecureEnv is nil", func(t *testing.T) {
		s := Service{
			Name:        "redis",
			InsecureEnv: nil,
			Env: []EnvVar{
				{Key: "SOME_KEY", Value: "val"},
			},
		}

		info, err := s.ApplyInsecure()
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if info != "" {
			t.Errorf("expected empty info, got %q", info)
		}
		// Env should be unchanged
		if len(s.Env) != 1 || s.Env[0].Key != "SOME_KEY" {
			t.Error("Env was modified despite error")
		}
	})

	t.Run("success with non-empty InsecureEnv", func(t *testing.T) {
		s := Service{
			Name: "mysql",
			Env: []EnvVar{
				{Key: "MYSQL_ROOT_PASSWORD", Value: "secret"},
			},
			InsecureEnv: []EnvVar{
				{Key: "MYSQL_ALLOW_EMPTY_PASSWORD", Value: "yes"},
			},
			InsecureInfo: "username: root, password: (empty)",
		}

		info, err := s.ApplyInsecure()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if info != "username: root, password: (empty)" {
			t.Errorf("info = %q, want %q", info, "username: root, password: (empty)")
		}
		if len(s.Env) != 1 || s.Env[0].Key != "MYSQL_ALLOW_EMPTY_PASSWORD" {
			t.Error("Env not replaced with InsecureEnv")
		}
	})

	t.Run("success with empty InsecureEnv drops all auth", func(t *testing.T) {
		s := Service{
			Name: "mongo",
			Env: []EnvVar{
				{Key: "MONGO_INITDB_ROOT_USERNAME", Value: "admin"},
				{Key: "MONGO_INITDB_ROOT_PASSWORD", Value: "secret"},
			},
			InsecureEnv:  []EnvVar{},
			InsecureInfo: "authentication disabled",
		}

		info, err := s.ApplyInsecure()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if info != "authentication disabled" {
			t.Errorf("info = %q, want %q", info, "authentication disabled")
		}
		if len(s.Env) != 0 {
			t.Errorf("expected empty Env after insecure, got %d items", len(s.Env))
		}
	})
}
