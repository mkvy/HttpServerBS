package service

import (
	"encoding/json"
)

type Service interface {
	Create(interface{}, string) (string, string)
	Update(interface{}, string, string) string
	Delete(string, string) string
	GetById(string, string, string) (json.RawMessage, string)
	GetByParameters(string, string, string) (json.RawMessage, string)
}
