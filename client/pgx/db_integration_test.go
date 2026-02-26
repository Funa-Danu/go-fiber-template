//go:build integration

package pgx

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func TestNewDB_WithTestContainer(t *testing.T) {
	if !hasDockerSocket() {
		t.Skip("integration test requires docker")
	}

	defer func() {
		if r := recover(); r != nil {
			t.Skipf("docker unavailable: %v", r)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	container, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
	)
	require.NoError(t, err)
	defer container.Terminate(ctx)

	mapped, err := container.MappedPort(ctx, "5432")
	require.NoError(t, err)

	cfg := DBConfig{
		Host:     "127.0.0.1",
		Port:     mapped.Port(),
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	}

	var dbConn interface{ Close() }
	var newErr error
	for i := 0; i < 10; i++ {
		dbConn, newErr = NewDB(ctx, cfg)
		if newErr == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	require.NoError(t, newErr)
	if dbConn != nil {
		dbConn.Close()
	}
}

func hasDockerSocket() bool {
	if _, err := os.Stat("/var/run/docker.sock"); err == nil {
		return true
	}
	if os.Getenv("DOCKER_HOST") != "" {
		return true
	}
	if runtimeDir := os.Getenv("XDG_RUNTIME_DIR"); runtimeDir != "" {
		if _, err := os.Stat(filepath.Join(runtimeDir, "docker.sock")); err == nil {
			return true
		}
	}
	return false
}
