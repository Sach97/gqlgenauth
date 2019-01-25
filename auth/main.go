package main

import (
	"fmt"

	"github.com/Sach97/gqlgenauth/auth/utils"
)

func main() {

	client := utils.New()

	token, err := client.GenerateString()
	if err != nil {
		panic(err)
	}

	value, err := client.GetToken(token)
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

}
