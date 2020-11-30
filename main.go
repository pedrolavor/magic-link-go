package main

import (
	"log"
	"magic-link/application/controllers"
	"magic-link/domain/repository"
	repositoryimpl "magic-link/infra/repository"
	"net/http"
)

const port = ":8080"

func main() {

	userRepository := &repositoryimpl.UserRepositoryPostgreSQLImpl{}
	tokenRepository := &repositoryimpl.TokenRepositoryPostgreSQLImpl{}
	loadRoutes(userRepository, tokenRepository)

	log.Printf("Server running on port %s...", port)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		panic(err.Error())
	}

}

func loadRoutes(userRepository repository.UserRepository, tokenRepository repository.TokenRepository) {

	http.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
		controllers.LoginHandler(res, req, userRepository, tokenRepository)
	})

	http.HandleFunc("/token", func(res http.ResponseWriter, req *http.Request) {
		controllers.TokenHandler(res, req, tokenRepository)
	})
}
