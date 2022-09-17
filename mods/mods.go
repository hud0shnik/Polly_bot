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
		"DQACAgIAAxkBAAMhYw0K7ZAuB0VxmZbKBZ4prZAIhI0AAtweAALyalFIp-xVoUVp4FIpBA",
		"DQACAgIAAxkBAAMiYw0LTq9N-K0ex0TYyHNv-EYV-e8AAuMeAALyalFIN8FCgUZLeZEpBA",
		"DQACAgIAAxkBAAMjYw0LeGSyR5tRJuErK9KE5DSvjlgAAvYeAALyalFImnfqhnwzzVMpBA",
		"DQACAgIAAxkBAAMkYw0Loy_3vQmZC2XbB4meiMmYGo8AAgQfAALyalFIckmm5I3zbKApBA",
		"DQACAgIAAxkBAAPVYySqb97Qh3kYroc4J9k2DeUbBhIAAuMbAAIeIiBJVFx8yqrrzCIpBA",
		"DQACAgIAAxkBAAPWYySraHtQyul6d9HLRQkedOV9m7IAAvIbAAIeIiBJR4UwrFf4VSkpBA",
		"DQACAgIAAxkBAAPXYySr1Sp9STwGRfla-PvpA5hkgToAAvUbAAIeIiBJwX7HTbJ-iywpBA",
		"DQACAgIAAxkBAAPYYySr-4iT6c2MYVuH3xtCy1B5NVoAAvsbAAIeIiBJxS5Uvna4d24pBA",
		"DQACAgIAAxkBAAIBZmMmEUyqmWgeVDC7_Io21aSLSbdnAAJaJAAC7j8pSQYbftcmtXnsKQQ",
		"DQACAgIAAxkBAAIBZWMmEUQ5rpXhNcpE-toyC-awgVE9AAI5JAAC7j8pScggvIVqF3VhKQQ",
	}

	SendVn(botUrl, update, SendVideoNote{
		ChatId:    update.Message.Chat.ChatId,
		VideoNote: answers[Random(len(answers))],
	})

}

// Функция заказа кружков
func CurrentBall8(botUrl string, update Update, num int) {

	answers := []string{
		"DQACAgIAAxkBAAMhYw0K7ZAuB0VxmZbKBZ4prZAIhI0AAtweAALyalFIp-xVoUVp4FIpBA",
		"DQACAgIAAxkBAAMiYw0LTq9N-K0ex0TYyHNv-EYV-e8AAuMeAALyalFIN8FCgUZLeZEpBA",
		"DQACAgIAAxkBAAMjYw0LeGSyR5tRJuErK9KE5DSvjlgAAvYeAALyalFImnfqhnwzzVMpBA",
		"DQACAgIAAxkBAAMkYw0Loy_3vQmZC2XbB4meiMmYGo8AAgQfAALyalFIckmm5I3zbKApBA",
		"DQACAgIAAxkBAAPVYySqb97Qh3kYroc4J9k2DeUbBhIAAuMbAAIeIiBJVFx8yqrrzCIpBA",
		"DQACAgIAAxkBAAPWYySraHtQyul6d9HLRQkedOV9m7IAAvIbAAIeIiBJR4UwrFf4VSkpBA",
		"DQACAgIAAxkBAAPXYySr1Sp9STwGRfla-PvpA5hkgToAAvUbAAIeIiBJwX7HTbJ-iywpBA",
		"DQACAgIAAxkBAAPYYySr-4iT6c2MYVuH3xtCy1B5NVoAAvsbAAIeIiBJxS5Uvna4d24pBA",
		"DQACAgIAAxkBAAIBZmMmEUyqmWgeVDC7_Io21aSLSbdnAAJaJAAC7j8pSQYbftcmtXnsKQQ",
		"DQACAgIAAxkBAAIBZWMmEUQ5rpXhNcpE-toyC-awgVE9AAI5JAAC7j8pScggvIVqF3VhKQQ",
	}

	// Проверка на доступность кружка
	if num > 0 && num <= len(answers) {
		SendVn(botUrl, update, SendVideoNote{
			ChatId:    update.Message.Chat.ChatId,
			VideoNote: answers[num-1],
		})
	} else {
		SendMsg(botUrl, update, "У меня нет кружка №"+update.Message.Text)
	}

}

// Приветствие и хелпа
func Help(botUrl string, update Update) {
	SendMsg(botUrl, update, "Привет, ты можешь посмотреть все кружки, для этого нужно просто прислать его номер (от 1 до 9). Ну или просто задай мне любой вопрос.")
}

// Функция инициализации конфига (всех токенов)
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
