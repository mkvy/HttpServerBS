package shoprepo

import (
	"github.com/google/uuid"
	"github.com/mkvy/HttpServerBS/internal/utils"
	"github.com/mkvy/HttpServerBS/model"
	"log"
	"sync"
)

type MemoryShopRepo struct {
	storage map[string]model.Shop
	mx      sync.RWMutex
}

func NewMemoryShopRepo() *MemoryShopRepo {
	return &MemoryShopRepo{storage: make(map[string]model.Shop)}
}

func (repo *MemoryShopRepo) Create(data model.Shop) (string, error) {
	log.Println("MemoryCustomerRepo Creating record")
	repo.mx.Lock()
	defer repo.mx.Unlock()
	id := uuid.New()
	repo.storage[id.String()] = data
	log.Println("Created record with id: " + id.String())
	return id.String(), nil
}

func (repo *MemoryShopRepo) GetById(id string) (model.Shop, error) {
	repo.mx.RLock()
	defer repo.mx.RUnlock()
	val, ok := repo.storage[id]
	log.Println("Getting record with id: " + id)
	if !ok {
		return model.Shop{}, utils.ErrNotFound
	}
	return val, nil
}

func (repo *MemoryShopRepo) Update(data model.Shop, id string) error {
	record, err := repo.GetById(id)
	repo.mx.Lock()
	defer repo.mx.Unlock()
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
	repo.mx.Lock()
	defer repo.mx.Unlock()
	if _, ok := repo.storage[id]; ok {
		delete(repo.storage, id)
	} else {
		return utils.ErrNotFound
	}
	return nil
}

func (repo *MemoryShopRepo) GetByName(name string) (model.Shop, error) {
	repo.mx.RLock()
	defer repo.mx.RUnlock()
	for key, element := range repo.storage {
		if element.Name == name {
			log.Println("Found by name " + name + " element with id " + key)
			return element, nil
		}
	}
	return model.Shop{}, utils.ErrNotFound
}
