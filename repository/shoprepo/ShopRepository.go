package shoprepo

import "github.com/mkvy/HttpServerBS/model"

type ShopRepository interface {
	Create(model.Shop) error
	Update(model.Shop, string) error
	Delete(string) error
	GetById(string) (model.Shop, error)
}
