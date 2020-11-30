package controllers

import (
	"log"
	"magic-link/domain/repository"
	"net/http"
)

// TokenHandler handles token validations requests
func TokenHandler(res http.ResponseWriter, req *http.Request, tokenRepository repository.TokenRepository) {
	queryParams := req.URL.Query()
	tokenReceived := queryParams.Get("token")

	token, err := tokenRepository.FindValidByToken(tokenReceived)

	if err != nil {
		log.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	err = tokenRepository.UpdateValidation(token.ID, false)

	if err != nil {
		log.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}

	res.Write([]byte("Congratulations!! You're logged in!"))
}
