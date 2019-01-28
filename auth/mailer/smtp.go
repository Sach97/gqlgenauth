package mailer

import (
	"log"
	"net/smtp"
)

type Service struct {
	Identity string
	Username string
	Password string
	Host     string
	Address  string
}

func NewMailer(c Service) *Service {
	return &Service{
		Identity: c.Identity,
		Username: c.Username,
		Password: c.Password,
		Host:     c.Host,
		Address:  c.Address,
	}
}

//TODO: confirmation email template

//SendEmail sends an email
func (c *Service) SendEmail(from string, to []string, msg []byte) error {
	auth := smtp.PlainAuth(c.Identity, c.Username, c.Password, c.Host)

	err := smtp.SendMail(c.Address, auth, from, to, msg)

	if err != nil {

		log.Fatal(err)

	}
	return err
}
