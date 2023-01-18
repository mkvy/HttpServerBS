package model

import "time"

type Customer struct {
	Surname     string     `json:"surname,omitempty" validate:"required"`
	Firstname   string     `json:"firstname,omitempty" validate:"required"`
	Patronym    string     `json:"patronym,omitempty" validate:"required"`
	Age         string     `json:"age,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
}
