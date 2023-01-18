package customerrepo

import (
	"github.com/mkvy/HttpServerBS/shared/model"
)

type CustomerRepository interface {
	Create(model.Customer) (string, error)
	Update(model.Customer, string) error
	Delete(string) error
	GetById(string) (model.Customer, error)
	GetBySurname(string) (model.Customer, error)
}
