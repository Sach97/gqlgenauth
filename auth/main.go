package main

import (
	"encoding/base64"
	"fmt"

	"github.com/Sach97/gqlgenauth/auth/context"
	"github.com/Sach97/gqlgenauth/auth/jwt"
)

type EmailMessage struct {
	ConfirmationUrl string
}

func main() {

	// // Context Stuffs
	cfg := context.LoadConfig(".")

	// // Token stuffs
	// RedisClient := tokenizer.NewRedisClient()
	// t := tokenizer.Tokenizer{RedisClient}

	// // Mail stuffs
	//m := mailer.NewMailer(cfg)

	// // Firebase STUFFS
	//d := deeplinker.NewFireBaseClient(cfg)

	//DB STUFFS
	//sql := db.Strategy(db.DriverSQL{Name: "postgres"})

	// s, err := sql.OpenDB(cfg)
	// if err != nil {
	// 	panic(err)
	// }

	//Signup(email,password) mutation done
	//create user done
	//send email confirmation done

	//ConfirmUser(token) query
	//Get userid from token done
	// Verify if user exists from userid done
	//sets user as confirmed in database from userid done
	//send boolean isConfirmed done

	//Login(user,password) -> AuthToken mutation
	//flutter side => save authtoken is shared preference or secure storage
	//add to header for next requests

	//Log stuffs
	//l := utils.NewLoggerService(cfg)

	// User service stuffs
	//u := user.NewUserService(s, l, &t, m, d)

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

	//JWT stuffs
	a := jwt.NewAuthService(cfg)
	customMap := CustomNamespace{
		Sub:   base64.StdEncoding.EncodeToString([]byte("1234567890")),
		Name:  base64.StdEncoding.EncodeToString([]byte("John Doe")),
		Admin: true,
		//Iat:   time.Now().Add(time.Second * *time.Duration(cfg.JWTExpireIn)).Unix(),
		HTTPSHasuraIoJwtClaims: HTTPSHasuraIoJwtClaims{
			XHasuraAllowedRoles: []string{"user", "editor"},
			XHasuraDefaultRole:  "user",
			XHasuraOrgID:        "123",
			XHasuraCustom:       "custom-value",
		},
	}
	token, _ := a.SignJWT(customMap)
	t := []byte(*token)
	fmt.Println(string(t))

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
