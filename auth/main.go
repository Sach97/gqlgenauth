package main

import (
	"github.com/Sach97/gqlgenauth/auth/context"
	"github.com/Sach97/gqlgenauth/auth/mailer"
)

func main() {

	// // Context Stuffs
	cfg := context.LoadConfig(".")

	// // Token stuffs
	// RedisClient := tokenizer.NewRedisClient()
	// client := tokenizer.Tokenizer{RedisClient}

	// token, err := client.GenerateToken("userid")
	// fmt.Println(token)
	// if err != nil {
	// 	panic(err)
	// }

	// value, err := client.GetUserID(token)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(value)

	// // Mail stuffs

	client := mailer.NewMailer(cfg)
	to := []string{"sacha.arbonel@hotmail.fr"}

	msg := []byte("To: recipient@example.net\r\n" +

		"Subject: testing with get env!\r\n" +

		"\r\n" +

		"This is the email body.\r\n")
	client.SendEmail("sacha.arbonel@hotmail.fr", to, msg)

	// // Firebase STUFFS

	// firebase := deeplinker.NewFireBaseClient(cfg)

	// link, _ := firebase.GetDynamicLink("randomstring", true)
	// fmt.Println(link)

	//DB STUFFS

	// sql := db.Strategy(db.DriverSQL{Name: "postgres"})

	// client, err := sql.OpenDB(config)
	// if err != nil {
	// 	panic(err)
	// }

	// client.
	// userService := service.NewUserService(db, roleService, log)
	// ctx := context.Background()

	//Signup(email,password) mutation
	//create user
	//send email confirmation

	//ConfirmUser(token) query
	//Get userid from token
	// Verify if user exists from userid
	//sets user as confirmed in database from userid
	//send boolean isConfirmed

	//Login(user,password) -> AuthToken mutation
	//flutter side => save authtoken is shared preference or secure storage
	//add to header for next requests

}
