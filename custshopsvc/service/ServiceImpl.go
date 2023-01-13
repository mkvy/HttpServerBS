package service

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/mkvy/HttpServerBS/custshopsvc/internal/config"
	"github.com/mkvy/HttpServerBS/custshopsvc/internal/utils"
	"github.com/mkvy/HttpServerBS/custshopsvc/model"
	"github.com/mkvy/HttpServerBS/custshopsvc/repository/customerrepo"
	"github.com/mkvy/HttpServerBS/custshopsvc/repository/shoprepo"
	"log"
)

type ServiceImpl struct {
	shopRepo shoprepo.ShopRepository
	custRepo customerrepo.CustomerRepository
}

func NewServiceImpl(cfg config.Config) *ServiceImpl {
	db, err := utils.GetDBConn(cfg)
	if err != nil {
		log.Println("ServiceImpl Error with database")
		panic("Error with database: " + err.Error())
	}
	customerRepo := customerrepo.NewDBCustomerRepository(db)
	shopRepo := shoprepo.NewDBShopRepository(db)
	//todo check err
	return &ServiceImpl{shopRepo: shopRepo, custRepo: customerRepo}
}

// todo maybe возвращать json.RawMessage, как и ниже
func (s *ServiceImpl) Create(jsonData json.RawMessage, modelType string) (string, string) {
	var id string
	if modelType == "shop" {
		var data model.Shop
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			log.Println(err)
			return "", err.Error()
		}
		validate := validator.New()
		err = validate.Struct(data)
		if err != nil {
			log.Println(err)
			return "", err.Error()
		}
		log.Println("Service Impl CreateShop: ", data)
		id, err = s.shopRepo.Create(data)
		if err != nil {
			log.Println(err)
			return "", err.Error()
		}
	} else if modelType == "customer" {
		var data model.Customer
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			log.Println(err)
			return "", err.Error()
		}
		validate := validator.New()
		err = validate.Struct(data)
		if err != nil {
			log.Println(err)
			return "", err.Error()
		}
		log.Println("Service Impl CreateCustomer: ", data)
		id, err = s.custRepo.Create(data)
		if err != nil {
			log.Println(err)
			return "", err.Error()
		}
	}
	return id, ""
}

func (s *ServiceImpl) Update(jsonData json.RawMessage, id string, modelType string) string {
	if modelType == "shop" {
		var data model.Shop
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			log.Println(err)
			return err.Error()
		}
		err = s.shopRepo.Update(data, id)
		if err != nil {
			log.Println("ServiceImpl Shop UpdateService error ", err)
			return err.Error()
		}
	} else if modelType == "customer" {
		var data model.Customer
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			log.Println(err)
			return err.Error()
		}
		err = s.custRepo.Update(data, id)
		if err != nil {
			log.Println("ServiceImpl Customer UpdateService error ", err)
			return err.Error()
		}
	}
	return ""
}

func (s *ServiceImpl) Delete(id string, modelType string) string {
	if modelType == "shop" {
		err := s.shopRepo.Delete(id)
		if err != nil {
			log.Println("ServiceImpl error while deleting Shop "+id+" ", err)
			return err.Error()
		}
	} else if modelType == "customer" {
		err := s.custRepo.Delete(id)
		if err != nil {
			log.Println("ServiceImpl error while deleting Customer id: "+id+" ", err)
			return err.Error()
		}
	}
	return ""
}

func (s *ServiceImpl) GetById(id string, field string, modelType string) (json.RawMessage, string) {
	log.Println("ServiceImpl: Get by id " + id + " in model " + modelType + "field options: " + field)
	if modelType == "shop" {
		data, err := s.shopRepo.GetById(id)
		if err != nil {
			log.Println(err)
			return nil, err.Error()
		}
		if field != "" {
			var shop model.Shop
			switch field {
			case "name":
				shop.Name = data.Name
			case "address":
				shop.Address = data.Address
			case "work_status":
				shop.WorkStatus = data.WorkStatus
			case "owner":
				shop.Owner = data.Owner
			default:
				shop = data
			}
			data = shop
		}
		returnData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			return nil, err.Error()
		}
		return returnData, ""
	} else if modelType == "customer" {
		data, err := s.custRepo.GetById(id)
		if err != nil {
			log.Println(err)
			return nil, err.Error()
		}
		if field != "" {
			var cust model.Customer
			switch field {
			case "surname":
				cust.Surname = data.Surname
			case "firstname":
				cust.Firstname = data.Firstname
			case "patronym":
				cust.Patronym = data.Patronym
			case "age":
				cust.Age = data.Age
			case "date_created":
				cust.DateCreated = data.DateCreated
			default:
				cust = data
			}
			data = cust
		}
		returnData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			return nil, err.Error()
		}
		return returnData, ""
	}
	return nil, ""
}

func (s *ServiceImpl) GetByParameters(searchParam string, field string, modelType string) (json.RawMessage, string) {
	log.Println("ServiceImpl: Get by param " + searchParam + " in model " + modelType + "field options: " + field)
	if modelType == "shop" {
		data, err := s.shopRepo.GetByName(searchParam)
		if err != nil {
			log.Println(err)
			return nil, err.Error()
		}
		if field != "" {
			var shop model.Shop
			switch field {
			case "name":
				shop.Name = data.Name
			case "address":
				shop.Address = data.Address
			case "work_status":
				shop.WorkStatus = data.WorkStatus
			case "owner":
				shop.Owner = data.Owner
			default:
				shop = data
			}
			data = shop
		}
		returnData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			return nil, err.Error()
		}
		return returnData, ""
	} else if modelType == "customer" {
		data, err := s.custRepo.GetBySurname(searchParam)
		if err != nil {
			log.Println(err)
			return nil, err.Error()
		}
		if field != "" {
			var cust model.Customer
			switch field {
			case "surname":
				cust.Surname = data.Surname
			case "firstname":
				cust.Firstname = data.Firstname
			case "patronym":
				cust.Patronym = data.Patronym
			case "age":
				cust.Age = data.Age
			case "date_created":
				cust.DateCreated = data.DateCreated
			default:
				cust = data
			}
			data = cust
		}
		returnData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			return nil, err.Error()
		}
		return returnData, ""
	}
	return nil, ""
}
