package entity

type Token struct {
	Jwt      string
	MaxAge   int
	Domain   string
	Secure   bool
	HttpOnly bool
}
