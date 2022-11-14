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
	query := fmt.Sprintf("select c.id, c.name from country_cities as cc inner join cities c on c.id = cc.city_id where country_id = ?")
	if err := f.conn.Raw(query, countryId).Scan(&cities); err != nil {
		return cities, err.Error
	}
	return cities, nil
}

func (f *FeedbackPostgres) GetAllLocalServices(cityID int) ([]types.Services, error) {
	var services []types.Services
	query := fmt.Sprintf("select s.id, s.name, s.description from local_services as ls inner join country_cities cc on ls.city_id = cc.city_id inner join services s on s.id = ls.service_id where ls.city_id = ?")
	if err := f.conn.Raw(query, cityID).Scan(&services); err != nil {
		return services, err.Error
	}
	return services, nil
}

func (f *FeedbackPostgres) CreateFeedback(fd *types.FeedBacks) error {
	query := fmt.Sprintf("insert into feedbacks(user_id, country_id, city_id, massage, photo)\n" +
		"values (?,?,?,?,?);")
	err := f.conn.Exec(query, fd.UserId, fd.CountryId, fd.CityId, fd.Massage, fd.Photo)
	return err.Error
}
