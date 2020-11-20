package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"magic-link/domain/models"
	infrarepo "magic-link/infra/repository"
	"net/http"
	"net/smtp"
)

var token string = ""

func main() {

	infrarepo.Init()
	infrarepo.Add(models.User{
		ID:    1,
		Name:  "Usuario 1",
		Email: "meuemail@mail.com",
	})

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/token", tokenHandler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}

}

func loginHandler(res http.ResponseWriter, req *http.Request) {
	method := req.Method
	queryParams := req.URL.Query()
	email := queryParams.Get("email")

	user, err := infrarepo.FindByEmail(email)

	if err != nil {
		log.Fatal(err)
	}

	token, err = generateToken(16)

	if err != nil {
		log.Fatal(err)
	}

	sendToken(user, token)

	res.Write([]byte(method))
	j, _ := json.Marshal(user)
	res.Write([]byte(j))
}

func tokenHandler(res http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	tokenReceived := queryParams.Get("token")

	if tokenReceived == token {
		res.Write([]byte("Congratulations!! You're logged in!"))
	} else {
		res.Write([]byte("Sorry :( Your link is not valid!"))
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
