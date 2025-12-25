package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/dahlhoffKevin/WebApiAuthKit/store"
	"github.com/google/uuid"
)

type sessionStore struct {
	db *sql.DB
}

func (s *sessionStore) Create(ctx context.Context, session *store.Session) error {
	const q = `
		insert into authkit.usersessions
		(id, userid, created_at, expires_at, ip, user_agent)
		values ($1, $2, $3, $4, $5, $6)
	`
	_, err := s.db.ExecContext(
		ctx,
		q,
		session.ID,
		session.UserID,
		session.CreatedAt,
		session.ExpiresAt,
		session.IP,
		session.UserAgent,
	)
	return err
}

func (s *sessionStore) GetByID(ctx context.Context, id string) (*store.Session, error) {
	const q = `
		select id, userid, created_at, expires_at,
		       last_seen_at, revoked_at, ip, user_agent
		from authkit.usersessions
		where id = $1
	`
	var session store.Session
	err := s.db.QueryRowContext(ctx, q, id).Scan(
		&session.ID,
		&session.UserID,
		&session.CreatedAt,
		&session.ExpiresAt,
		&session.LastSeenAt,
		&session.RevokedAt,
		&session.IP,
		&session.UserAgent,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *sessionStore) Touch(ctx context.Context, id string, t time.Time) error {
	const q = `
		update authkit.usersessions
		set last_seen_at = $2
		where id = $1
		  and revoked_at is null
	`
	_, err := s.db.ExecContext(ctx, q, id, t)
	return err
}

func (s *sessionStore) Revoke(ctx context.Context, sessionID string, revokedAt time.Time) error {
	const q = `
		update authkit.usersessions
		set revoked_at = $2
		WHERE id = $1 and revoked_at is null
	`

	_, err := s.db.ExecContext(ctx, q, sessionID, revokedAt)
	return err
}

func (s *sessionStore) RevokeAllForUser(ctx context.Context, userID uuid.UUID, revokedAt time.Time) error {
	return nil
}

func (s *sessionStore) RevokeAllForUserExcept(ctx context.Context, userID uuid.UUID, exceptSessionID string, revokedAt time.Time) error {
	return nil
}

func (s *sessionStore) ListForUser(ctx context.Context, userID uuid.UUID, limit int) ([]store.Session, error) {
	return nil, nil
}
