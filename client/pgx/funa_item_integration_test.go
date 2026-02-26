//go:build integration

package pgx

import (
	"context"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed sql/test/truncate_funa_item.sql
var truncateFunaItemSQL string

func TestFunaItemRepository_CRUD_Integration(t *testing.T) {
	require.NotNil(t, globalTestDB)

	ctx := context.Background()
	_, err := globalTestDB.Exec(ctx, truncateFunaItemSQL)
	require.NoError(t, err)

	repo := NewFunaItemRepository(globalTestDB)
	created, err := repo.CreateFunaItem(ctx, "초기", "desc-1")
	require.NoError(t, err)
	require.NotZero(t, created.ID)

	got, err := repo.GetFunaItem(ctx, created.ID)
	require.NoError(t, err)
	require.Equal(t, created.ID, got.ID)
	require.Equal(t, "초기", got.Name)

	updated, err := repo.UpdateFunaItem(ctx, FunaItem{ID: created.ID, Name: "수정", Description: "desc-2"})
	require.NoError(t, err)
	require.Equal(t, int64(1), updated)

	items, err := repo.ListFunaItemNamesByPrefix(ctx, "수", 10, 0)
	require.NoError(t, err)
	require.Len(t, items, 1)
	require.Equal(t, "수정", items[0].Name)

	deleted, err := repo.DeleteFunaItem(ctx, created.ID)
	require.NoError(t, err)
	require.Equal(t, int64(1), deleted)
}
