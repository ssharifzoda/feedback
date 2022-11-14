package types

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

type FeedBacks struct {
	UserId  int    `json:"user_id"`
	CityId  int    `json:"city_id"`
	Massage string `json:"massage"` //maybe chang area name to description
	Photo   string `json:"photo"`
}
