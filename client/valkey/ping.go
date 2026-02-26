package valkey

import "context"

// CheckAlive validates connectivity by pinging the client.
func CheckAlive(ctx context.Context, client ValkeyClient) error {
	if client == nil {
		return ErrNilClient
	}
	return client.Ping(ctx).Err()
}
