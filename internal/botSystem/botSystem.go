package botSystem

import (
	"feedback/internal/service"
	"feedback/internal/types"
	"feedback/pkg/logging"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

var feedbackCh chan types.FeedBackTG

func NewBotSystem(s *service.Service) {
	feedbackCh = make(chan types.FeedBackTG)
	RunBot(s, feedbackCh)
}

func RunBot(s *service.Service, ch chan types.FeedBackTG) {
	logger := logging.GetLogger()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		logger.Fatalf("Invalid token: %s", err)
	}
	bot.Debug = false
	logger.Printf("Authorized on account %s", bot.Self.UserName)
	var ucfg = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updates := bot.GetUpdatesChan(ucfg)
	for {
		select {
		case update := <-updates:
			BotLogicCommands(update, s)
		case r := <-ch:
			msg := tgbotapi.NewMessage(int64(r.ChatId), Response(r))
			_, err := bot.Send(msg)
			if err != nil {
				log.Println(err)
			}
			if r.PhotoPath != nil {
				for _, i := range r.PhotoPath {
					ph := tgbotapi.NewPhoto(int64(r.ChatId), tgbotapi.FilePath(i))
					_, err := bot.Send(ph)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
}

func Response(r types.FeedBackTG) string {
	resp := fmt.Sprintf("Новый отзыв:\nID: %d\nГород: %s\nТекст: %s", r.FeedbackId, r.City, r.Massage)
	return resp
}
