package types

import "time"

type Countries struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Cities struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Services struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type LocalServices struct {
	Service []Services
}

type Feedbacks struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	CityId    int       `json:"city_id"`
	Massage   string    `json:"massage"` //maybe chang area name to description
	Status    bool      `json:"status"`
	Photo     []string  `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
}

type FeedBackTG struct {
	FeedbackId int
	ChatId     int
	City       string `json:"city_id"`
	Massage    string `json:"massage"` //maybe chang area name to description
	PhotoPath  []string
}
