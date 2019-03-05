package user

import (
	"database/sql"

	"github.com/Sach97/ninshoo/model"
	"github.com/rs/xid"
)

// createUser creates a new user
func (u *Service) createUser(user *model.User) (*model.User, error) {
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

// confirmUser is a service that sets a confirmed user
func (u *Service) confirmUser(userID string) (bool, error) {
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

// findByEmail find a user by email
func (u *Service) findByEmail(email string) (*model.User, error) {
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

// findByID find a user by email
func (u *Service) findByID(userID string) (*model.User, error) {
	//TODO: public facing errors
	user := &model.User{}

	userSQL := `SELECT * FROM users WHERE ID = $1`
	udb := u.db.Unsafe()
	row := udb.QueryRowx(userSQL, userID)
	err := row.StructScan(user)
	if err == sql.ErrNoRows {
		return user, nil
	}
	if err != nil {
		u.log.Errorf("Error in retrieving user by his id %s: %v", userID, err)
		return nil, err
	}

	return user, nil
}

// UserIDExists returns true if user exists with is id
func (u *Service) userIDExists(userID string) bool {
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

// userEmailExists returns true if user exists with this email
func (u *Service) userEmailExists(userEmail string) bool {
	user := &model.User{}
	userSQL := `SELECT * FROM users WHERE Email = $1`
	udb := u.db.Unsafe()
	row := udb.QueryRowx(userSQL, userEmail)
	err := row.StructScan(user)
	if err == sql.ErrNoRows {
		return false
	}
	if err != nil {
		u.log.Errorf("Error verifying if user exists by its id %s: %v", userEmail, err)
		return false
	}

	return true
}
