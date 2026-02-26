package valkey

import "context"

// PingResult models the result of a ping operation.
type PingResult interface {
	Err() error
}

// ValkeyClient defines a minimal client interface used across the project.
type ValkeyClient interface {
	Ping(ctx context.Context) PingResult
	Close() error
}
