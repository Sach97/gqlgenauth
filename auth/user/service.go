package user

import (
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/Sach97/gqlgenauth/auth/context"
	"github.com/Sach97/gqlgenauth/auth/deeplinker"
	"github.com/Sach97/gqlgenauth/auth/jwt"
	"github.com/Sach97/gqlgenauth/auth/mailer"
	"github.com/Sach97/gqlgenauth/auth/model"
	"github.com/Sach97/gqlgenauth/auth/tokenizer"
	"github.com/jmoiron/sqlx"
	"github.com/op/go-logging"
	"github.com/rs/xid"
)

//Service holds the user service struct
type Service struct {
	msg        *context.MessageService
	db         *sqlx.DB
	log        *logging.Logger
	tokenizer  *tokenizer.Tokenizer
	mailer     *mailer.Service
	deeplinker *deeplinker.FireBaseClient
	jwt        *jwt.AuthService
}

//EmailMessage holds our email struct
type EmailMessage struct {
	ConfirmationURL string
}

type CustomNamespace struct {
	Sub                    string                 `json:"sub"`
	Name                   string                 `json:"name"`
	Admin                  bool                   `json:"admin"`
	Iat                    int64                  `json:"iat"`
	HTTPSHasuraIoJwtClaims HTTPSHasuraIoJwtClaims `json:"https://hasura.io/jwt/claims"`
}

type HTTPSHasuraIoJwtClaims struct {
	XHasuraAllowedRoles []string `json:"x-hasura-allowed-roles"`
	XHasuraDefaultRole  string   `json:"x-hasura-default-role"`
	XHasuraUserID       string   `json:"x-hasura-user-id"`
	XHasuraOrgID        string   `json:"x-hasura-org-id"`
	XHasuraCustom       string   `json:"x-hasura-custom"`
}

// NewUserService instantiates user service
func NewUserService(msg *context.MessageService, db *sqlx.DB, log *logging.Logger, jwt *jwt.AuthService, tokenizer *tokenizer.Tokenizer, mailer *mailer.Service,
	deeplinker *deeplinker.FireBaseClient) *Service {
	return &Service{msg: msg, db: db, log: log, jwt: jwt, tokenizer: tokenizer, mailer: mailer, deeplinker: deeplinker}
}

// CreateUser creates a new user
func (u *Service) CreateUser(user *model.User) (*model.User, error) {
	userID := xid.New()
	user.ID = userID.String()
	userSQL := `INSERT INTO users (id, email, password,confirmed) VALUES (:id, :email, :password, :confirmed)`
	err := user.HashedPassword()
	if err != nil {
		u.log.Errorf("Error during hashing password : %v", err)
	}
	user.Confirmed = false
	_, err = u.db.NamedExec(userSQL, user)
	if err != nil {
		u.log.Errorf("Error during sql execution query of user creation : %v", err)
		return nil, err
	}
	return user, nil
}

func (u *Service) Signup(credentials *model.UserCredentials) string {
	user := &model.User{
		Email:    credentials.Email,
		Password: credentials.Password,
	}
	newUser, err := u.CreateUser(user)
	if err != nil {
		u.log.Errorf("Error during user creation : %v", err)
		return fmt.Sprintf("This email already exists")
	}
	//fmt.Println(user.ID)
	err = u.SendConfirmationEmail(newUser)
	if err != nil {
		u.log.Errorf("Error during sending email")
		return fmt.Sprintf("Sorry an error occured please try again")
	}
	return fmt.Sprintf("We've just sent you a confirmation email to %s", user.Email)

}
func (u *Service) Login(credentials *model.UserCredentials) (string, error) {
	user, err := u.ComparePassword(credentials)
	if err != nil {
		return "", err
	}

	token, err := u.SignJwt(user)
	if err != nil {
		u.log.Errorf("Error during jwt signing of user: %v", err)
	}
	return token, nil
}

