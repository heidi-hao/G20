package models

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

func NewUser(name, email string, age int) *User {
	return &User{
		Name:  name,
		Email: email,
		Age:   age,
	}
}
