package model

type Shop struct {
	Name       string `json:"name,omitempty" validate:"required"`
	Address    string `json:"address,omitempty" validate:"required"`
	WorkStatus *bool  `json:"work_status,omitempty" validate:"required"`
	Owner      string `json:"owner,omitempty"`
}
