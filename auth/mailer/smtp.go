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

//TODO: retrieve values from config

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

//TODO: SendEmail(input) for mirroring aws SES API https://github.com/awsdocs/aws-doc-sdk-examples/blob/7a81218bd33bd74e7364561b9df66814df876cdb/go/example_code/ses/ses_send_email.go#L107

//SendEmail sends an email
func (c *Service) SendEmail(from string, to []string, msg []byte) error { //SendEmail(to []string, template string)
	auth := smtp.PlainAuth(c.Identity, c.Username, c.Password, c.Host)
	err := smtp.SendMail(c.Address, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// func (c *Service) SendConfirmationEmail(from string, to []string, msg []byte) error {
// 	SendEmail()
// 	return err
// }
