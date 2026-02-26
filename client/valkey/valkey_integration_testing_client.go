package valkey

import (
	"context"
	"log"

	valkey_module "github.com/testcontainers/testcontainers-go/modules/valkey"
)

func MustNewValkeyContainer(ctx context.Context) (*valkey_module.ValkeyContainer, *Config) {
	container, err := valkey_module.Run(ctx, "valkey/valkey:8.1.0")
	if err != nil {
		return nil, nil
	}

	endpoint, err := container.Endpoint(ctx, "")
	if err != nil {
		log.Printf("read valkey endpoint failed: %v", err)
		if err := container.Terminate(context.Background()); err != nil {
			log.Printf("terminate valkey container after endpoint error: %v", err)
		}
		return nil, nil
	}

	return container, &Config{Address: endpoint, DB: 0}
}
