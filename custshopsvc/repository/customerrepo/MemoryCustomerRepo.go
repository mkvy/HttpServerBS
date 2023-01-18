package customerrepo

import (
	"github.com/google/uuid"
	"github.com/mkvy/HttpServerBS/custshopsvc/internal/utils"
	"github.com/mkvy/HttpServerBS/shared/model"
	"log"
	"sync"
	"time"
)

type MemoryCustomerRepo struct {
	storage map[string]model.Customer
	mx      sync.RWMutex
}

func NewMemoryCustomerRepo() *MemoryCustomerRepo {
	return &MemoryCustomerRepo{storage: make(map[string]model.Customer)}
}

func (repo *MemoryCustomerRepo) Create(data model.Customer) (string, error) {
	log.Println("MemoryCustomerRepo Creating record")
	repo.mx.Lock()
	defer repo.mx.Unlock()
	id := uuid.New()
	dateCreated := time.Now()
	data.DateCreated = &dateCreated
	repo.storage[id.String()] = data
	log.Println("Created record with id: " + id.String())
	return id.String(), nil
}

func (repo *MemoryCustomerRepo) GetById(id string) (model.Customer, error) {
	repo.mx.RLock()
	defer repo.mx.RUnlock()
	val, ok := repo.storage[id]
	log.Println("Getting record with id: " + id)
	if !ok {
		return model.Customer{}, utils.ErrNotFound
	}
	return val, nil
}

func (repo *MemoryCustomerRepo) Update(data model.Customer, id string) error {
	record, err := repo.GetById(id)
	repo.mx.Lock()
	defer repo.mx.Unlock()
	if err != nil {
		return err
	}
	if data.Surname != "" {
		record.Surname = data.Surname
	}
	if data.Firstname != "" {
		record.Firstname = data.Firstname
	}
	if data.Patronym != "" {
		record.Patronym = data.Patronym
	}
	if data.Age != "" {
		record.Age = data.Age
	}
	if data.DateCreated != nil {
		*record.DateCreated = *data.DateCreated
	}

	repo.storage[id] = record
	return nil
}

func (repo *MemoryCustomerRepo) Delete(id string) error {
	repo.mx.Lock()
	defer repo.mx.Unlock()
	if _, ok := repo.storage[id]; ok {
		delete(repo.storage, id)
	} else {
		return utils.ErrNotFound
	}
	return nil
}

func (repo *MemoryCustomerRepo) GetBySurname(surname string) (model.Customer, error) {
	repo.mx.RLock()
	defer repo.mx.RUnlock()
	for key, element := range repo.storage {
		if element.Surname == surname {
			log.Println("Found by surname " + surname + " element with id " + key)
			return element, nil
		}
	}
	return model.Customer{}, utils.ErrNotFound
}
