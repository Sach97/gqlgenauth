package main

import (
	deeplinking "github.com/Sach97/gqlgenauth/auth/deeplinking"
)

func main() {

	// client := utils.NewTokenizer()

	// token, err := client.GenerateString()
	// if err != nil {
	// 	panic(err)
	// }

	// value, err := client.GetToken(token)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(value)

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

	apiKey := "AIzaSyCZz285uQR6-XDAs6pdINuN8y73RO6kGf4"
	androidInfo := deeplinking.AndroidInfo{AndroidPackageName: "home.sacha.firebasetest"}
	dynamicLinkInfo := deeplinking.DynamicLinkInfo{DomainURIPrefix: "https://minorys.page.link", Link: "https://ecstatic-heisenberg-ea2789.netlify.com/confirmation", AndroidInfo: &androidInfo}

	payload := deeplinking.Payload{DynamicLinkInfo: &dynamicLinkInfo}
	firebase := deeplinking.NewFireBaseClient(apiKey)
	firebase.GetDynamicLink(&payload)
}
