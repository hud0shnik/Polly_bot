package mods

import (
	"math/rand"
	"time"

	"github.com/spf13/viper"
)

// Структуры для работы с Telegram API

type TelegramResponse struct {
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatId int `json:"id"`
}

// Функция помощи пользователю
func Help(botUrl string, update Update) {
	SendMsg(botUrl, update, "Просто задай свой вопрос, я на него отвечу")
}

// Функция генерации псевдослучайных чисел
func Random(n int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(n)
}

// Функция генерации псевдослучайных ответов
func Ball8(botUrl string, update Update) {

	answers := []string{
		"DQACAgIAAxkBAAMXYws-dMDsuuw-DqYOblkX7RU7MS0AAtkeAALyalFItN_JGJbMnZ4pBA",
	}

	SendVn(botUrl, update, SendVideoNote{
		ChatId:    update.Message.Chat.ChatId,
		VideoNote: answers[Random(len(answers))],
	})

}

// Функция инициализации конфига (всех токенов)
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
