package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"tgBot/mods"

	"github.com/spf13/viper"
)

func main() {
	// https://core.telegram.org/bots/api#video

	// Инициализация конфига (токенов)
	err := mods.InitConfig()
	if err != nil {
		fmt.Println("Config error: ", err)
		return
	}

	// Url бота для отправки и приёма сообщений
	botUrl := "https://api.telegram.org/bot" + viper.GetString("token")
	offSet := 0

	// Цикл работы приложения
	for {

		// Получение апдейтов
		updates, err := getUpdates(botUrl, offSet)
		if err != nil {
			fmt.Println("Something went wrong: ", err)
		}

		// Обработка апдейтов
		for _, update := range updates {
			respond(botUrl, update)
			offSet = update.UpdateId + 1
		}

		// Вывод в консоль для тестов
		fmt.Println(updates)
	}
}

// Функция получения апдейтов
func getUpdates(botUrl string, offset int) ([]mods.Update, error) {

	// Rest запрос для получения апдейтов
	resp, err := http.Get(botUrl + "/getUpdates?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Запись и обработка полученных данных
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse mods.TelegramResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}

	return restResponse.Result, nil
}

//	Функция генерации ответа
func respond(botUrl string, update mods.Update) error {

	if len(update.Message.Text) <= 0 {
		return nil
	}

	switch update.Message.Text {
	case "/start", "/help":
		mods.Help(botUrl, update)
		return nil
	}

	num, err := strconv.Atoi(update.Message.Text)
	if err != nil {
		mods.Ball8(botUrl, update)
		return nil
	}

	mods.CurrentBall8(botUrl, update, num)

	return nil
}
