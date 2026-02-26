package valkey

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckAlive_NilClient(t *testing.T) {
	err := CheckAlive(context.Background(), nil)
	require.Equal(t, ErrNilClient, err)
}
