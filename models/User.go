package models

type User struct {
	ID       uint
	Username string
	Email    string
	Password string
	Role     string
	Todos    []Todo
}
