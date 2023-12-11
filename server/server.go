package server

import (
	"encoding/json"
	"fmt"
	"github/Danila331/testlyceumbot/models"
	"github/Danila331/testlyceumbot/pkg"
	"io"
	"log"
	"net/http"
	"strconv"

	_ "modernc.org/sqlite"
)

// Ссылка на telegram api bot
const BaseUrl = "https://api.telegram.org/bot"

// Функция для работы с update
func GetInfoUpdate(botUrl string, offset int) ([]models.Update, error) {

	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))

	if err != nil {
		return []models.Update{}, err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	var Result models.RestResponse

	err = json.Unmarshal(body, &Result)

	if err != nil {
		return []models.Update{}, err
	}

	return Result.Result, nil
}

// Функция запуска сервера
func StartServer(bottoken string) {
	botUrl := BaseUrl + bottoken
	offset := 0
	for {
		updates, err := GetInfoUpdate(botUrl, offset)
		if err != nil {
			log.Println("Ошибка работы функции GetInfoUpdate")
		}
		for _, update := range updates {
			fmt.Println(update.Message.Text)
			offset = update.UpadateId + 1
			if update.Message.Text == "/start" {
				err := pkg.HandlerSendMessageComandStart(botUrl, update)
				if err != nil {
					log.Println("Ошибка работы функции SendMessageComandStart")
				}
			} else if pkg.ValidUrl(update.Message.Text) {
				err := pkg.HandlerGetQrCode(botUrl, update)
				if err != nil {
					log.Println("Ошибка работы функции GetQrCode")
				}
			} else {
				err := pkg.HandlerWrongMessage(botUrl, update)
				if err != nil {
					log.Println("Ошибка работы функции WrongMessage")
				}
			}
		}
	}
}
