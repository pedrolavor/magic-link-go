package infrarepo

import "magic-link/domain/models"

var userRepository []models.User
var lastIndex int = 0

// Init initialize repo
func Init() {
	userRepository = make([]models.User, 10)
}

// Add add new user
func Add(user models.User) {
	userRepository[lastIndex] = user
	lastIndex++
}

// FindByEmail returns user by email
func FindByEmail(email string) (models.User, error) {
	for i := 0; i < len(userRepository); i++ {
		if userRepository[i].Email == email {
			return userRepository[i], nil
		}
	}
	return models.User{}, nil
}
