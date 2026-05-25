package models

type User struct {
	ID         int    `json:"id"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	SecondName string `json:"secondname"`
	Date       string `json:"date"`
	Role       int    `json:"role"`
}
