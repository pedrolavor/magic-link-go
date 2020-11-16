package repository

import "models"

type userRepository interface {
	FindByEmail() models.User
}
