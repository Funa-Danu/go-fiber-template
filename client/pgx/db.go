package pgx

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:generate go run github.com/golang/mock/mockgen -source=db.go -destination=pgx_client_mock.go -package=pgx

// Client defines minimum database operations used across templates.
type Client interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

// DBConfig holds pgx connection settings.
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// NewDB opens a pgx pool and validates connectivity.
func NewDB(ctx context.Context, cfg DBConfig) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, buildConnString(cfg))
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("pgx ping failed: %w", err)
	}

	return pool, nil
}

func buildConnString(cfg DBConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
}
