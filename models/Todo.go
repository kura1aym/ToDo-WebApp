package models

type Todo struct {
	ID     uint
	Text   string
	Done   bool
	UserID uint
}
