package authkit

import (
	"net/http"
	"time"

	"github.com/dahlhoffKevin/WebApiAuthKit/errorhandler"
)

func (a *Authkit) RequireSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// cookie lesen
		c, err := r.Cookie(a.cfg.SessionCookieName)
		if err != nil || c.Value == "" {
			errorhandler.Write(w, errorhandler.Unauthorized())
			return
		}
		sessionID := c.Value

		// session laden
		sess, err := a.store.Sessions().GetByID(r.Context(), sessionID)
		if err != nil {
			// internes problem -> nicht zu viel leaken
			// TODO: logging
			errorhandler.Write(w, errorhandler.Internal())
			return
		}
		if sess == nil {
			errorhandler.Write(w, errorhandler.Unauthorized())
			return
		}

		// revoked / expire check
		if sess.RevokedAt != nil {
			errorhandler.Write(w, errorhandler.Unauthorized())
			return
		}
		if time.Now().After(sess.ExpiresAt) {
			errorhandler.Write(w, errorhandler.Unauthorized())
			return
		}

		// in context legen
		ctx := withSession(r.Context(), sess)
		r = r.WithContext(ctx)

		// weiter
		next.ServeHTTP(w, r)
	})
}
