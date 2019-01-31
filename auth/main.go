package main

import (
	"github.com/Sach97/gqlgenauth/auth/context"
	"github.com/Sach97/gqlgenauth/auth/mailer"
)

type EmailMessage struct {
	ConfirmationUrl string
}

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

	p := EmailMessage{
		ConfirmationUrl: "https://ecstatic-heisenberg-ea2789.netlify.com/confirmation",
	}
	client := mailer.NewMailer(cfg)
	to := []string{"sacha.arbonel@hotmail.fr"}
	recipients := "recipient@example.ne"
	subject := "Confirmation email"
	sender := "sacha.arbonel@hotmail.fr"
	inputs := mailer.Inputs{
		Recipients: recipients,
		Subject:    subject,
		Sender:     sender,
		To:         to,
	}

	client.SendEmailTemplate(inputs, "confirmation", p)

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
