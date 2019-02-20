package user

import (
	"fmt"

	"github.com/Sach97/gqlgenauth/auth/model"
)

//TODO: strategy pattern or something if the user doesnt want to use deeplinks but just an email

//Signup register a user in db and send an email and return instructions
func (u *Service) Signup(credentials *model.UserCredentials) string {
	user := &model.User{
		Email:    credentials.Email,
		Password: credentials.Password,
	}
	userExists := u.userEmailExists(user.Email)
	if userExists && !user.Confirmed {
		return fmt.Sprintf("Sorry this email already exists. You must login or confirm your email address")
	}
	newUser, err := u.createUser(user)
	if err != nil {
		u.log.Errorf("Error during user creation : %v", err)
		return fmt.Sprintf("Sorry an error occured please try again")
	}

	err = u.sendConfirmationEmail(newUser)
	if err != nil {
		u.log.Errorf("Error during sending email : %v", err)
		return fmt.Sprintf("Sorry an error occured please try again")
	}
	return fmt.Sprintf("We've just sent you a confirmation email to %s", user.Email)

}

//Login return a jwt token if user is confirmed
func (u *Service) Login(credentials *model.UserCredentials) (string, error) {
	//TODO: builder input
	userExists := u.userEmailExists(credentials.Email)

	if !userExists {
		return "", u.msg.CredentialsError()
	}
	user, err := u.comparePassword(credentials)
	if err != nil {
		return "", err
	}

	if !user.Confirmed {
		return "", fmt.Errorf("We've sent you an email to %s please click on the click in order to complete your registration", user.Email)
	}

	token, err := u.signJWT(user) //TODO: same builder input here
	if err != nil {
		u.log.Errorf("Error during jwt signing of user: %v", err)
	}
	return token, nil
}

//VerifyUserToken decode userid from token and verify if exists
func (u *Service) VerifyUserToken(token string) (bool, error) {
	userID, err := u.tokenizer.GetUserID(token)
	if err != nil {
		u.log.Errorf("Error in retrieving userid for token %s: %v", token, err)
		return false, fmt.Errorf("Your session has expired. ")
	}

	//TODO: handle session expired

	userExists := u.userIDExists(userID)
	if !userExists {
		return false, fmt.Errorf("This user doesnt exists")
	}
	return u.confirmUser(userID)

}

func (u *Service) FindByID(userID string) (*model.User, error) {

	userExists := u.userIDExists(userID)
	if !userExists {
		return nil, fmt.Errorf("This user doesnt exists")
	}
	return u.findByID(userID)
}
