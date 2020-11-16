package telegram

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/doublegrey/lotus/api/app/apps"
	"github.com/doublegrey/lotus/api/app/logs"
	"github.com/doublegrey/lotus/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	restartChan = make(chan bool)
)

// Start creates new tg bot instance and listens for new events
func Start() {
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		result := utils.DB.Collection("integrations").FindOne(ctx, bson.M{"name": "telegram"})
		var settings Settings
		err := result.Decode(&settings)
		if err != nil {
			log.Println(err)
		}
		bot, err := tgbotapi.NewBotAPI(settings.Token)
		if err != nil {
			log.Println(err)
		}
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates, err := bot.GetUpdatesChan(u)
		for {
			select {
			case event := <-logs.Events:
				for _, user := range settings.Users {
					for _, app := range user.Apps {
						if app.ID == event.App {
							a, exists := apps.Cache.Load(app.ID)
							if exists {
								var text = fmt.Sprintf("%s\n=-=-=-=-=-=-=-=\n", a.(apps.App).Name)
								for key, value := range event.Data {
									text += fmt.Sprintf("* %s: %v\n", key, value)
								}
								msg := tgbotapi.NewMessage(int64(user.Chat), text)
								bot.Send(msg)
							}
						}
					}
				}
			case <-restartChan:
				break
			case update := <-updates:
				if update.Message == nil {
					continue
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprint(update.Message.Chat.ID))
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}

		}

	}
}
