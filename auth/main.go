package main

import (
	"github.com/Sach97/gqlgenauth/auth/utils"
)

func main() {

	client := utils.New()

	// token, err := client.GenerateString()
	// if err != nil {
	// 	fmt.Println(token)
	// }

	token := "abed1971-3fd9-4319-a2d4-d6fc4e5943bd"
	utils.GetToken(token)

}
