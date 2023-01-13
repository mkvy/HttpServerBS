package service

import (
	"encoding/json"
)

type Service interface {
	Create(json.RawMessage, string) (string, string)
	Update(json.RawMessage, string, string) string
	Delete(string, string) string
	GetById(string, string, string) (json.RawMessage, string)
	GetByParameters(string, string, string) (json.RawMessage, string)
}
