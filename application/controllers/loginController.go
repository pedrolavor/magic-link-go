package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"magic-link/application/services"
	"magic-link/domain/models"
	"magic-link/domain/repository"
	"net/http"
	"net/smtp"
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

func sendToken(user models.User, token string) {

	auth := smtp.PlainAuth("", "from@email.com", "password", "smtp.gmail.com")

	to := []string{user.Email}
	msg := []byte("Hey there " + user.Name + "!" +
		"\r\nHere is your login link: http://localhost:8080/token?token=" + token)

	err := smtp.SendMail("smtp.gmail.com:587", auth, "magic-link-go@mail.go", to, msg)
	if err != nil {
		log.Fatal(err)
	}

}

func generateToken(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), err
}
