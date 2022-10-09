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

				responseData, _ := ioutil.ReadAll(res.Body)

				var responseActivity ResponseActivtiy

				_ = json.Unmarshal(responseData, &responseActivity)

				msgText := fmt.Sprintf("Activity : %v\nType : %v\nParticipants : %v\nLink : %v\nPrice : %v\nAccessiblity : %v", responseActivity.Activity, responseActivity.Type, responseActivity.Participants, responseActivity.Link, responseActivity.Price, responseActivity.Accessibility)

				op := tgbotapi.NewMessage(message.Chat.ID, msgText)
				bot.Send(op)

				res.Body.Close()

			case "/document":
				msg := "activity - Description of the queried activity\naccessibility - A factor describing how possible an event is to do with zero being the most accessible [0.0, 1.0]\ntype - Type of the activity [\"education\", \"recreational\", \"social\", \"diy\", \"charity\", \"cooking\", \"relaxation\", \"music\", \"busywork\"]\nparticipants - The number of people that this activity could involve [0, n]\nprice - A factor describing the cost of the event with zero being free [0, 1]"

				op := tgbotapi.NewMessage(message.Chat.ID, msg)
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
