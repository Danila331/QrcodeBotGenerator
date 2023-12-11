package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github/Danila331/testlyceumbot/models"
	"github/Danila331/testlyceumbot/store"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

var Count = 0

func HandlerSendMessageComandStart(boturl string, update models.Update) error {
	var BotMessage models.BotMessage

	ChatId := update.Message.Chat.ChatId
	UserName := update.Message.Chat.UserName
	BotMessage.ChatId = ChatId
	BotMessage.ParseMode = "HTML"
	BotMessage.Text = "Далеко-далеко, во вселенной Star Wars, бот-император смотрит на вас с доброй улыбкой и складывает ушки. 👽\n\nПриветствую, путник! Я - твой верный помощник в создании QR-кодов по ссылке. Просто отправь мне ссылку, и я превращу ее в магический QR-код! 🚀✨ \n\nНапример, отправь мне ссылку на сайт, например: https://www.example.com, и я тут же сотворю для тебя QR-код! 🌐💫"
	data, err := json.Marshal(BotMessage)

	if err != nil {
		return err
	}

	database, err := store.NewDB("./server/telegram.db")

	if err != nil {
		database.Close()
		return err
	}

	user, err := database.SearchByChatid(ChatId)
	fmt.Println(user.ChatId)
	if user.ChatId == 0 {
		err = database.CreateUser(ChatId, UserName)
		if err != nil {
			database.Close()
			return err
		}
	}

	database.Close()

	_, err = http.Post(boturl+"/sendMessage", "application/json", bytes.NewBuffer(data))

	if err != nil {
		return err
	}

	return nil
}

func HandlerWrongMessage(boturl string, update models.Update) error {
	var BotMessage models.BotMessage

	ChatId := update.Message.Chat.ChatId
	BotMessage.ChatId = ChatId
	BotMessage.Text = fmt.Sprintf("Тьма окружает нас, молодой джедай. Ты отправил мне нечто странное, не похожее на ссылку. Может быть, это сообщение из другой галактики? Повтори попытку и отправь мне ссылку, а я обязательно помогу! 🌌✨\n\nМожет у тебя возникли другие вопросы? Я всегда готов помочь, как только почувствую твою силу! Пусть с тобой всегда будет Сила! 💫✨")
	BotMessage.ParseMode = "HTML"
	data, err := json.Marshal(BotMessage)

	if err != nil {
		return err
	}

	_, err = http.Post(boturl+"/sendMessage", "application/json", bytes.NewBuffer(data))

	if err != nil {
		return err
	}

	return nil
}

// http://qrcoder.ru/code/?https%3A%2F%2Ft.me%2Fturbogolang&10&0
func HandlerGetQrCode(boturl string, update models.Update) error {
	filename, err := GeneratorQr(update.Message.Text)

	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("photo", file.Name())

	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)

	caption := fmt.Sprintf("Молодой падаван! 😊 Ты попросил мою помощь в создании QR-кода, и я исполнил твое желание. Вот твой QR-код, он наполнен силой, готов помочь тебе в твоих приключениях. Помни, сила всегда с тобой! И пусть Сопровождает тебя Сила, молодой джедай. 🌌✨")
	_ = writer.WriteField("chat_id", strconv.Itoa(update.Message.Chat.ChatId))
	_ = writer.WriteField("caption", caption)

	err = writer.Close()

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", boturl+"/sendPhoto", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file.Close()
	fmt.Println(file.Name())
	err = os.Remove(filename)
	if err != nil {
		fmt.Printf("%e", err)
		return err
	}
	return nil
}
