package shoprepo

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mkvy/HttpServerBS/model"
	"log"
)

type MemoryShopRepo struct {
	//todo config
	storage map[string]model.Shop
}

func NewMemoryShopRepo() *MemoryShopRepo {
	return &MemoryShopRepo{storage: make(map[string]model.Shop)}
}

func (repo *MemoryShopRepo) Create(data model.Shop) error {
	id := uuid.New()
	repo.storage[id.String()] = data
	log.Println("Created record with id: " + id.String())
	return nil
}

func (repo *MemoryShopRepo) GetById(id string) (model.Shop, error) {
	val, ok := repo.storage[id]
	log.Println("Getting record with id: " + id)
	if !ok {
		return model.Shop{}, errors.New("Not found by ID")
	}
	return val, nil
}

func (repo *MemoryShopRepo) Update(data model.Shop, id string) error {
	record, err := repo.GetById(id)
	if err != nil {
		return err
	}
	if data.Name != "" {
		record.Name = data.Name
	}
	if data.WorkStatus != nil {
		*record.WorkStatus = *data.WorkStatus
	}
	if data.Address != "" {
		record.Address = data.Address
	}
	if data.Owner != "" {
		record.Owner = data.Owner
	}
	repo.storage[id] = record
	return nil
}

func (repo *MemoryShopRepo) Delete(id string) error {
	if _, ok := repo.storage[id]; ok {
		delete(repo.storage, id)
	} else {
		return errors.New("No record with ID " + id)
	}
	return nil
}
