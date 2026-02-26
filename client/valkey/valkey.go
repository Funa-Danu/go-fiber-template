package valkey

import (
	"context"
	"time"

	"github.com/pkg/errors"
	valkey_module "github.com/valkey-io/valkey-go"
)

// Config wraps valkey client options.
type Config struct {
	Address string
	DB      int
}

// ValkeyStore represents a concrete valkey client.
type ValkeyStore struct {
	ctx    context.Context
	client valkey_module.Client
}

// NewClient creates a valkey client and validates connectivity.
func NewClient(ctx context.Context, cfg *Config) (ValkeyClient, error) {
	if cfg == nil {
		cfg = &Config{}
	}
	if cfg.Address == "" {
		cfg.Address = "localhost:6379"
	}

	client, err := valkey_module.NewClient(valkey_module.ClientOption{
		InitAddress: []string{cfg.Address},
		SelectDB:    cfg.DB,
	})
	if err != nil {
		return nil, errors.Wrap(err, "valkey: new client")
	}

	if err := client.Do(ctx, client.B().Ping().Build()).Error(); err != nil {
		client.Close()
		return nil, errors.Wrap(err, "valkey: ping")
	}

	return &ValkeyStore{ctx: ctx, client: client}, nil
}

func (v *ValkeyStore) Get(key string) (string, error) {
	return v.client.Do(v.ctx, v.client.B().Get().Key(key).Build()).ToString()
}

func (v *ValkeyStore) Set(key, value string, expiration time.Duration) error {
	return v.client.Do(v.ctx, v.client.B().Set().Key(key).Value(value).ExSeconds(int64(expiration.Seconds())).Build()).Error()
}

func (v *ValkeyStore) Delete(key string) error {
	return v.client.Do(v.ctx, v.client.B().Del().Key(key).Build()).Error()
}

func (v *ValkeyStore) Close() error {
	v.client.Close()
	return nil
}

// GetItem gets string by key.
func GetItem(client ValkeyClient, key string) (string, error) {
	v, err := client.Get(key)
	if err != nil {
		return "", errors.Wrap(err, "valkey: get item")
	}
	return v, nil
}

// SetItem sets string with expiration.
func SetItem(client ValkeyClient, key, value string, expiration time.Duration) error {
	if err := client.Set(key, value, expiration); err != nil {
		return errors.Wrap(err, "valkey: set item")
	}
	return nil
}

// DeleteItem deletes key.
func DeleteItem(client ValkeyClient, key string) error {
	if err := client.Delete(key); err != nil {
		return errors.Wrap(err, "valkey: delete item")
	}
	return nil
}
