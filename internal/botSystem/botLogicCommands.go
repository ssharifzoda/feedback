package botSystem

import (
	"feedback/internal/service"
	"feedback/internal/types"
	"feedback/pkg/logging"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

func BotLogicCommands(update tgbotapi.Update, s *service.Service) {
	logger := logging.GetLogger()
	log.Println(update.Message.Chat.ID)
	text := update.Message.Text
	repl := strings.ReplaceAll(text, " ", "")
	sliceStr := strings.Split(repl, "-")
	for _, val := range sliceStr {
		switch val {
		case "Ok":
			feedbackID, _ := strconv.Atoi(sliceStr[1])
			if err := s.Bot.UpdateFeedbackStatus(feedbackID); err != nil {
				logger.Println(err)
			}
		}
	}
}

func Sender(response types.FeedBackTG) {
	feedbackCh <- response
}
