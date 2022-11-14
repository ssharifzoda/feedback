package service

import (
	"feedback/internal/database"
	"feedback/internal/types"
	"io"
)

type Feedback interface {
	GetAllCountries() ([]types.Countries, error)
	GetCountryCities(countryId int) ([]types.Cities, error)
	ValidateImage(size int64) error
	SaveImage(file io.Reader, fileName string, feedback *types.FeedBacks) (*types.FeedBacks, error)
	CreateFeedback(feedback *types.FeedBacks) error
}

type Service struct {
	Feedback
}

func NewService(db *database.Database, imagPath string) *Service {
	return &Service{
		Feedback: NewFeedbackService(db, imagPath),
	}
}
