package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var port int = 8080

type User struct {
	Name  string
	Email string
}

func main() {

	userRepository := {User{"Pedro", "pedro@email.com"}}

	http.HandleFunc("/send", sendHandler)
	http.HandleFunc("/validate", validateHandler)
	fmt.Println("Server running on port :" + strconv.Itoa(port))
	http.ListenAndServe(":"+strconv.Itoa(port), nil)

}

func sendHandler(res http.ResponseWriter, req *http.Request) {

	queryParams := req.URL.Query()
	email := queryParams.Get("email")

	io.WriteString(res, email)

}

func validateHandler(res http.ResponseWriter, req *http.Request) {

	io.WriteString(res, "Hello Go!")

}

func findByEmail(email string) user {
	return
}
