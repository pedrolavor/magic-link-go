package services

import (
	"magic-link/application/utils"
	"magic-link/domain/models"
	"net/smtp"
)

// SendToken sends token to the user e-mail
func SendToken(user models.User) error {

	token, err := utils.GenerateToken(32)

	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", "your_email", "your_password", "smtp.gmail.com")
	to := []string{user.Email}
	msg := []byte("Hey there " + user.Name + "!\r\nHere is your login link: http://localhost:8080/token?token=" + token)

	return smtp.SendMail("smtp.gmail.com:587", auth, "magic-link-go@mail.go", to, msg)

}
