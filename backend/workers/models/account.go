package models

type Account struct {
	Id       int64  `json:"-" db:"id"`
	Name     string `json:"name" db:"name"`
	Surname  string `json:"surname" db:"surname"`
	Email    string `json:"mail" db:"email"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
	Sex      string `json:"sex" db:"sex"`
	Date     string `json:"date" db:"date"`
}
