package server

import (
	"encoding/json"
	"fmt"
	"github.com/mkvy/HttpServerBS/service"
	"log"
	"net/http"
	"strings"
)

type Controller struct {
	service service.Service
}

func NewController(s service.Service) *Controller {
	return &Controller{service: s}
}

// todo: переделать на одинаковые хендлеры, передавать лишь параметром shoprepo/customerrepo тип! повторяешься
// todo: сверить тогда насчет required полей, или опустить в сервис
// todo: получить из json слайс байт и передавать в логику сервиса

// POST /api/v1/shop/
// PATCH /api/v1/shop/{id}
// DELETE /api/v1/shop/{id}
// GET /api/v1/shop/?name={name}
// GET /api/v1/shop?name={name}&field={field}
func (c *Controller) ShopController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// если без <id> запрос
	modelType := "shop"
	if r.URL.Path == "/api/v1/shop/" {
		if r.Method == http.MethodPost {
			c.createHandler(w, r, modelType)
			return
		}
		if r.Method == http.MethodGet {
			// логика на случай если параметры name = или name = или field =
			keys, ok := r.URL.Query()["name"]
			if !ok || len(keys[0]) < 1 {
				log.Println("Url Param 'name' is missing")
			}
			log.Println(keys)
			c.getAllHandler(w, r, modelType)
			return
		}
		http.Error(w, fmt.Sprintf("expect method GET or POST at /api/v1/shop/, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	} else {
		path := strings.Trim(r.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 2 {
			http.Error(w, "expect /api/v1/shop/<id> in task handler", http.StatusBadRequest)
			return
		}
		if len(pathParts) > 4 {
			http.Error(w, "wrong request", http.StatusBadRequest)
			return
		}
		if r.Method == http.MethodPatch {
			c.patchHandler(w, r, modelType)
			return
		}
		if r.Method == http.MethodDelete {
			c.deleteHandler(w, r, modelType)
			return
		}
		if r.Method == http.MethodGet {
			c.getByIDHandler(w, r, modelType)
			return
		}
		http.Error(w, fmt.Sprintf("expect method DELETE or PATCH or GET at /api/v1/shop/<id>, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

}

// POST /api/v1/customer/
// PATCH /api/v1/customer/{id}
// DELETE /api/v1/customer/{id}
// GET /api/v1/customer/?surname={surname}
// GET /api/v1/customer?surname={surname}&field={field}
func (c *Controller) CustController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	modelType := "customer"
	// если без <id> запрос
	if r.URL.Path == "/api/v1/customer/" {
		if r.Method == http.MethodPost {
			c.createHandler(w, r, modelType)
			return
		}
		if r.Method == http.MethodGet {
			// логика на случай если параметры surname = или surname = или field =
			keys, ok := r.URL.Query()["surname"]
			if !ok || len(keys[0]) < 1 {
				log.Println("Url Param 'name' is missing")
			}
			log.Println(keys)
			c.getAllHandler(w, r, modelType)
			return
		}
		http.Error(w, fmt.Sprintf("expect method GET or POST at /api/v1/customer/, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	} else {
		path := strings.Trim(r.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 2 {
			http.Error(w, "expect /api/v1/customer/<id> in task handler", http.StatusBadRequest)
			return
		}
		if len(pathParts) > 4 {
			http.Error(w, "wrong request", http.StatusBadRequest)
			return
		}
		if r.Method == http.MethodPatch {
			c.patchHandler(w, r, modelType)
			return
		}
		if r.Method == http.MethodDelete {
			c.deleteHandler(w, r, modelType)
			return
		}
		if r.Method == http.MethodGet {
			c.getByIDHandler(w, r, modelType)
		}
		http.Error(w, fmt.Sprintf("expect method DELETE or PATCH or GET at /api/v1/customer/<id>, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

}

func (c *Controller) createHandler(w http.ResponseWriter, r *http.Request, model string) {
	var msg json.RawMessage
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.service.Create(msg, model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	return
}

func (c *Controller) patchHandler(w http.ResponseWriter, r *http.Request, model string) {
	path := strings.Trim(r.URL.Path, "/")
	pathParts := strings.Split(path, "/")
	id := pathParts[3]
	var msg json.RawMessage
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Controller sending patch " + model + " with id " + id)
	c.service.Update(msg, id, model)
	return
}

func (c *Controller) deleteHandler(w http.ResponseWriter, r *http.Request, model string) {
	path := strings.Trim(r.URL.Path, "/")
	pathParts := strings.Split(path, "/")
	id := pathParts[3]
	log.Println("Controller sending delete " + model + " with id " + id)
	c.service.Delete(id, model)
	//todo handle error in response
}

func (c *Controller) getAllHandler(w http.ResponseWriter, r *http.Request, model string) {

}

func (c *Controller) getByIDHandler(w http.ResponseWriter, r *http.Request, model string) {
	key, ok := r.URL.Query()["field"]
	if ok {
		fmt.Println("Have a key field: ", key)
	}
	path := strings.Trim(r.URL.Path, "/")
	pathParts := strings.Split(path, "/")
	id := pathParts[3]
	msg := c.service.GetById(id, "", model)
	w.Header().Set("Content-Type", "application/json")
	w.Write(msg)
}
