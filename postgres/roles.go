package postgres

import (
	"context"
	"database/sql"

	"github.com/dahlhoffKevin/WebApiAuthKit/store"
)

type roleStore struct {
	db *sql.DB
}

func (r *roleStore) GetByName(ctx context.Context, name string) (*store.Role, error) {
	return nil, nil
}
