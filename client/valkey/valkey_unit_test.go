//go:build !integration
// +build !integration

package valkey

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfigFromEnv_Defaults(t *testing.T) {
	// isolate env
	t.Setenv("VALKEY_ADDR", "")
	t.Setenv("VALKEY_USERNAME", "")
	t.Setenv("VALKEY_PASSWORD", "")
	t.Setenv("VALKEY_DB", "")

	cfg := LoadConfigFromEnv()
	require.Equal(t, "localhost:6379", cfg.Address)
	require.Equal(t, "", cfg.Username)
	require.Equal(t, "", cfg.Password)
	require.Equal(t, 0, cfg.DB)
}

func TestLoadConfigFromEnv_Overrides(t *testing.T) {
	t.Setenv("VALKEY_ADDR", "valkey:6379")
	t.Setenv("VALKEY_USERNAME", "u")
	t.Setenv("VALKEY_PASSWORD", "p")
	t.Setenv("VALKEY_DB", "2")

	cfg := LoadConfigFromEnv()
	require.Equal(t, "valkey:6379", cfg.Address)
	require.Equal(t, "u", cfg.Username)
	require.Equal(t, "p", cfg.Password)
	require.Equal(t, 2, cfg.DB)
}

func TestNew_InvalidDBFallsBackToZeroAndFailsWithoutServer(t *testing.T) {
	cfg := LoadConfigFromEnv()
	cfg.Address = "127.0.0.1:0000"

	_, err := New(context.Background(), cfg)
	require.Error(t, err)
}

func TestLoadConfigFromEnv_DBInvalid(t *testing.T) {
	old := os.Getenv("VALKEY_DB")
	t.Setenv("VALKEY_DB", "not-a-number")

	cfg := LoadConfigFromEnv()
	require.Equal(t, 0, cfg.DB)

	if old == "" {
		t.Setenv("VALKEY_DB", "")
	} else {
		t.Setenv("VALKEY_DB", old)
	}
}
