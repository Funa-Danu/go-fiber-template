package pgx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildConnString(t *testing.T) {
	cfg := DBConfig{Host: "localhost", Port: "5432", User: "postgres", Password: "secret", Database: "app"}
	conn := buildConnString(cfg)
	require.Equal(t, "host=localhost port=5432 user=postgres password=secret dbname=app sslmode=disable", conn)
}

func TestNewDB_InvalidPort(t *testing.T) {
	ctx := context.Background()
	_, err := NewDB(ctx, DBConfig{Host: "x", Port: "-1", User: "u", Password: "p", Database: "d"})
	require.Error(t, err)
}
