package database

import (
	"feedback/internal/types"
	"gorm.io/gorm"
)

type Feedback interface {
	GetAllCountries() ([]types.Countries, error)
	GetCountryCities(countryId int) ([]types.Cities, error)
	CreateFeedback(feedback *types.FeedBacks) error
}

type Database struct {
	Feedback
}

func NewDatabase(conn *gorm.DB) *Database {
	return &Database{
		Feedback: NewFeedbackPostgres(conn),
	}
}
