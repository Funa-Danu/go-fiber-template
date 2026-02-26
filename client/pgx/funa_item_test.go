package pgx

import (
	"context"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/require"
)

func TestMapFunaItem(t *testing.T) {
	row := fakeRow{values: []any{int64(7), "name", "desc"}}
	item, err := mapFunaItem(row)
	require.NoError(t, err)
	require.Equal(t, int64(7), item.ID)
	require.Equal(t, "name", item.Name)
	require.Equal(t, "desc", item.Description)
}

func TestFunaItemRepository_CreateAndDelete(t *testing.T) {
	ctx := context.Background()

	c := &fakePgxClient{
		queryRowFn: func(ctx context.Context, sql string, args ...any) pgx.Row {
			require.Equal(t, "INSERT INTO funa_item (name, description) VALUES ($1, $2) RETURNING id, name, description", strings.Join(strings.Fields(sql), " "))
			require.Equal(t, "n", args[0])
			require.Equal(t, "d", args[1])
			return fakeRow{values: []any{int64(1), "n", "d"}}
		},
		execFn: func(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
			require.Equal(t, "DELETE FROM funa_item WHERE id = $1", strings.Join(strings.Fields(sql), " "))
			require.Equal(t, int64(1), args[0])
			return pgconn.NewCommandTag("DELETE 1"), nil
		},
	}

	repo := NewFunaItemRepository(c)
	item, err := repo.CreateFunaItem(ctx, "n", "d")
	require.NoError(t, err)
	require.Equal(t, int64(1), item.ID)

	deleted, err := repo.DeleteFunaItem(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, int64(1), deleted)
}

// fake implementations for repository unit test
type fakePgxClient struct {
	queryRowFn func(ctx context.Context, sql string, args ...any) pgx.Row
	execFn     func(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

func (f *fakePgxClient) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return f.queryRowFn(ctx, sql, args...)
}

func (f *fakePgxClient) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return nil, nil
}

func (f *fakePgxClient) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return f.execFn(ctx, sql, args...)
}

// fakeRow implements pgx.Row.
type fakeRow struct {
	values []any
	err    error
}

func (r fakeRow) Scan(dest ...any) error {
	if r.err != nil {
		return r.err
	}
	for i, dst := range dest {
		v := r.values[i]
		switch target := dst.(type) {
		case *int64:
			*target = v.(int64)
		case *string:
			*target = v.(string)
		default:
			return nil
		}
	}
	return nil
}
