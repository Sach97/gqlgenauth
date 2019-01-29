package main

import (
	"fmt"

	"github.com/Sach97/gqlgenauth/auth/context"
	"github.com/Sach97/gqlgenauth/auth/deeplinker"
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
	// sendgrid := os.Getenv("SENDGRID")
	// fmt.Println(sendgrid)
	// cfg := mailer.Service{
	// 	"",
	// 	"apikey",
	// 	sendgrid,
	// 	"smtp.sendgrid.net",
	// 	"smtp.sendgrid.net:25",
	// }apiKey := "AIzaSyCZz285uQR6-XDAs6pdINuN8y73RO6kGf4"
	// androidInfo := deeplinker.AndroidInfo{AndroidPackageName: "home.sacha.firebasetest"}
	// dynamicLinkInfo := deeplinker.DynamicLinkInfo{DomainURIPrefix: "https://minorys.page.link", Link: "https://ecstatic-heisenberg-ea2789.netlify.com/confirmation", AndroidInfo: &androidInfo}
	// client := mailer.NewMailer(cfg)
	// to := []string{"sacha.arbonel@hotmail.fr"}

	// msg := []byte("To: recipient@example.net\r\n" +

	// 	"Subject: testing with get env!\r\n" +

	// 	"\r\n" +

	// 	"This is the email body.\r\n")
	// client.SendEmail("sacha.arbonel@hotmail.fr", to, msg)

	// // Firebase stuffs

	firebase := deeplinker.NewFireBaseClient(cfg)

	link, _ := firebase.GetDynamicLink("randomstring", true)
	fmt.Println(link)

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
