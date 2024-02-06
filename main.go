package main

import (
	"fmt"

	composicao "github.com/rzj/composicao"
)

func main() {
	var requestUser1 = composicao.User{
		ID:    "1",
		Name:  "Reinaldo",
		Idade: 46,
	}
	createUser1 := composicao.CreateUser(requestUser1)

	fmt.Println(createUser1)
}
