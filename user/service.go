package user

import (
	"fmt"

	"github.com/Sach97/ninshoo/builder"
	gcontext "github.com/Sach97/ninshoo/context"
	"github.com/Sach97/ninshoo/deeplinker"
	"github.com/Sach97/ninshoo/jwt"
	"github.com/Sach97/ninshoo/mailer"
	"github.com/Sach97/ninshoo/model"
	"github.com/Sach97/ninshoo/tokenizer"
	"github.com/jmoiron/sqlx"
	"github.com/op/go-logging"
)

//Service holds the user service struct
type Service struct {
	msg        *gcontext.MessageService
	db         *sqlx.DB
	log        *logging.Logger
	tokenizer  *tokenizer.Tokenizer
	mailer     *mailer.Service
	deeplinker *deeplinker.FireBaseClient
	jwt        *jwt.AuthService
	builder    *builder.BuilderService
}

//EmailMessage holds our email struct
type EmailMessage struct {
	ConfirmationURL string
	//Username
}

// NewUserService instantiates user service
func NewUserService(msg *gcontext.MessageService, db *sqlx.DB, log *logging.Logger, jwt *jwt.AuthService, tokenizer *tokenizer.Tokenizer, mailer *mailer.Service,
	deeplinker *deeplinker.FireBaseClient, builder *builder.BuilderService) *Service {
	return &Service{msg: msg, db: db, log: log, jwt: jwt, tokenizer: tokenizer, mailer: mailer, deeplinker: deeplinker, builder: builder}
}

//signJWT sign a user jwt
func (u *Service) signJWT(user *model.User) (string, error) { //TODO: cleaner way to do this
	claims := u.builder.BuildCustomClaims(user)
	fmt.Println(claims)
	//TODO: fetch roles from db
	token, err := u.jwt.SignJWT(claims)
	return token, err
}

// sendConfirmationEmail sends an email with a confirmation link to a new user
func (u *Service) sendConfirmationEmail(user *model.User) error {
	token, err := u.tokenizer.GenerateToken(user.ID)

	if err != nil {
		u.log.Errorf("Error during generation of confirmation token related to the userID %s: %v", user.ID, err)
		return err
	}
	link, err := u.deeplinker.GetDynamicLink(token, true)
	if err != nil {
		u.log.Errorf("Error retrieving deeplink from firebase : %v", err)
		return err
	}
	//TODO: deeplinker strategy None or Firebase. None return example.com/confirmation?token=<token>
	message := EmailMessage{
		ConfirmationURL: link,
	}
	//TODO: put username or something else in message

	to := []string{user.Email}
	recipients := ""
	subject := "Confirmation email"
	sender := user.Email
	inputs := mailer.Inputs{
		Recipients: recipients,
		Subject:    subject,
		Sender:     sender,
		To:         to,
	}

	//TODO: make a builder for email
	err = u.mailer.SendEmailTemplate(inputs, "confirmation", message)
	if err != nil {
		u.log.Errorf("An error occured when sending email : %v", err)
	}
	return err
}

//ComparePassword compares two passwords
func (u *Service) comparePassword(userCredentials *model.UserCredentials) (*model.User, error) {
	user, err := u.findByEmail(userCredentials.Email)
	if err != nil {
		return nil, u.msg.CredentialsError()
	}
	result, err := user.ComparePassword(userCredentials.Password)
	if err != nil {
		u.log.Errorf("Error comparing passwords : %v", err)
	}
	if !result {
		return nil, u.msg.CredentialsError() //TODO: inspiration from this for error message config
	}
	return user, nil
}
