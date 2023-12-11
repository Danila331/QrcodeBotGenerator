package models

type BotPicture struct {
	Photo   string `json:"photo"`
	ChatId  int    `json:"chat_id"`
	Caption string `json:"caption"`
}
