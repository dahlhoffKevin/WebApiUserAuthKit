package postgres

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/dahlhoffKevin/WebApiAuthKit/store"
	"github.com/google/uuid"
)

type userStore struct {
	db *sql.DB
}

func (u *userStore) GetByID(ctx context.Context, id uuid.UUID) (*store.User, error) {
	const q = `
		select id, firstname, lastname, username, email, password_hash,
			password_changed_at, roleid, created_at
		from authkit.users
		where id = $1
	`

	var user store.User
	err := u.db.QueryRowContext(ctx, q, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.PasswordChangedAt,
		&user.RoleID,
		&user.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userStore) GetByEmailCI(ctx context.Context, email string) (*store.User, error) {
	const q = `
		select id, firstname, lastname, username, email, password_hash,
			password_changed_at, roleid, created_at
		from authkit.users
		where lower(email) = lower($1)
	`
	email = strings.TrimSpace(email)

	var user store.User
	err := u.db.QueryRowContext(ctx, q, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.PasswordChangedAt,
		&user.RoleID,
		&user.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userStore) GetByUsername(ctx context.Context, username string) (*store.User, error) {
	const q = `
		select id, firstname, lastname, username, email, password_hash,
			password_changed_at, roleid, created_at
		from authkit.users
		where lower(username) = lower($1)
	`

	username = strings.TrimSpace(username)

	var user store.User
	err := u.db.QueryRowContext(ctx, q, username).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.PasswordChangedAt,
		&user.RoleID,
		&user.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userStore) Create(ctx context.Context, user *store.User) error {
	return nil
}

func (u *userStore) UpdatePassword(ctx context.Context, userID uuid.UUID, newHash string, changedAt time.Time) error {
	return nil
}

func (u *userStore) UpdateRole(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) error {
	return nil
}
