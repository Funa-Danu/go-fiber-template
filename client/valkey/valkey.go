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

// Ping checks connectivity and satisfies ValkeyClient.
func (c *Client) Ping(ctx context.Context) PingResult {
	return c.Client.Ping(ctx)
}

// New builds a Valkey client from config.
func New(ctx context.Context, cfg Config) (ValkeyClient, error) {
	if cfg.Address == "" {
		cfg.Address = defaultAddress
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := redisClient.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("valkey ping failed: %w", err)
	}

	return &Client{Client: redisClient, cfg: cfg}, nil
}

// CloseClient closes underlying client.
func CloseClient(c ValkeyClient) error {
	if c == nil {
		return nil
	}
	return c.Close()
}
