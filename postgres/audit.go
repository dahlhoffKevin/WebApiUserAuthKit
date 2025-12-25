package postgres

import (
	"context"
	"database/sql"

	"github.com/dahlhoffKevin/WebApiAuthKit/store"
)

type auditStore struct {
	db *sql.DB
}

func (a *auditStore) Add(ctx context.Context, ev *store.AuditEvent) error {
	return nil
}
