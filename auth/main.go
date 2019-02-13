package main

import (
	"fmt"

	"github.com/Sach97/gqlgenauth/auth/context"
	"github.com/Sach97/gqlgenauth/auth/db"
	"github.com/Sach97/gqlgenauth/auth/deeplinker"
	"github.com/Sach97/gqlgenauth/auth/jwt"
	"github.com/Sach97/gqlgenauth/auth/mailer"
	"github.com/Sach97/gqlgenauth/auth/model"
	"github.com/Sach97/gqlgenauth/auth/tokenizer"
	"github.com/Sach97/gqlgenauth/auth/user"
	"github.com/Sach97/gqlgenauth/auth/utils"
)

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

	//Message service stuffs
	msg := context.NewMessageService(cfg)

	// User service stuffs
	u := user.NewUserService(msg, s, l, a, &t, m, d)

	credentials := model.UserCredentials{Email: "sacha.arbonel@hotmail.fr", Password: "secretpassword"}
	signup := u.Signup(&credentials)
	fmt.Println(signup)

	// token := "e51ea03b-4eea-4db8-b81e-81365d4350e0"
	// verified, err := u.VerifyUserToken(token)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(verified)

	// token, err := u.Login(&credentials)
	// fmt.Println(token)
	// fmt.Println(err)
}
