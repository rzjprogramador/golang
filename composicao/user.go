package composicao

type User struct {
	ID    string
	Name  string
	Idade int
}

func CreateUser(u User) *User {
	var validUser = User{
		ID:    u.ID,
		Name:  u.Name,
		Idade: u.Idade,
	}
	return &validUser
}
