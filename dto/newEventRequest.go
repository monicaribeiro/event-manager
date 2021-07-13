package dto

import (
	"time"
)

type NewEventRequest struct {
	Name      string `json:"name"`
	City      string `json:"city"`
	State     string `json:"state"`
	PhotoUrl  string `json:"photoUrl"`
	Datetime  time.Time `json:"datetime"`
}