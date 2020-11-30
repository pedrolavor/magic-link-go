package repository

import "magic-link/domain/models"

// TokenRepository interface
type TokenRepository interface {
	Save(token models.Token) (models.Token, error)
	FindByToken(token string) (models.Token, error)
	FindValidByToken(token string) (models.Token, error)
	UpdateValidation(id int, valid bool) error
}
