package postgres

import (
	"database/sql"

	"github.com/dahlhoffKevin/WebApiAuthKit/store"
)

// Aggregrator
type Store struct {
	db *sql.DB

	users          store.UserStore
	roles          store.RoleStore
	sessions       store.SessionStore
	passwordResets store.PasswordResetStore
	audit          store.AuditStore
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:             db,
		users:          &userStore{db: db},
		roles:          &roleStore{db: db},
		sessions:       &sessionStore{db: db},
		passwordResets: &passwordResetStore{db: db},
		audit:          &auditStore{db: db},
	}
}

// --- store.Store interface ---

func (s *Store) Users() store.UserStore {
	return s.users
}

func (s *Store) Roles() store.RoleStore {
	return s.roles
}

func (s *Store) Sessions() store.SessionStore {
	return s.sessions
}

func (s *Store) PasswordResets() store.PasswordResetStore {
	return s.passwordResets
}

func (s *Store) Audit() store.AuditStore {
	return s.audit
}
