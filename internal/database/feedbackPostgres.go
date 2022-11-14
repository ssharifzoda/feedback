package database

import (
	"feedback/internal/types"
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
	var countries []types.Countries
	if err := f.conn.Find(&countries); err.Error != nil {
		return nil, err.Error
	}
	return countries, nil
}

func (f *FeedbackPostgres) GetCountryCities(countryId int) ([]types.Cities, error) {
	var cities []types.Cities
	query := fmt.Sprintf("select ct.id, ct.name from cities as ct where country_id = ?")
	if err := f.conn.Raw(query, countryId).Scan(&cities); err != nil {
		return cities, err.Error
	}
	return cities, nil
}

func (f *FeedbackPostgres) CreateFeedback(fd *types.FeedBacks) error {
	query := fmt.Sprintf("insert into feedbacks(user_id, city_id, massage, photo)\n" +
		"values (?,?,?,?);")
	err := f.conn.Exec(query, fd.UserId, fd.CityId, fd.Massage, fd.Photo)
	return err.Error
}
