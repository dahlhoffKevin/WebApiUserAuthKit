package authkit

type CookieConfig struct {
	Name     string
	Secure   bool
	HttpOnly bool
	SameSite string
	Path     string
}
