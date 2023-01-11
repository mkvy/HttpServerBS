package customerrepo

import (
	"github.com/mkvy/HttpServerBS/model"
)

type CustomerRepository interface {
	Create(model.Customer) error
	Update(model.Customer, string) error
	Delete(string) error
	GetById(string, string) (model.Customer, error)
}
