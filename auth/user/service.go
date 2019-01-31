package user

import (
	"database/sql"
	"errors"

	"github.com/Sach97/gqlgenauth/auth/model"
	"github.com/Sach97/gqlgenauth/auth/tokenizer"
	"github.com/Sach97/gqlgenauth/auth/deeplinker"
	"github.com/Sach97/gqlgenauth/auth/mailer"
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
)


// UserService holds the user service struct
type UserService struct {
	db        *sqlx.DB
	log       *logging.Logger
	tokenizer *tokenizer.Tokenizer
	mailer *mailer.Mailer
	deeplinker *deeplinker.FireBaseClient
}

// NewUserService instantiates user service
func NewUserService(db *sqlx.DB, log *logging.Logger,tokenizer *tokenizer.Tokenizer,mailer *mailer.Mailer
	deeplinker *deeplinker.Deeplinker) *UserService {
	return &UserService{db: db, log: log, tokenizer: tokenizer, mailer:mailer,deeplinker:deeplinker}
}

// CreateUser creates a new user
func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	userID := xid.New()
	user.ID = userId.String()
	userSQL := `INSERT INTO users (id, email, password) VALUES (:id, :email, :password)`
	user.HashedPassword()
	user.Confirmed := false
	_, err := u.db.NamedExec(userSQL, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// SendEmail sends an email with a confirmation link to a new user
func (u *UserService) SendConfirmationEmail(user *model.User, p *tokenizer.Payload) error {
	token,err := u.tokenizer.GenerateToken(user.ID)
	if err != nil {
		return err
	}
	
	link, _ := u.deeplinker.GetDynamicLink(token, true)
	u.mailer.SendConfirmationEmail("Activate your account by clicking on this link",user.Email)
}

// ConfirmUser is a service that sets a confirmed user
func (u *UserService) ConfirmUser(userID string) (bool, error) {
	user := &model.User{}
	updateUserSQL := `UPDATE users SET confirmed = TRUE WHERE id = $1;`

	udb := u.db.Unsafe()
	row := udb.QueryRowx(updateUserSQL, userID)
	err := row.StructScan(user)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		u.log.Errorf("Error in retrieving user : %v", err)
		return false, err
	}
	return user, nil

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
