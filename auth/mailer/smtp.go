package smtp

import (
	"log"
	"net/smtp"
)

type Config struct {
	string identity,
	string username,
	string password,
	string host,
	string address,
}

func (c *Config) SendMail(from string,to []string , msg []byte) {
	auth := smtp.PlainAuth(c.identity, c.username, c.password, c.host) //smtp.PlainAuth("", "user@example.com", "password", "mail.example.com")

	// Connect to the server, authenticate, set the sender and recipient,

	// and send the email all in one step.

	// to := []string{"recipient@example.net"}

	// msg := []byte("To: recipient@example.net\r\n" +

	// 	"Subject: discount Gophers!\r\n" +

	// 	"\r\n" +

	// 	"This is the email body.\r\n")

	err := smtp.SendMail(c.address, auth, from, to, msg)

	if err != nil {

		log.Fatal(err)

	}
}
