package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mkvy/HttpServerBS/internal/utils"
	"github.com/mkvy/HttpServerBS/model"
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

// todo: сверить тогда насчет required полей, или опустить в сервис
// todo: получить из json слайс байт и передавать в логику сервиса

// POST /api/v1/shop/
// PATCH /api/v1/shop/{id}
// DELETE /api/v1/shop/{id}
// GET /api/v1/shop/?name={name}
// GET /api/v1/shop?name={name}&field={field}
// GET /api/v1/shop/{id}?field={field}
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
			c.getByParameters(w, r, modelType)
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
// GET /api/v1/customer/{id}?field={field}
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
			c.getByParameters(w, r, modelType)
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
			return
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
	id, err := c.service.Create(msg, model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"id": "%s"}`, id)
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
	err := c.service.Delete(id, model)
	if err != nil {
		if err == utils.ErrNotFound {
			returnError(w, err.Error(), http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (c *Controller) getByParameters(w http.ResponseWriter, r *http.Request, model string) {
	// логика на случай если параметр surname // field
	var param string
	if model == "shop" {
		param = "name"
	} else if model == "customer" {
		param = "surname"
	}
	srch, ok := r.URL.Query()[param]
	searchParam := ""
	if !ok || len(srch[0]) < 1 {
		log.Println("Url Param is missing")
		http.Error(w, errors.New("Url Param is missing").Error(), http.StatusBadRequest)
	} else {
		searchParam = srch[0]
		log.Println("Controller GetByParameters with search option: ", searchParam)
	}
	key, ok := r.URL.Query()["field"]
	var field string
	if ok {
		field = key[0]
		log.Println(model+" Controller GetByParameters with field option: ", field)
	} else {
		log.Println(model + " Controller GetByParameters")
		field = ""
	}
	msg, err := c.service.GetByParameters(searchParam, field, model)
	if err != nil {
		if err == utils.ErrNotFound {
			returnError(w, err.Error(), http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(msg)
}

func (c *Controller) getByIDHandler(w http.ResponseWriter, r *http.Request, model string) {
	path := strings.Trim(r.URL.Path, "/")
	pathParts := strings.Split(path, "/")
	id := pathParts[3]
	key, ok := r.URL.Query()["field"]
	var field string
	if ok {
		field = key[0]
		log.Println(model+" Controller GetByID with field option: ", field)
	} else {
		log.Println(model + " Controller GetByID")
		field = ""
	}
	msg, err := c.service.GetById(id, field, model)
	if err != nil {
		if err == utils.ErrNotFound {
			returnError(w, err.Error(), http.StatusNotFound)
			return
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(msg)
}

func returnError(w http.ResponseWriter, err string, status int) {
	log.Println("Error occured ", err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err == "" {
		fmt.Fprintf(w, `{"error": "%s"}`, http.StatusText(status))
		return
	}
	fmt.Fprintf(w, `{"error": "%s: %s"}`, http.StatusText(status), err)
}

func (c *Controller) testEndpoint(w http.ResponseWriter, r *http.Request) {
	/*msg := model.Shop{
		Name: "asd",
	}
	m, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json")
	w.Write(m)*/
	var msg json.RawMessage
	_ = json.NewDecoder(r.Body).Decode(&msg)
	var msg1 model.Customer
	json.Unmarshal(msg, msg1)
	fmt.Println(msg1)
}

//todo: delete by not only id?
//todo: return id
