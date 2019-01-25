package main

import (
	"fmt"
	"os"

	"github.com/Sach97/gqlgenauth/auth/mailer"
)

func main() {

	// client := utils.New()

	// token, err := client.GenerateString()
	// if err != nil {
	// 	panic(err)
	// }

	// value, err := client.GetToken(token)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(value)

	sendgrid := os.Getenv("SENDGRID")
	fmt.Println(sendgrid)
	cfg := mailer.Service{
		"",
		"apikey",
		sendgrid,
		"smtp.sendgrid.net",
		"smtp.sendgrid.net:25",
	}
	client := mailer.NewMailer(cfg)
	to := []string{"sacha.arbonel@hotmail.fr"}

	msg := []byte("To: recipient@example.net\r\n" +

		"Subject: testing with get env!\r\n" +

		"\r\n" +

		"This is the email body.\r\n")
	client.SendEmail("sacha.arbonel@hotmail.fr", to, msg)

}
