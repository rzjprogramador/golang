package composicao

import "github.com/rzj/helpers"

type User struct {
	ID    string
	Name  string
	Idade int
}

func CreateUser(u User) *User {
	var validUser = User{
		ID:    u.ID,
		Name:  helpers.MinString(u.Name, 2),
		Idade: helpers.MinAge(u.Idade, 18),
	}
	return &validUser
}
