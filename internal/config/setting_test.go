package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSetting_Defaults(t *testing.T) {
	t.Setenv("SERVICE_NAME", "")
	t.Setenv("ENV", "")
	t.Setenv("PORT", "")
	t.Setenv("VALKEY_ADDR", "")
	t.Setenv("VALKEY_DB", "")
	t.Setenv("PG_HOST", "")
	t.Setenv("PG_PORT", "")
	t.Setenv("PG_USER", "")
	t.Setenv("PG_PASSWORD", "")
	t.Setenv("PG_DATABASE", "")

	s := NewSetting()

	require.Equal(t, "go-fiber-template", s.ServiceName)
	require.Equal(t, "local", s.Env)
	require.Equal(t, "3000", s.Port)
	require.Equal(t, "localhost:6379", s.ValkeyAddress)
	require.Equal(t, 0, s.ValkeyDB)
	require.Equal(t, "localhost", s.PostgresHost)
	require.Equal(t, "5432", s.PostgresPort)
	require.Equal(t, "postgres", s.PostgresUser)
	require.Equal(t, "postgres", s.PostgresPassword)
	require.Equal(t, "postgres", s.PostgresDatabase)
}

func TestNewSetting_CustomValues(t *testing.T) {
	t.Setenv("SERVICE_NAME", "demo-service")
	t.Setenv("ENV", "dev")
	t.Setenv("PORT", "8081")
	t.Setenv("VALKEY_ADDR", "cache:6379")
	t.Setenv("VALKEY_DB", "2")
	t.Setenv("PG_HOST", "pg")
	t.Setenv("PG_PORT", "5433")
	t.Setenv("PG_USER", "user")
	t.Setenv("PG_PASSWORD", "pwd")
	t.Setenv("PG_DATABASE", "appdb")

	s := NewSetting()

	require.Equal(t, "demo-service", s.ServiceName)
	require.Equal(t, "dev", s.Env)
	require.Equal(t, "8081", s.Port)
	require.Equal(t, "cache:6379", s.ValkeyAddress)
	require.Equal(t, 2, s.ValkeyDB)
	require.Equal(t, "pg", s.PostgresHost)
	require.Equal(t, "5433", s.PostgresPort)
	require.Equal(t, "user", s.PostgresUser)
	require.Equal(t, "pwd", s.PostgresPassword)
	require.Equal(t, "appdb", s.PostgresDatabase)
}

func TestConfig_New_Validate(t *testing.T) {
	cfg, err := New(Setting{Port: "abc"})
	require.Error(t, err)

	cfg, err = New(Setting{Port: "3000", ValkeyDB: 0, ServiceName: "svc", Env: "local", ValkeyAddress: "localhost:6379", PostgresHost: "localhost", PostgresPort: "5432"})
	require.NoError(t, err)
	require.NotNil(t, cfg)
}

func TestMissingEnvError(t *testing.T) {
	err := (&MissingEnvError{Key: "TEST_KEY"}).Error()
	require.Contains(t, err, "TEST_KEY")
}
