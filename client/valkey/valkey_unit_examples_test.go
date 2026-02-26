package valkey

import (
	"context"
	"fmt"
)

// ExampleClient demonstrates creating a client from env config and calling ping.
func ExampleClient() {
	cfg := LoadConfigFromEnv()
	client, err := New(context.Background(), cfg)
	if err != nil {
		fmt.Println("cannot connect")
		return
	}
	defer client.Close()

	fmt.Println("ready")
}
