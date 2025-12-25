package authkit

import (
	"context"

	"github.com/dahlhoffKevin/WebApiAuthKit/store"
)

// context-key als eigener type (verhindert collision)
type ctxKey int

const (
	ctxSessionKey ctxKey = iota + 1
	ctxUserKey
)

func withSession(ctx context.Context, s *store.Session) context.Context {
	return context.WithValue(ctx, ctxSessionKey, s)
}

func SessionFromContext(ctx context.Context) (*store.Session, bool) {
	s, ok := ctx.Value(ctxSessionKey).(*store.Session)
	return s, ok && s != nil
}

func withUser(ctx context.Context, u *store.User) context.Context {
	return context.WithValue(ctx, ctxUserKey, u)
}

func UserFromContext(ctx context.Context) (*store.User, bool) {
	u, ok := ctx.Value(ctxUserKey).(*store.User)
	return u, ok && u != nil
}
