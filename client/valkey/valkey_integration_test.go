//go:build integration
// +build integration

package valkey

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestIntegration_NewAndPingWithContainer(t *testing.T) {
	if !hasDockerSocket() {
		t.Skip("integration-test requires Docker runtime")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	req := testcontainers.ContainerRequest{
		Image:        "valkey/valkey:7-alpine",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForListeningPort("6379/tcp"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	defer func() {
		_ = container.Terminate(context.Background())
	}()

	host, err := container.Host(ctx)
	require.NoError(t, err)
	port, err := container.MappedPort(ctx, "6379/tcp")
	require.NoError(t, err)
	addr := fmt.Sprintf("%s:%s", host, port.Port())

	cfg := Config{Address: addr}
	c, err := New(ctx, cfg)
	require.NoError(t, err)
	require.NotNil(t, c)
	defer func() { _ = CloseClient(c) }()

	err = c.Ping(ctx).Err()
	require.NoError(t, err)
}

func hasDockerSocket() bool {
	if _, err := os.Stat("/var/run/docker.sock"); err == nil {
		return true
	}

	if r := os.Getenv("XDG_RUNTIME_DIR"); r != "" {
		if _, err := os.Stat(filepath.Join(r, "docker.sock")); err == nil {
			return true
		}
	}

	return os.Getenv("DOCKER_HOST") != ""
}
