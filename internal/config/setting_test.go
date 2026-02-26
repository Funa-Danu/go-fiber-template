package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSetting_Defaults(t *testing.T) {
	t.Setenv("SERVICE_NAME", "")
	t.Setenv("ENV", "")
	t.Setenv("PORT", "")
	t.Setenv("VALKEY_ADDR", "")
	t.Setenv("VALKEY_DB", "")

	s := NewSetting()

	require.Equal(t, "go-fiber-template", s.ServiceName)
	require.Equal(t, "local", s.Env)
	require.Equal(t, "3000", s.Port)
	require.Equal(t, "localhost:6379", s.ValkeyAddress)
	require.Equal(t, 0, s.ValkeyDB)
}

func TestNewSetting_CustomValues(t *testing.T) {
	t.Setenv("SERVICE_NAME", "demo-service")
	t.Setenv("ENV", "dev")
	t.Setenv("PORT", "8081")
	t.Setenv("VALKEY_ADDR", "cache:6379")
	t.Setenv("VALKEY_DB", "2")

	s := NewSetting()

	require.Equal(t, "demo-service", s.ServiceName)
	require.Equal(t, "dev", s.Env)
	require.Equal(t, "8081", s.Port)
	require.Equal(t, "cache:6379", s.ValkeyAddress)
	require.Equal(t, 2, s.ValkeyDB)
}

func TestConfig_New_Validate(t *testing.T) {
	cfg, err := New(Setting{Port: "abc"})
	require.Error(t, err)

	cfg, err = New(Setting{Port: "3000", ValkeyDB: 0, ServiceName: "svc", Env: "local", ValkeyAddress: "localhost:6379"})
	require.NoError(t, err)
	require.NotNil(t, cfg)
}

func TestMissingEnvError(t *testing.T) {
	err := (&MissingEnvError{Key: "TEST_KEY"}).Error()
	require.Contains(t, err, "TEST_KEY")
	require.Equal(t, "", os.Getenv("NEVER_SET"))
}
