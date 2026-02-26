//go:build integration

package pgx

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

var (
	globalTestDB *pgxpool.Pool
)

func TestMain(m *testing.M) {
	if !hasDockerSocket() {
		fmt.Println("integration test requires docker")
		os.Exit(0)
	}

	ctx := context.Background()

	container, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		postgres.WithInitScripts(filepath.Join(".", "sql", "funa_item_schema.sql")),
	)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			panic(err)
		}
	}()

	mapped, err := container.MappedPort(ctx, "5432")
	if err != nil {
		panic(err)
	}

	var db *pgxpool.Pool
	for i := 0; i < 10; i++ {
		db, err = NewDB(ctx, DBConfig{
			Host:     "127.0.0.1",
			Port:     mapped.Port(),
			User:     "postgres",
			Password: "postgres",
			Database: "postgres",
		})
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		panic(err)
	}
	globalTestDB = db
	defer db.Close()

	os.Exit(m.Run())
}

func TestNewDB_WithTestContainer(t *testing.T) {
	require.NotNil(t, globalTestDB)
	require.NoError(t, globalTestDB.Ping(context.Background()))
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
