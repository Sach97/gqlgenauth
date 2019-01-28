package main

import (
	gcontext "github.com/Sach97/gqlgenauth/auth/context"
	db "github.com/Sach97/gqlgenauth/auth/db"
)

func main() {

	// // Token stuffs
	// RedisClient := tokenizer.NewRedisClient()
	// client := tokenizer.Tokenizer{RedisClient}

	// token, err := client.GenerateString()
	// fmt.Println(token)
	// if err != nil {
	// 	panic(err)
	// }

	// value, err := client.GetToken(token)
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
	// }
	// client := mailer.NewMailer(cfg)
	// to := []string{"sacha.arbonel@hotmail.fr"}

	// msg := []byte("To: recipient@example.net\r\n" +

	// 	"Subject: testing with get env!\r\n" +

	// 	"\r\n" +

	// 	"This is the email body.\r\n")
	// client.SendEmail("sacha.arbonel@hotmail.fr", to, msg)

	// // Firebase stuffs
	// apiKey := "AIzaSyCZz285uQR6-XDAs6pdINuN8y73RO6kGf4"
	// androidInfo := deeplinker.AndroidInfo{AndroidPackageName: "home.sacha.firebasetest"}
	// dynamicLinkInfo := deeplinker.DynamicLinkInfo{DomainURIPrefix: "https://minorys.page.link", Link: "https://ecstatic-heisenberg-ea2789.netlify.com/confirmation", AndroidInfo: &androidInfo}

	// payload := deeplinker.Payload{DynamicLinkInfo: &dynamicLinkInfo}
	// firebase := deeplinker.NewFireBaseClient(apiKey)
	// firebase.GetDynamicLink(&payload)

	// // DB Stuffs
	config := gcontext.LoadConfig(".")

	postgres := db.Strategy(db.Postgres{})

	client, err := postgres.OpenDB(config)
	if err != nil {
		panic(err)
	}
	client.Ping()
}
