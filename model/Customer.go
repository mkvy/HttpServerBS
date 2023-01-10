package model

import "time"

type Customer struct {
	Surname     string    `json:"surname" mapstructure:"surname" validate:"required"`
	Firstname   string    `json:"firstname" mapstructure:"firstname" validate:"required"`
	Patronym    string    `json:"patronym" mapstructure:"patronym"validate:"required"`
	Age         string    `json:"age" mapstructure:"age"`
	DateCreated time.Time `json:"date_created" mapstructure:"date_created" validate:"required"`
}
