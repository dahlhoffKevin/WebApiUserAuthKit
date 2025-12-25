package store

import (
	"encoding/json"
	"net"
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID   uuid.UUID
	name string
}

type User struct {
	ID                uuid.UUID
	FirstName         *string
	LastName          *string
	Username          string
	Email             string
	PasswordHash      string
	PasswordChangedAt *time.Time
	RoleID            uuid.UUID
	CreatedAt         time.Time
}

type Session struct {
	// ID is the cookie session value (base64url string)
	ID         string
	UserID     uuid.UUID
	CreatedAt  time.Time
	ExpiresAt  time.Time
	LastSeenAt *time.Time
	RevokedAt  *time.Time
	IP         *net.IP
	UserAgent  *string
}

type PasswordResetToken struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	TokenHash []byte
	CreatedAt time.Time
	ExpiresAt time.Time
	UsedAt    *time.Time
}

type AuditEventType string

const (
	AuditLoginSuccess    AuditEventType = "login_success"
	AuditLoginFailed     AuditEventType = "login_failed"
	AuditResetRequested  AuditEventType = "reset_requested"
	AuditResetCompleted  AuditEventType = "reset_completed"
	AuditLogout          AuditEventType = "logout"
	AuditSessionRevoed   AuditEventType = "session_revoked"
	AuditPasswordChanged AuditEventType = "password_changed"
)

type AuditEvent struct {
	ID        uuid.UUID
	UserID    *uuid.UUID
	EventType AuditEventType
	CreatedAt time.Time
	IP        *net.IP
	UserAgent *string
	Metadata  json.RawMessage
}
