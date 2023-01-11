package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/mkvy/HttpServerBS/model"
	"github.com/mkvy/HttpServerBS/repository/shoprepo"
	"log"
)

type ServiceImpl struct {
	shopRepo shoprepo.ShopRepository
}

func NewServiceImpl() *ServiceImpl {
	return &ServiceImpl{shopRepo: shoprepo.NewMemoryShopRepo()}
}

func (s *ServiceImpl) Create(jsonData json.RawMessage, modelType string) error {
	if modelType == "shop" {
		var data model.Shop
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			log.Println(err)
			return err
		}
		validate := validator.New()
		err = validate.Struct(data)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Println("Service Impl CreateShop: ", data)
		err = s.shopRepo.Create(data)
		if err != nil {
			log.Println(err)
			return err
		}
	} else if modelType == "customer" {
		var data model.Customer
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			log.Println(err)
		}
		validate := validator.New()
		err = validate.Struct(data)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println(data)
	}
	return nil
}

func (s *ServiceImpl) Update(jsonData json.RawMessage, id string, modelType string) error {
	if modelType == "shop" {
		var data model.Shop
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			log.Println(err)
			return err
		}
		err = s.shopRepo.Update(data, id)
		if err != nil {
			log.Println("UpdateService error ", err)
			return err
		}
	} else if modelType == "customer" {
		var data model.Customer
		err := json.Unmarshal(jsonData, &data)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (s *ServiceImpl) Delete(id string, modelType string) error {
	if modelType == "shop" {
		err := s.shopRepo.Delete(id)
		if err != nil {
			log.Println("ServiceImpl error while deleting ", err)
			return err
		}
	} else if modelType == "customer" {

	}
	return nil
}

func (s *ServiceImpl) GetById(id string, field string, modelType string) json.RawMessage {
	if modelType == "shop" {
		data, err := s.shopRepo.GetById(id)
		if err != nil {
			log.Println(err)
			return nil
		}
		returnData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			return nil
		}
		return returnData
	} else if modelType == "customer" {

	}
	return nil
}
