package controllers

import (
	"encoding/json"
	"log"
	"magic-link/application/services"
	"magic-link/domain/repository"
	"net/http"
)

// LoginHandler handles login requests
func LoginHandler(res http.ResponseWriter, req *http.Request, userRepository repository.UserRepository) {
	method := req.Method

	if method == "POST" {

		email := req.URL.Query().Get("email")

		user, err := userRepository.FindByEmail(email)

		if err != nil {
			log.Println(err)
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(err.Error()))
			return
		}

		err = services.SendToken(user)

		if err != nil {
			log.Println(err)
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(err.Error()))
			return
		}

		j, _ := json.Marshal(user)
		res.Write([]byte(j))

	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}
