package store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type UserStore interface {
	// Lookups
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByEmailCI(ctx context.Context, email string) (*User, error)     // case-insensitive
	GetByUsername(ctx context.Context, username string) (*User, error) // case-insensitive

	// Mutations
	Create(ctx context.Context, u *User) error
	UpdatePassword(ctx context.Context, userID uuid.UUID, newHash string, changedAt time.Time) error
	UpdateRole(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error
}

type RoleStore interface {
	GetByName(ctx context.Context, name string) (*Role, error)
}

type SessionStore interface {
	Create(ctx context.Context, s *Session) error
	GetByID(ctx context.Context, sessionID string) (*Session, error)
	Touch(ctx context.Context, sessionID string, lastSeenAt time.Time) error
	Revoke(ctx context.Context, sessionID string, revokedAt time.Time) error

	RevokeAllForUser(ctx context.Context, userID uuid.UUID, revokedAt time.Time) error
	RevokeAllForUserExcept(ctx context.Context, userID uuid.UUID, exceptSessionID string, revokedAt time.Time) error

	// Optional nice-to-have
	ListForUser(ctx context.Context, userID uuid.UUID, limit int) ([]Session, error)
}

type PasswordResetStore interface {
	// Create a new reset token; caller ensures old ones are invalidated if desired.
	Create(ctx context.Context, t *PasswordResetToken) error

	// Invalidate/expire old tokens for a user (optional but recommended)
	InvalidateAllForUser(ctx context.Context, userID uuid.UUID, now time.Time) error

	// Lookup by hash, validate in service (expiry/used) or here; your choice.
	GetByHash(ctx context.Context, tokenHash []byte) (*PasswordResetToken, error)
	MarkUsed(ctx context.Context, tokenID uuid.UUID, usedAt time.Time) error
}

type AuditStore interface {
	Add(ctx context.Context, ev *AuditEvent) error
}

type Store interface {
	Users() UserStore
	Roles() RoleStore
	Sessions() SessionStore
	PasswordResets() PasswordResetStore
	Audit() AuditStore
}
