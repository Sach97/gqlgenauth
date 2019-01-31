package main

import (
	"fmt"

	"github.com/Sach97/gqlgenauth/auth/context"
	"github.com/Sach97/gqlgenauth/auth/mailer"
)

type Inputs struct {
	Recipients string
	Subject    string
	Body       string
}

// func NewMessage(inputs Inputs)[]byte{
// 	msg := []byte(fmt.Sprintf("To: %s\r\n", recipients) +

// 		fmt.Sprintf("Subject: %s\r\n", subject) +

// 		"\r\n" +

// 		fmt.Sprintf("%s\r\n", body))
// }

func NewMessage(inputs Inputs) []byte {
	msg := []byte(fmt.Sprintf("To: %s\r\n", inputs.Recipients) +

		fmt.Sprintf("Subject: %s\r\n", inputs.Subject) +

		"\r\n" +

		fmt.Sprintf("%s\r\n", inputs.Body))
	return msg
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

	client := mailer.NewMailer(cfg)
	to := []string{"sacha.arbonel@hotmail.fr"}
	recipients := "recipient@example.ne"
	subject := "testing with func NewMsg!"
	body := "This is the email body from newmsg"

	inputs := Inputs{
		Recipients: recipients,
		Subject:    subject,
		Body:       body,
	}
	msg := NewMessage(inputs)
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
