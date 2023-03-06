package models

type User struct {
	Id          int    `json:"id"`
	Login       string `json:"login"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	PhoneNumber string `json:"phone_number"`
}
