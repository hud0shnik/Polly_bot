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

// Функция генерации псевдослучайных чисел
func Random(n int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(n)
}

// Функция генерации псевдослучайных ответов
func Ball8(botUrl string, update Update) {

	answers := []string{
		"DQACAgIAAxkBAAMgYw0KxGa_-bdaJ6km1XEAAXDrb_2rAAL2HgAC8mpRSJp36oZ8M81TKQQ",
		"DQACAgIAAxkBAAMhYw0K7ZAuB0VxmZbKBZ4prZAIhI0AAtweAALyalFIp-xVoUVp4FIpBA",
		"DQACAgIAAxkBAAMiYw0LTq9N-K0ex0TYyHNv-EYV-e8AAuMeAALyalFIN8FCgUZLeZEpBA",
		"DQACAgIAAxkBAAMjYw0LeGSyR5tRJuErK9KE5DSvjlgAAvYeAALyalFImnfqhnwzzVMpBA",
		"DQACAgIAAxkBAAMkYw0Loy_3vQmZC2XbB4meiMmYGo8AAgQfAALyalFIckmm5I3zbKApBA",
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
