package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var token = ""

func main() {
	bot, err := tgbotapi.NewBotAPI("5440361925:AAH5b-xTZUAZs9azk3CfY4Qn9BdOGtkXjNA")

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			message := update.Message
			switch update.Message.Text {
			case "/start":
				op := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Hello, %s how can you feel now?", message.From.FirstName))
				bot.Send(op)
			case "/new":
				url := "https://www.boredapi.com/api/activity"
				client := http.Client{}
				req, _ := http.NewRequest("GET", url, nil)

				res, _ := client.Do(req)

				defer res.Body.Close()

				responseData, _ := ioutil.ReadAll(res.Body)

				var responseActivity ResponseActivtiy

				_ = json.Unmarshal(responseData, &responseActivity)

				op := tgbotapi.NewMessage(message.Chat.ID, responseActivity.Activity)
				bot.Send(op)
			}

		}
	}
}

type ResponseActivtiy struct {
	Activity      string `json:"activity"`
	Type          string `json:"type"`
	Participants  int    `json:"participants"`
	Price         int    `json:"price"`
	Link          string `json:"link"`
	Key           string `json:"key"`
	Accessibility int    `json:"accessibility"`
}
