package models

// Token is a representation
type Token struct {
	ID        int
	Token     string
	Email     string
	ExpiresAt string
	Valid     bool
}
