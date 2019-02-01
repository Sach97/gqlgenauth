package user

import (
	"database/sql"
	"encoding/base64"
	"errors"
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

//Service holds the user service struct
type Service struct {
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

// NewUserService instantiates user service
func NewUserService(db *sqlx.DB, log *logging.Logger, jwt *jwt.AuthService, tokenizer *tokenizer.Tokenizer, mailer *mailer.Service,
	deeplinker *deeplinker.FireBaseClient) *Service {
	return &Service{db: db, log: log, jwt: jwt, tokenizer: tokenizer, mailer: mailer, deeplinker: deeplinker}
}

// CreateUser creates a new user
func (u *Service) CreateUser(user *model.User) (*model.User, error) {
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

// func (u *Service) Login(user *model.User)  {
// if password match user exists and is confirmed signjwt and return token
// }

func (u *Service) SignJwt(user *model.User) (string, error) { //TODO cleaner way to do this
	customMap := CustomNamespace{
		Sub:   base64.StdEncoding.EncodeToString([]byte(user.ID)),
		Name:  base64.StdEncoding.EncodeToString([]byte(user.Username)),
		Admin: true,
		//Iat:   time.Now().Add(time.Second * *time.Duration(cfg.JWTExpireIn)).Unix(),
		HTTPSHasuraIoJwtClaims: HTTPSHasuraIoJwtClaims{
			XHasuraAllowedRoles: []string{"user", "editor"},
			XHasuraDefaultRole:  "user",
			XHasuraOrgID:        base64.StdEncoding.EncodeToString([]byte(user.ID)),
			XHasuraCustom:       "custom-value",
		},
	}

	//TODO: if user if user is confirmed sign token
	tokenb, err := u.jwt.SignJWT(customMap)
	t := []byte(*tokenb)
	token := string(t)
	return token, err
}

// SendConfirmationEmail sends an email with a confirmation link to a new user
func (u *Service) SendConfirmationEmail(user *model.User) error {
	token, err := u.tokenizer.GenerateToken(user.ID)
	if err != nil {
		return err
	}
	link, err := u.deeplinker.GetDynamicLink(token, true)
	if err != nil {
		return err
	}
	message := EmailMessage{
		ConfirmationURL: link,
	}

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

	return u.mailer.SendEmailTemplate(inputs, "confirmation", message)
}

//TODO: handler when email already exists in database (pq: duplicate key value violates unique constraint "users_email_key")

//Get userid from token
//
//send boolean isConfirmed

//VerifyUserToken decode userid from token and verify if exists
func (u *Service) VerifyUserToken(token string) (bool, error) {
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
		u.log.Errorf("Error in updating user : %v", err)
		return user.Confirmed, err
	}
	return user.Confirmed, nil

}

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
		u.log.Errorf("Error in retrieving user : %v", err)
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
		u.log.Errorf("Error in retrieving user : %v", err)
		return false
	}

	return true
}

//ComparePassword compares two passwords
func (u *Service) ComparePassword(userCredentials *model.UserCredentials) (*model.User, error) {
	user, err := u.FindByEmail(userCredentials.Email)
	if err != nil {
		return nil, errors.New(context.UnauthorizedAccess)
	}
	if result := user.ComparePassword(userCredentials.Password); !result {
		return nil, errors.New(context.UnauthorizedAccess)
	}
	return user, nil
}
