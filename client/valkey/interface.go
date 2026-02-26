package valkey

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -package valkey -self_package "go-fiber-template/client/valkey" -destination=./valkey_client_mock.go "go-fiber-template/client/valkey" ValkeyClient

import "time"

// ValkeyClient defines only the operations needed for cache tests.
type ValkeyClient interface {
	Get(key string) (string, error)
	Set(key, value string, expiration time.Duration) error
	Delete(key string) error
	Close() error
}
