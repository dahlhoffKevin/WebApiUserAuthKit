package authkit

import "github.com/dahlhoffKevin/WebApiAuthKit/store"

type Authkit struct {
	cfg   Config
	store store.Store
}

func New(cfg Config, store store.Store) *Authkit {
	if cfg.SessionCookieName == "" {
		cfg.SessionCookieName = "__Host-admin-session"
	}

	return &Authkit{cfg: cfg, store: store}
}
