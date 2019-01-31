package mailer

import (
	"log"
	"net/smtp"

	"github.com/Sach97/gqlgenauth/auth/context"
)

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

//Service holds our service struct
type Service struct {
	Identity string
	Username string
	Password string
	Host     string
	Address  string
}

//Inputs holds our inputs struct
type Inputs struct {
	Recipients string
	Subject    string
	Body       string
	Sender     string
	To         []string
}

//Message holds our message struct
type Message struct {
	Msg    []byte
	Sender string
	To     []string
}

//NewMessage carefully craft a new message from inputs struct
func (s *Service) NewMessage(inputs Inputs) Message {
	body := "To: " + inputs.To[0] + "\r\nSubject: " + inputs.Subject + "\r\n" + MIME + "\r\n" + inputs.Body
	msg := []byte(body)
	return Message{
		Msg:    msg,
		To:     inputs.To,
		Sender: inputs.Sender,
	}
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

//SendEmail sends an email
func (s *Service) SendEmail(message Message) error {
	auth := smtp.PlainAuth(s.Identity, s.Username, s.Password, s.Host)
	err := smtp.SendMail(s.Address, auth, message.Sender, message.To, message.Msg)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// func (c *Service) SendConfirmationEmail(from string, to []string, msg []byte) error {
// 	SendEmail
// 	return err
// }
