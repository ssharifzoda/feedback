package service

import (
	"feedback/internal/database"
	"feedback/internal/types"
	"mime/multipart"
)

type Feedback interface {
	GetAllCountries() ([]types.Countries, error)
	GetCountryCities(countryId int) ([]types.Cities, error)
	ValidateImage(size int64) error
	SaveImage(reader *multipart.Form, feedback *types.Feedbacks) (*types.Feedbacks, error)
	CreateFeedback(feedback *types.Feedbacks) (int, error)
	GetAllFeedbacks(page, limit int, term string) ([]types.Feedbacks, error)
	SearchFeedbacks(phoneNumber string, page, limit int) ([]types.Feedbacks, error)
}

type Bot interface {
	UpdateFeedbackStatus(feedbackId int) error
}

type Service struct {
	Feedback
	Bot
}

func NewService(db *database.Database, imagPath string) *Service {
	return &Service{
		Feedback: NewFeedbackService(db, imagPath),
		Bot:      NewBotService(db),
	}
}
