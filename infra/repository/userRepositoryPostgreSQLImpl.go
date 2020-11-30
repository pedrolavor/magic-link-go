package repositoryimpl

import (
	"errors"
	"log"
	"magic-link/domain/models"

	// postgres driver connection
	_ "github.com/lib/pq"
)

// UserRepositoryPostgreSQLImpl repository struct
type UserRepositoryPostgreSQLImpl struct{}

// FindByEmail returns user by email
func (r *UserRepositoryPostgreSQLImpl) FindByEmail(email string) (models.User, error) {
	db := getConnection()

	query, err := db.Query("SELECT * FROM public.user WHERE public.user.email = $1", email)

	if err != nil {
		log.Fatalln(err.Error())

		return models.User{}, err
	}

	var user models.User

	if query.Next() {
		var id int
		var nome, email string

		err := query.Scan(&id, &nome, &email)

		if err != nil {
			log.Fatalln(err.Error())

			return models.User{}, err
		}

		user = models.User{
			ID:    id,
			Name:  nome,
			Email: email,
		}

	} else {
		defer query.Close()
		defer db.Close()
		return models.User{}, errors.New("E-mail " + email + " not found.")
	}

	defer query.Close()
	defer db.Close()

	return user, nil
}
