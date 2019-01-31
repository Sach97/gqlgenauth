package main

import (
	"bytes"
	"fmt"
	"html/template"
)

type EmailMessage struct {
	ConfirmationUrl string
}

// func FormatResponse(data []byte, templ string) string {
// 	jsonData := fmt.Sprintf("%s", data)

// 	t := template.Must(template.New("").Parse(templ))

// 	m := map[string]interface{}{}
// 	if err := json.Unmarshal([]byte(jsonData), &m); err != nil {
// 		panic(err)
// 	}

// 	var buff bytes.Buffer
// 	t.Execute(&buff, m)
// 	body := buff.String()
// 	fmt.Println(body)
// }
type Person struct {
	UserName string
}

func main() {
	// t := template.New("fieldname example")
	// t, _ = t.Parse("hello {{.UserName}}!")
	// p := Person{UserName: "Astaxie"}
	// t.Execute(os.Stdout, p)

	// // Context Stuffs
	//cfg := context.LoadConfig(".")

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

	t := template.Must(template.ParseFiles("confirmation.html"))
	p := EmailMessage{
		ConfirmationUrl: "https://ecstatic-heisenberg-ea2789.netlify.com/confirmation",
	}
	var buff bytes.Buffer
	t.Execute(&buff, p)
	body := buff.String()
	fmt.Println(body)

	// client := mailer.NewMailer(cfg)
	// to := []string{"sacha.arbonel@hotmail.fr"}
	// recipients := "recipient@example.ne"
	// subject := "testing with func NewMsg! inputs"
	// sender := "sacha.arbonel@hotmail.fr"
	// inputs := mailer.Inputs{
	// 	Recipients: recipients,
	// 	Subject:    subject,
	// 	Body:       body,
	// 	Sender:     sender,
	// 	To:         to,
	// }
	// msg := client.NewMessage(inputs)
	// client.SendEmail(msg)

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
