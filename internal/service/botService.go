package service

import (
	"feedback/internal/database"
)

type BotService struct {
	db database.Bot
}

func NewBotService(db database.Bot) *BotService {
	return &BotService{db: db}
}

func (b *BotService) UpdateFeedbackStatus(feedbackId int) error {
	return b.db.UpdateFeedbackStatus(feedbackId)
}
