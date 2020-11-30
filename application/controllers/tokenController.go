package controllers

import (
	"magic-link/domain/repository"
	"net/http"
)

// TokenHandler handles token validations requests
func TokenHandler(res http.ResponseWriter, req *http.Request, userRepository repository.UserRepository) {
	queryParams := req.URL.Query()
	tokenReceived := queryParams.Get("token")

	// TODO: implement persistence
	if tokenReceived == tokenReceived {
		res.Write([]byte("Congratulations!! You're logged in!"))
	} else {
		res.Write([]byte("Sorry :( Your link is not valid!"))
	}
}
