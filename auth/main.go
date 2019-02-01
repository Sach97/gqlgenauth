package main

import (
	"fmt"

	"github.com/Sach97/gqlgenauth/auth/context"
	"github.com/Sach97/gqlgenauth/auth/db"
	"github.com/Sach97/gqlgenauth/auth/deeplinker"
	"github.com/Sach97/gqlgenauth/auth/mailer"
	"github.com/Sach97/gqlgenauth/auth/tokenizer"
	"github.com/Sach97/gqlgenauth/auth/user"
	"github.com/Sach97/gqlgenauth/auth/utils"
)

type EmailMessage struct {
	ConfirmationUrl string
}

func main() {

	// // Context Stuffs
	cfg := context.LoadConfig(".")

	// // Token stuffs
	RedisClient := tokenizer.NewRedisClient()
	t := tokenizer.Tokenizer{RedisClient}

	// // Mail stuffs
	m := mailer.NewMailer(cfg)

	// // Firebase STUFFS
	d := deeplinker.NewFireBaseClient(cfg)

	//DB STUFFS
	sql := db.Strategy(db.DriverSQL{Name: "postgres"})

	s, err := sql.OpenDB(cfg)
	if err != nil {
		panic(err)
	}

	//Signup(email,password) mutation
	//create user done
	//send email confirmation done

	//ConfirmUser(token) query
	//Get userid from token
	// Verify if user exists from userid
	//sets user as confirmed in database from userid
	//send boolean isConfirmed

	//Login(user,password) -> AuthToken mutation
	//flutter side => save authtoken is shared preference or secure storage
	//add to header for next requests

	//Log stuffs
	l := utils.NewLoggerService(cfg)

	// User service stuffs
	u := user.NewUserService(s, l, &t, m, d)

	// fmt.Println("we are here")
	// user := &model.User{
	// 	Email:    "sacha.arbonel@hotmail.fr",
	// 	Password: "secretpassword",
	// }
	// _, err = u.CreateUser(user)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(user.ID)
	// u.SendConfirmationEmail(user)

	token := "22ab6e4d-143d-4941-a0cc-805df3748270"
	verified, err := u.VerifyToken(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(verified)

}
