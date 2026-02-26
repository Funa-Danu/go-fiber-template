package valkey_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"go-fiber-template/client/valkey"
	"go-fiber-template/client/valkey/mocks"
)

// fakePingResult is a tiny test double for PingResult.
type fakePingResult struct {
	err error
}

func (f fakePingResult) Err() error { return f.err }

func TestCheckAlive_SucceedsWithMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mockclient.NewMockValkeyClient(ctrl)
	mockClient.EXPECT().Ping(context.Background()).Return(fakePingResult{err: nil})

	err := valkey.CheckAlive(context.Background(), mockClient)
	require.NoError(t, err)
}

func TestCheckAlive_ReturnsErrorFromPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mockclient.NewMockValkeyClient(ctrl)
	mockClient.EXPECT().Ping(context.Background()).Return(fakePingResult{err: context.Canceled})

	err := valkey.CheckAlive(context.Background(), mockClient)
	require.Equal(t, context.Canceled, err)
}
