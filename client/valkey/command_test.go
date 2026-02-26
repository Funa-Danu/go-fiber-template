package valkey

import (
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

type fakeStore struct {
	GetFn    func(key string) (string, error)
	SetFn    func(key, value string, expiration time.Duration) error
	DeleteFn func(key string) error
}

func (f fakeStore) Get(key string) (string, error) { return f.GetFn(key) }
func (f fakeStore) Set(key, value string, expiration time.Duration) error {
	return f.SetFn(key, value, expiration)
}
func (f fakeStore) Delete(key string) error { return f.DeleteFn(key) }
func (f fakeStore) Close() error            { return nil }

func TestGetItem(t *testing.T) {
	t.Run("valkey error", func(t *testing.T) {
		store := fakeStore{
			GetFn: func(key string) (string, error) {
				require.Equal(t, "missing", key)
				return "", errors.New("error")
			},
		}

		got, err := GetItem(store, "missing")
		require.Equal(t, "", got)
		require.ErrorContains(t, err, "valkey: get item")
	})

	t.Run("success", func(t *testing.T) {
		store := fakeStore{
			GetFn: func(key string) (string, error) {
				require.Equal(t, "key", key)
				return "value", nil
			},
		}

		got, err := GetItem(store, "key")
		require.NoError(t, err)
		require.Equal(t, "value", got)
	})
}

func TestSetItem(t *testing.T) {
	t.Run("set error", func(t *testing.T) {
		store := fakeStore{
			SetFn: func(key, value string, expiration time.Duration) error {
				require.Equal(t, "key", key)
				require.Equal(t, "value", value)
				require.Equal(t, time.Second, expiration)
				return errors.New("error")
			},
		}

		err := SetItem(store, "key", "value", time.Second)
		require.ErrorContains(t, err, "valkey: set item")
	})

	t.Run("success", func(t *testing.T) {
		store := fakeStore{
			SetFn: func(key, value string, expiration time.Duration) error {
				require.Equal(t, "key", key)
				require.Equal(t, "value", value)
				require.Equal(t, time.Second, expiration)
				return nil
			},
		}

		err := SetItem(store, "key", "value", time.Second)
		require.NoError(t, err)
	})
}

func TestDeleteItem(t *testing.T) {
	t.Run("valkey error", func(t *testing.T) {
		store := fakeStore{
			DeleteFn: func(key string) error {
				require.Equal(t, "key", key)
				return errors.New("error")
			},
		}

		err := DeleteItem(store, "key")
		require.ErrorContains(t, err, "valkey: delete item")
	})

	t.Run("success", func(t *testing.T) {
		store := fakeStore{
			DeleteFn: func(key string) error {
				require.Equal(t, "key", key)
				return nil
			},
		}

		err := DeleteItem(store, "key")
		require.NoError(t, err)
	})
}
