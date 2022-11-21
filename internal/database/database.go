package database

import (
	"feedback/internal/types"
	"gorm.io/gorm"
)

type Feedback interface {
	GetAllCountries() ([]types.Countries, error)
	GetCountryCities(countryId int) ([]types.Cities, error)
	CreateFeedback(feedback *types.Feedbacks) (int, error)
	GetAllFeedback(limit, offset int, term string) ([]types.Feedbacks, error)
	SearchFeedbacks(phoneNumber string, limit, offset int) ([]types.Feedbacks, error)
}

type Bot interface {
	UpdateFeedbackStatus(feedbackId int) error
}

type Database struct {
	Feedback
	Bot
}

func NewDatabase(conn *gorm.DB) *Database {
	return &Database{
		Feedback: NewFeedbackPostgres(conn),
		Bot:      NewBotPostgres(conn),
	}
}
