package main

import (
	"github.com/Sach97/gqlgenauth/auth/context"
	"github.com/Sach97/gqlgenauth/auth/db"
	"github.com/Sach97/gqlgenauth/auth/deeplinker"
	"github.com/Sach97/gqlgenauth/auth/jwt"
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
	RedisClient := tokenizer.NewRedisClient(cfg)
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

	//Signup(email,password) mutation done
	//create user done
	//send email confirmation done

	//ConfirmUser(token) query
	//Get userid from token done
	// Verify if user exists from userid done
	//sets user as confirmed in database from userid done
	//send boolean isConfirmed done

	//create jwt token done
	//verify middleware chi

	//Login(user,password) -> AuthToken mutation
	//flutter side => save authtoken is shared preference or secure storage
	//add to header for next requests

	//Log stuffs
	l := utils.NewLoggerService(cfg)

	//JWT stuffs
	a := jwt.NewAuthService(cfg)

	// User service stuffs
	u := user.NewUserService(s, l, a, &t, m, d)

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

	// token := "e51ea03b-4eea-4db8-b81e-81365d4350e0"
	// verified, err := u.VerifyUserToken(token)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(verified)

}

