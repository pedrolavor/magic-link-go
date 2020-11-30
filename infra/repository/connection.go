package repositoryimpl

import (
	"database/sql"
	"log"
)

func getConnection() *sql.DB {

	connection := "host=your_host user=your_user password=your_passowrd dbname=your_db sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	log.Println("Connection to database successfully...")

	return db
}
