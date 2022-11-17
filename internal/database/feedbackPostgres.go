package database

import (
	"feedback/internal/types"
	"feedback/pkg/logging"
	"fmt"
	"gorm.io/gorm"
)

type FeedbackPostgres struct {
	conn *gorm.DB
}

func NewFeedbackPostgres(conn *gorm.DB) *FeedbackPostgres {
	return &FeedbackPostgres{conn: conn}
}
func (f *FeedbackPostgres) GetAllCountries() ([]types.Countries, error) {
	log := logging.GetLogger()
	var countries []types.Countries
	if err := f.conn.Find(&countries); err.Error != nil {
		log.Print(err)
		return nil, err.Error
	}
	return countries, nil
}

func (f *FeedbackPostgres) GetCountryCities(countryId int) ([]types.Cities, error) {
	log := logging.GetLogger()
	var cities []types.Cities
	query := fmt.Sprintf("select ct.id, ct.name from cities as ct where country_id = ?")
	if err := f.conn.Raw(query, countryId).Scan(&cities); err != nil {
		log.Print(err)
		return cities, err.Error
	}
	return cities, nil
}

func (f *FeedbackPostgres) CreateFeedback(fd *types.Feedbacks) (int, error) {
	log := logging.GetLogger()
	result := f.conn.Table("feedbacks").Create(&fd)
	if result.Error != nil {
		log.Print(result.Error)
	}
	id := fd.ID
	return id, result.Error
}
