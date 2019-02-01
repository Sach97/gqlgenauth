package user

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Sach97/gqlgenauth/auth/context"
	"github.com/Sach97/gqlgenauth/auth/deeplinker"
	"github.com/Sach97/gqlgenauth/auth/mailer"
	"github.com/Sach97/gqlgenauth/auth/model"
	"github.com/Sach97/gqlgenauth/auth/tokenizer"
	"github.com/jmoiron/sqlx"
	"github.com/op/go-logging"
	"github.com/rs/xid"
)

// UserService holds the user service struct
type UserService struct {
	db         *sqlx.DB
	log        *logging.Logger
	tokenizer  *tokenizer.Tokenizer
	mailer     *mailer.Service
	deeplinker *deeplinker.FireBaseClient
}

type EmailMessage struct {
	ConfirmationUrl string
}

// NewUserService instantiates user service
func NewUserService(db *sqlx.DB, log *logging.Logger, tokenizer *tokenizer.Tokenizer, mailer *mailer.Service,
	deeplinker *deeplinker.FireBaseClient) *UserService {
	return &UserService{db: db, log: log, tokenizer: tokenizer, mailer: mailer, deeplinker: deeplinker}
}

// CreateUser creates a new user
func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	userID := xid.New()
	user.ID = userID.String()
	userSQL := `INSERT INTO users (id, email, password,confirmed) VALUES (:id, :email, :password, :confirmed)`
	user.HashedPassword()
	user.Confirmed = false
	_, err := u.db.NamedExec(userSQL, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// SendEmail sends an email with a confirmation link to a new user
func (u *UserService) SendConfirmationEmail(user *model.User) error {
	token, err := u.tokenizer.GenerateToken(user.ID)
	if err != nil {
		return err
	}
	link, err := u.deeplinker.GetDynamicLink(token, true)
	if err != nil {
		return err
	}
	message := EmailMessage{
		ConfirmationUrl: link,
	}

	to := []string{user.Email}
	recipients := ""
	subject := "Confirmation email"
	sender := "sacha.arbonel@hotmail.fr"
	inputs := mailer.Inputs{
		Recipients: recipients,
		Subject:    subject,
		Sender:     sender,
		To:         to,
	}

	return u.mailer.SendEmailTemplate(inputs, "confirmation", message)
}

//TODO: handler when email already exists in database (pq: duplicate key value violates unique constraint "users_email_key")

//Get userid from token
// Verify if user exists from userid
//send boolean isConfirmed

func (u *UserService) VerifyUserToken(token string) (bool, error) {
	userID, err := u.tokenizer.GetUserID(token)
	if err != nil {
		u.log.Errorf("Error in retrieving userid : %v", err)
		return false, err
	}
	userExists := u.UserExists(userID)
	if !userExists {
		return false, fmt.Errorf("This user doesnt exists")
	}
	return u.ConfirmUser(userID)

}

// ConfirmUser is a service that sets a confirmed user
func (u *UserService) ConfirmUser(userID string) (bool, error) {
	user := &model.User{}
	updateUserSQL := `UPDATE users SET confirmed = TRUE WHERE id = $1 RETURNING confirmed;`

	udb := u.db.Unsafe()
	row := udb.QueryRowx(updateUserSQL, userID)
	err := row.StructScan(user)
	if err == sql.ErrNoRows {
		return user.Confirmed, nil
	}
	if err != nil {
		u.log.Errorf("Error in updating user : %v", err)
		return user.Confirmed, err
	}
	return user.Confirmed, nil

}

// FindByEmail find a user by email
func (u *UserService) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	userSQL := `SELECT * FROM users WHERE email = $1`
	udb := u.db.Unsafe()
	row := udb.QueryRowx(userSQL, email)
	err := row.StructScan(user)
	if err == sql.ErrNoRows {
		return user, nil
	}
	if err != nil {
		u.log.Errorf("Error in retrieving user : %v", err)
		return nil, err
	}

	return user, nil
}

// UserExists returns true if user exists
func (u *UserService) UserExists(userID string) bool {
	user := &model.User{}

	userSQL := `SELECT * FROM users WHERE ID = $1`
	udb := u.db.Unsafe()
	row := udb.QueryRowx(userSQL, userID)
	err := row.StructScan(user)
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		u.log.Errorf("Error in retrieving user : %v", err)
		return false
	}

	return true
}

//ComparePassword compares two passwords
func (u *UserService) ComparePassword(userCredentials *model.UserCredentials) (*model.User, error) {
	user, err := u.FindByEmail(userCredentials.Email)
	if err != nil {
		return nil, errors.New(context.UnauthorizedAccess)
	}
	if result := user.ComparePassword(userCredentials.Password); !result {
		return nil, errors.New(context.UnauthorizedAccess)
	}
	return user, nil
}
