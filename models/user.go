package models

type User struct {
	Id       int    `json:"id"`
	ChatId   int    `json:"chatid"`
	UserName string `json:"username"`
}
