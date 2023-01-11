package service

import (
	"encoding/json"
)

type Service interface {
	Create(json.RawMessage, string) error
	Update(json.RawMessage, string, string) error
	Delete(string, string) error
	GetById(string, string, string) json.RawMessage
}
