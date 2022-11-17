package botSystem

import (
	"feedback/internal/service"
	"feedback/internal/types"
	"feedback/pkg/logging"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"strings"
)

func BotLogicCommands(bot *tgbotapi.BotAPI, update tgbotapi.Update, s *service.Service) {
	logger := logging.GetLogger()
	log.Println(update.Message.Chat.ID)
	text := update.Message.Text
	repl := strings.ReplaceAll(text, " ", "")
	sliceStr := strings.Split(repl, "-")
	for _, val := range sliceStr {
		switch val {
		case "Ok":
			feedbackID, err := strconv.Atoi(sliceStr[1])
			if err != nil {
				ValidateResponseStatus(bot)
				continue
			}
			if err := s.Bot.UpdateFeedbackStatus(feedbackID); err != nil {
				logger.Println(err)
			}
		}
	}
}

func Sender(response types.FeedBackTG) {
	feedbackCh <- response
}

func ValidateResponseStatus(bot *tgbotapi.BotAPI) {
	logger := logging.GetLogger()
	logger.Println("error Atoi func")
	msg := tgbotapi.NewMessage(int64(viper.GetInt("chatid")), "Неправильно введен ID отзыва, проверьте и отправьте ещё раз")
	if _, err := bot.Send(msg); err != nil {
		logger.Println(err)
	}
}
