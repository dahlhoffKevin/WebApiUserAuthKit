package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/dahlhoffKevin/WebApiAuthKit/store"
	"github.com/google/uuid"
)

type passwordResetStore struct {
	db *sql.DB
}

// Create a new reset token; caller ensures old ones are invalidated if desired.
func (p *passwordResetStore) Create(ctx context.Context, t *store.PasswordResetToken) error {
	return nil
}

// Invalidate/expire old tokens for a user (optional but recommended)
func (p *passwordResetStore) InvalidateAllForUser(ctx context.Context, userID uuid.UUID, now time.Time) error {
	return nil
}

// Lookup by hash, validate in service (expiry/used) or here; your choice.
func (p *passwordResetStore) GetByHash(ctx context.Context, tokenHash []byte) (*store.PasswordResetToken, error) {
	return nil, nil
}

func (p *passwordResetStore) MarkUsed(ctx context.Context, tokenID uuid.UUID, usedAt time.Time) error {
	return nil
}
