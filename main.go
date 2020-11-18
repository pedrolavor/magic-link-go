package main

import (
	"encoding/json"
	"magic-link/domain/models"
	infrarepo "magic-link/infra/repository"
	"net/http"
)

func main() {

	infrarepo.Init()
	infrarepo.Add(models.User{
		ID:    1,
		Name:  "Pedro Lavor",
		Email: "pedro@email.com",
	})

	http.HandleFunc("/login", loginHandler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}

}

func loginHandler(res http.ResponseWriter, req *http.Request) {
	method := req.Method
	queryParams := req.URL.Query()
	email := queryParams.Get("email")

	user := infrarepo.FindByEmail(email)
	res.Write([]byte(method))
	j, _ := json.Marshal(user)
	res.Write([]byte(j))
}
