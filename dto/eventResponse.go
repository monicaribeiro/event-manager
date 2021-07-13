package dto

import "time"

type EventResponse struct {
	Id        int64 `json:"id"`
	Name      string `json:"name"`
	City      string `json:"city"`
	State     string `json:"state"`
	PhotoUrl  string `json:"photoUrl"`
	Datetime  time.Time `json:"datetime"`
	CreatedOn time.Time `json:"createdOn"`
}
