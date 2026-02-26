package valkey

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// Config holds the Valkey/Redis connection settings.
type Config struct {
	Address  string
	Password string
	Username string
	DB       int
}

const (
	defaultAddress  = "localhost:6379"
	defaultDB       = 0
	defaultUsername = ""
	defaultPassword = ""
)

// LoadConfigFromEnv loads connection config from environment variables.
//
// Supported env vars:
// - VALKEY_ADDR (default: localhost:6379)
// - VALKEY_USERNAME (default: "")
// - VALKEY_PASSWORD (default: "")
// - VALKEY_DB (default: 0)
func LoadConfigFromEnv() Config {
	cfg := Config{
		Address:  defaultAddress,
		Username: defaultUsername,
		Password: defaultPassword,
		DB:       defaultDB,
	}

	if v := os.Getenv("VALKEY_ADDR"); v != "" {
		cfg.Address = v
	}
	if v := os.Getenv("VALKEY_USERNAME"); v != "" {
		cfg.Username = v
	}
	if v := os.Getenv("VALKEY_PASSWORD"); v != "" {
		cfg.Password = v
	}
	if v := os.Getenv("VALKEY_DB"); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil {
			cfg.DB = parsed
		}
	}

	return cfg
}

// Client wraps redis.Client with a few convenience helpers.
type Client struct {
	*redis.Client
	cfg Config
}

// New builds a Valkey client from config.
func New(ctx context.Context, cfg Config) (*Client, error) {
	if cfg.Address == "" {
		cfg.Address = defaultAddress
	}

	c := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := c.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("valkey ping failed: %w", err)
	}

	return &Client{Client: c, cfg: cfg}, nil
}

// CloseClient closes underlying client.
func CloseClient(c *Client) error {
	if c == nil || c.Client == nil {
		return nil
	}
	return c.Close()
}
