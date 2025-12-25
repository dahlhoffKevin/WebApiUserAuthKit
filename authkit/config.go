package authkit

import "time"

type AutoSchemaConfig struct {
	Enabled bool
	LockKey int
}

type Config struct {
	SessionCookieName string
	SessionTTL        time.Duration
}

type AlternateConfig struct {
	Store        string
	Cookie       CookieConfig
	AdminOrigin  string
	PublicOrigin string
	APIOrigin    string
	CSRF         CSRFConfig
	AutoSchema   AutoSchemaConfig
	RateLimit    RateLimitConfig
}
