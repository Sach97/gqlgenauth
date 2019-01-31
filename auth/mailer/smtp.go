package mailer

import (
	"log"
	"net/smtp"

	"github.com/Sach97/gqlgenauth/auth/context"
)

//Service holds our service struct
type Service struct {
	Identity string
	Username string
	Password string
	Host     string
	Address  string
}

//NewMailer instantiates the mailer service from config file
func NewMailer(config *context.Config) *Service {
	if config.SMTPPassword == "" {
		panic("You must set your smtp password")
	}
	return &Service{
		Identity: config.SMTPIdentity,
		Username: config.SMTPUsername,
		Password: config.SMTPPassword,
		Host:     config.SMTPHost,
		Address:  config.SMTPAddress,
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

// func (c *Service) SendEmailTemplate(from string, to []string, data []byte, emailType string) error {

// }

// func (c *Service) SendConfirmationEmail(from string, to []string, msg []byte) error {
// 	SendEmail
// 	return err
// }
