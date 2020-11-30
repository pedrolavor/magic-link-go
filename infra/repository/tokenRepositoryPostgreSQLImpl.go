package repositoryimpl

import (
	"errors"
	"log"
	"magic-link/domain/models"

	// postgres driver connection
	_ "github.com/lib/pq"
)

// TokenRepositoryPostgreSQLImpl repository struct
type TokenRepositoryPostgreSQLImpl struct{}

// Save persistss a new token
func (r *TokenRepositoryPostgreSQLImpl) Save(token models.Token) (models.Token, error) {
	db := getConnection()

	var insertedID int
	query := db.QueryRow("INSERT INTO public.token (token, email, valid) VALUES ($1, $2, $3) RETURNING id", token.Token, token.Email, token.Valid)

	err := query.Scan(&insertedID)

	if err != nil {
		log.Println(err.Error())

		return models.Token{}, err
	}

	token.ID = int(insertedID)

	defer db.Close()

	return token, nil
}

// FindByToken returns user by token
func (r *TokenRepositoryPostgreSQLImpl) FindByToken(receivedToken string) (models.Token, error) {
	db := getConnection()

	query, err := db.Query("SELECT id, token, email, valid FROM public.token WHERE public.token.token = $1", receivedToken)

	if err != nil {
		log.Println(err.Error())

		return models.Token{}, err
	}

	var token models.Token

	if query.Next() {
		var id int
		var tokenStr, email string
		var valid bool

		err := query.Scan(&id, &tokenStr, &email, &valid)

		if err != nil {
			log.Println(err.Error())

			return models.Token{}, err
		}

		token = models.Token{
			ID:    id,
			Token: tokenStr,
			Email: email,
			Valid: valid,
		}

	} else {
		defer query.Close()
		defer db.Close()
		return models.Token{}, errors.New("Token " + receivedToken + " not found.")
	}

	defer query.Close()
	defer db.Close()

	return token, nil
}

// FindValidByToken returns user by token
func (r *TokenRepositoryPostgreSQLImpl) FindValidByToken(receivedToken string) (models.Token, error) {
	db := getConnection()

	query, err := db.Query("SELECT id, token, email, valid FROM public.token WHERE public.token.token = $1 AND valid = true", receivedToken)

	if err != nil {
		log.Println(err.Error())

		return models.Token{}, err
	}

	var token models.Token

	if query.Next() {
		var id int
		var tokenStr, email string
		var valid bool

		err := query.Scan(&id, &tokenStr, &email, &valid)

		if err != nil {
			log.Println(err.Error())

			return models.Token{}, err
		}

		token = models.Token{
			ID:    id,
			Token: tokenStr,
			Email: email,
			Valid: valid,
		}

	} else {
		defer query.Close()
		defer db.Close()
		return models.Token{}, errors.New("Token " + receivedToken + " not found.")
	}

	defer query.Close()
	defer db.Close()

	return token, nil
}

// UpdateValidation updates token valid property
func (r *TokenRepositoryPostgreSQLImpl) UpdateValidation(id int, valid bool) error {
	db := getConnection()

	stmt, err := db.Prepare("UPDATE public.token SET valid = $2 WHERE id = $1")

	if err != nil {
		log.Println(err.Error())

		return err
	}

	_, err = stmt.Exec(id, valid)

	if err != nil {
		log.Println(err.Error())

		return err
	}

	defer stmt.Close()
	defer db.Close()

	return nil
}
