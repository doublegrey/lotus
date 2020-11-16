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
	"go.mongodb.org/mongo-driver/mongo"
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
			if err == mongo.ErrNoDocuments {
				return
			}
		}
		bot, err := tgbotapi.NewBotAPI(settings.Token)
		if err != nil {
			log.Println(err)
			return
		}
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates, err := bot.GetUpdatesChan(u)
		for {
			select {
			case event := <-logs.Events:
				for _, user := range settings.Users {
					a, exists := apps.Cache.Load(event.App)
					if exists {
						var text = fmt.Sprintf("%s\n=-=-=-=-=-=-=-=\n", a.(apps.App).Name)
						for key, value := range event.Data {
							text += fmt.Sprintf("* %s: %v\n", key, value)
						}
						msg := tgbotapi.NewMessage(int64(user.Chat), text)
						bot.Send(msg)
					}
				}
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
