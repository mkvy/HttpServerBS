package model

type Shop struct {
	Name       string `json:"name" mapstructure:"name" validate:"required"`
	Address    string `json:"address" mapstructure:"address" validate:"required"`
	WorkStatus *bool  `json:"work_status" mapstructure:"work_status" validate:"required"`
	Owner      string `json:"owner" mapstructure:"owner"`
}
