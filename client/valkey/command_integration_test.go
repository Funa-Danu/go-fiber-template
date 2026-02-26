//go:build integration

package valkey

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/valkey-io/valkey-go"
)

var valkeyConfig *Config

func TestMain(m *testing.M) {
	if !hasDockerSocket() {
		m.Run()
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	container, cfg := MustNewValkeyContainer(ctx)
	if container != nil && cfg != nil {
		valkeyConfig = cfg
		defer func() {
			if err := container.Terminate(context.Background()); err != nil {
			}
		}()
	}

	m.Run()
}

func TestSetItem_WithIntegrated(t *testing.T) {
	if valkeyConfig == nil {
		t.Skip("integration test requires docker")
	}

	t.Run("success", func(t *testing.T) {
		store, err := NewClient(t.Context(), valkeyConfig)
		require.NoError(t, err)
		t.Cleanup(func() { _ = store.Close() })

		err = SetItem(store, "success", "value", time.Hour)
		require.NoError(t, err)

		result, err := GetItem(store, "success")
		require.NoError(t, err)
		require.Equal(t, "value", result)
	})
}

func TestGetItem_WithIntegrated(t *testing.T) {
	if valkeyConfig == nil {
		t.Skip("integration test requires docker")
	}

	t.Run("failed to get item", func(t *testing.T) {
		store, err := NewClient(t.Context(), valkeyConfig)
		require.NoError(t, err)
		t.Cleanup(func() { _ = store.Close() })

		result, err := GetItem(store, "missing")
		require.Equal(t, "", result)
		require.ErrorContains(t, err, "valkey: get item")
	})

	t.Run("success", func(t *testing.T) {
		store, err := NewClient(t.Context(), valkeyConfig)
		require.NoError(t, err)
		t.Cleanup(func() { _ = store.Close() })

		require.NoError(t, SetItem(store, "success", "value", time.Minute))
		result, err := GetItem(store, "success")
		require.NoError(t, err)
		require.Equal(t, "value", result)
	})
}

func TestDeleteItem_WithIntegrated(t *testing.T) {
	if valkeyConfig == nil {
		t.Skip("integration test requires docker")
	}

	t.Run("success", func(t *testing.T) {
		store, err := NewClient(t.Context(), valkeyConfig)
		require.NoError(t, err)
		t.Cleanup(func() { _ = store.Close() })

		require.NoError(t, SetItem(store, "to-delete", "value", time.Minute))
		require.NoError(t, DeleteItem(store, "to-delete"))

		value, err := GetItem(store, "to-delete")
		require.Error(t, err)
		require.Equal(t, "", value)
	})
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

var _ = valkey.Nil
