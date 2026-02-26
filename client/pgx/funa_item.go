package pgx

import (
	"context"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed sql/funa_item_create.sql
var createFunaItemSQL string

//go:embed sql/funa_item_get.sql
var getFunaItemSQL string

//go:embed sql/funa_item_update.sql
var updateFunaItemSQL string

//go:embed sql/funa_item_delete.sql
var deleteFunaItemSQL string

//go:embed sql/funa_item_list_by_name_like.sql
var listFunaItemByNameLikeSQL string

//go:embed sql/funa_item_list_by_name_like_with_escape.sql
var listFunaItemByNameLikeWithEscapeSQL string

// FunaItem represents the sample table record used by the pgx CRUD example.
type FunaItem struct {
	ID          int64
	Name        string
	Description string
}

// FunaItemRepository is a minimal data access example for funa_item.
type FunaItemRepository struct {
	db Client
}

// NewFunaItemRepository returns a repository that uses the provided pgx client.
func NewFunaItemRepository(db Client) *FunaItemRepository {
	return &FunaItemRepository{db: db}
}

// CreateFunaItem inserts one row and returns the created record.
func (r *FunaItemRepository) CreateFunaItem(ctx context.Context, name, description string) (FunaItem, error) {
	item, err := mapFunaItem(r.db.QueryRow(
		ctx,
		createFunaItemSQL,
		name,
		description,
	))
	if err != nil {
		return FunaItem{}, fmt.Errorf("funa_item create: %w", err)
	}
	return item, nil
}

// GetFunaItem retrieves one row by ID.
func (r *FunaItemRepository) GetFunaItem(ctx context.Context, id int64) (FunaItem, error) {
	item, err := mapFunaItem(r.db.QueryRow(
		ctx,
		getFunaItemSQL,
		id,
	))
	if err != nil {
		return FunaItem{}, fmt.Errorf("funa_item get: %w", err)
	}
	return item, nil
}

// UpdateFunaItem updates name/description and returns affected row count.
func (r *FunaItemRepository) UpdateFunaItem(ctx context.Context, item FunaItem) (int64, error) {
	tag, err := r.db.Exec(
		ctx,
		updateFunaItemSQL,
		item.Name,
		item.Description,
		item.ID,
	)
	if err != nil {
		return 0, fmt.Errorf("funa_item update: %w", err)
	}
	return int64(tag.RowsAffected()), nil
}

// DeleteFunaItem deletes by ID and returns affected row count.
func (r *FunaItemRepository) DeleteFunaItem(ctx context.Context, id int64) (int64, error) {
	tag, err := r.db.Exec(
		ctx,
		deleteFunaItemSQL,
		id,
	)
	if err != nil {
		return 0, fmt.Errorf("funa_item delete: %w", err)
	}
	return int64(tag.RowsAffected()), nil
}

// ListFunaItemNamesByPrefix returns rows where name contains keyword, paging by limit/offset.
func (r *FunaItemRepository) ListFunaItemNamesByPrefix(ctx context.Context, namePrefix string, limit, offset int32) ([]FunaItem, error) {
	rows, err := r.db.Query(
		ctx,
		listFunaItemByNameLikeSQL,
		"%"+namePrefix+"%",
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("funa_item query: %w", err)
	}
	defer rows.Close()

	results := make([]FunaItem, 0)
	for rows.Next() {
		item, err := mapFunaItemRows(rows)
		if err != nil {
			return nil, fmt.Errorf("funa_item scan: %w", err)
		}
		results = append(results, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("funa_item query: %w", err)
	}
	return results, nil
}

func mapFunaItem(rowQueryRow pgxLikeRow) (FunaItem, error) {
	var item FunaItem
	if err := rowQueryRow.Scan(&item.ID, &item.Name, &item.Description); err != nil {
		return FunaItem{}, err
	}
	return item, nil
}

func mapFunaItemRows(rows pgxLikeRows) (FunaItem, error) {
	var item FunaItem
	if err := rows.Scan(&item.ID, &item.Name, &item.Description); err != nil {
		return FunaItem{}, err
	}
	return item, nil
}

// Small interfaces to keep test mocks easier if needed.
type pgxLikeRow interface {
	Scan(dest ...any) error
}

type pgxLikeRows interface {
	Close()
	Err() error
	Next() bool
	Scan(dest ...any) error
}

func sanitizePrefix(prefix string) string {
	trimmed := strings.TrimSpace(prefix)
	if trimmed == "" {
		return "%"
	}
	if strings.Contains(trimmed, "%") || strings.Contains(trimmed, "_") {
		trimmed = strings.ReplaceAll(trimmed, "%", "\\%")
		trimmed = strings.ReplaceAll(trimmed, "_", "\\_")
	}
	return trimmed
}

func normalizePattern(p string) string {
	if p == "" {
		return "%"
	}
	return "%" + sanitizePrefix(p) + "%"
}

func (r *FunaItemRepository) ListFunaItemNamesByPrefixWithEscape(ctx context.Context, namePrefix string, limit, offset int32) ([]FunaItem, error) {
	pattern := normalizePattern(namePrefix)
	rows, err := r.db.Query(
		ctx,
		listFunaItemByNameLikeWithEscapeSQL,
		pattern,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("funa_item query: %w", err)
	}
	defer rows.Close()

	results := make([]FunaItem, 0)
	for rows.Next() {
		item, err := mapFunaItemRows(rows)
		if err != nil {
			return nil, fmt.Errorf("funa_item scan: %w", err)
		}
		results = append(results, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("funa_item query: %w", err)
	}
	return results, nil
}
