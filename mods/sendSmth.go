package mods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Структуры для отправки сообщений, стикеров и картинок

type SendMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
}

type SendSticker struct {
	ChatId     int    `json:"chat_id"`
	StickerUrl string `json:"sticker"`
}

type SendPhoto struct {
	ChatId   int    `json:"chat_id"`
	PhotoUrl string `json:"photo"`
	Caption  string `json:"caption"`
}

type SendVideoNote struct {
	ChatId    int    `json:"chat_id"`
	VideoNote string `json:"video_note"`
}

// Функции отправки сообщений, стикеров и картинок
// Отправка сообщения
func SendMsg(botUrl string, update Update, msg string) error {

	// Формирование сообщения
	botMessage := SendMessage{
		ChatId: update.Message.Chat.ChatId,
		Text:   msg,
	}
	buf, err := json.Marshal(botMessage)
	if err != nil {
		fmt.Println("Marshal json error: ", err)
		return err
	}

	// Отправка сообщения
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		fmt.Println("SendMessage method error: ", err)
		return err
	}
	return nil
}

// Функция отправки кружков
func SendVn(botUrl string, update Update, vn SendVideoNote) error {

	// Формирование кружка
	buf, err := json.Marshal(vn)
	if err != nil {
		fmt.Println("Marshal json error: ", err)
		return err
	}

	// Отправка кружка
	_, err = http.Post(botUrl+"/sendVideoNote", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		fmt.Println("SendVn method error: ", err)
		return err
	}
	return nil
}
