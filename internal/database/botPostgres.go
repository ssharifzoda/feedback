package database

import (
	"feedback/pkg/logging"
	"gorm.io/gorm"
)

type BotPostgres struct {
	conn *gorm.DB
}

func NewBotPostgres(conn *gorm.DB) *BotPostgres {
	return &BotPostgres{conn: conn}
}
func (b *BotPostgres) UpdateFeedbackStatus(feedbackId int) error {
	log := logging.GetLogger()
	if err := b.conn.Table("feedbacks").Where("id = ?", feedbackId).Update("status", true); err.Error != nil {
		log.Println(err)
		return err.Error
	}
	return nil
}