func (u *Service) SignJwt(user *model.User) (string, error) { //TODO: cleaner way to do this
	customMapClaims := CustomNamespace{
		Sub:   base64.StdEncoding.EncodeToString([]byte(user.ID)),
		Name:  base64.StdEncoding.EncodeToString([]byte(user.Username)),
		Admin: true, //TODO: change this
		//Iat:   time.Now().Add(time.Second * *time.Duration(cfg.JWTExpireIn)).Unix(),
		HTTPSHasuraIoJwtClaims: HTTPSHasuraIoJwtClaims{
			XHasuraAllowedRoles: []string{"user", "editor"},
			XHasuraDefaultRole:  "user",
			XHasuraOrgID:        base64.StdEncoding.EncodeToString([]byte(user.ID)),
			XHasuraCustom:       "custom-value",
		},
	}
	tokenb, err := u.jwt.SignJWT(customMapClaims)
	t := []byte(*tokenb)
	token := string(t)
	return token, err
}

// SendConfirmationEmail sends an email with a confirmation link to a new user
func (u *Service) SendConfirmationEmail(user *model.User) error {
	token, err := u.tokenizer.GenerateToken(user.ID)

	if err != nil {
		u.log.Errorf("Error during generation of confirmation token related to the userID %s: %v", user.ID, err)
		return err
	}
	link, err := u.deeplinker.GetDynamicLink(token, true)
	if err != nil {
		u.log.Errorf("Error during retrieving deeplink from firebase : %v", err)
		return err
	}
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
	err = u.mailer.SendEmailTemplate(inputs, "confirmation", message)
	if err != nil {
		u.log.Errorf("An error occured when sending email : %v", err)
	}
	return err
}

//TODO: handler when email already exists in database (pq: duplicate key value violates unique constraint "users_email_key")

//VerifyUserToken decode userid from token and verify if exists
func (u *Service) VerifyUserToken(token string) (bool, error) {
	userID, err := u.tokenizer.GetUserID(token)
	if err != nil {
		u.log.Errorf("Error in retrieving userid for token %s: %v", token, err)
		return false, err
	}
	userExists := u.UserExists(userID)
	if !userExists {
		u.log.Errorf("Error verifying if user corresponding to userID %s exists : %v", userID, err)
		return false, fmt.Errorf("This user doesnt exists")
	}
	return u.ConfirmUser(userID)

}

// ConfirmUser is a service that sets a confirmed user
func (u *Service) ConfirmUser(userID string) (bool, error) {
	user := &model.User{}
	updateUserSQL := `UPDATE users SET confirmed = TRUE WHERE id = $1 RETURNING confirmed;`

	udb := u.db.Unsafe()
	row := udb.QueryRowx(updateUserSQL, userID)
	err := row.StructScan(user)
	if err == sql.ErrNoRows {
		return user.Confirmed, nil
	}
	if err != nil {
		u.log.Errorf("Error in setting user to confirmed %s : %v", userID, err)
		return user.Confirmed, err
	}
	return user.Confirmed, nil

}

//TODO: strategy pattern for other database than sql

// FindByEmail find a user by email
func (u *Service) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	userSQL := `SELECT * FROM users WHERE email = $1`
	udb := u.db.Unsafe()
	row := udb.QueryRowx(userSQL, email)
	err := row.StructScan(user)
	if err == sql.ErrNoRows {
		return user, nil
	}
	if err != nil {
		u.log.Errorf("Error in retrieving user by his email %s: %v", email, err)
		return nil, err
	}

	return user, nil
}

// UserExists returns true if user exists
func (u *Service) UserExists(userID string) bool {
	user := &model.User{}
	userSQL := `SELECT * FROM users WHERE ID = $1`
	udb := u.db.Unsafe()
	row := udb.QueryRowx(userSQL, userID)
	err := row.StructScan(user)
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		u.log.Errorf("Error verifying if user exists by its id %s: %v", userID, err)
		return false
	}

	return true
}

//ComparePassword compares two passwords
func (u *Service) ComparePassword(userCredentials *model.UserCredentials) (*model.User, error) {
	user, err := u.FindByEmail(userCredentials.Email)
	if err != nil {
		return nil, u.msg.CredentialsError()
	}
	result, err := user.ComparePassword(userCredentials.Password)
	if err != nil {
		u.log.Errorf("Error comparing passwords : %v", err)
	}
	if !result {
		return nil, u.msg.CredentialsError()
	}
	return user, nil
}
