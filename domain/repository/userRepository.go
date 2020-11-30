package repository

import "magic-link/domain/models"

// UserRepository interface
type UserRepository interface {
	FindByEmail(email string) (models.User, error)
}
