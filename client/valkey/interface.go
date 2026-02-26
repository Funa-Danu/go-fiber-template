package valkey

//go:generate mockgen -source=interface.go -destination=./mocks/mock_valkey.go -package=mockclient

import "time"

// ValkeyClient defines only the operations needed for cache tests.
type ValkeyClient interface {
	Get(key string) (string, error)
	Set(key, value string, expiration time.Duration) error
	Delete(key string) error
	Close() error
}
