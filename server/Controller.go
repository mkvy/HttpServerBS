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

// todo: переделать на одинаковые хендлеры, передавать лишь параметром shop/customer тип! повторяешься
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
	if r.URL.Path == "/api/v1/shop/" {
		if r.Method == http.MethodPost {
			c.createShopHandler(w, r)
			return
		}
		if r.Method == http.MethodGet {
			// логика на случай если параметры name = или name = или field =
			keys, ok := r.URL.Query()["name"]
			if !ok || len(keys[0]) < 1 {
				log.Println("Url Param 'name' is missing")
			}
			log.Println(keys)
			c.getAllShopHandler(w, r)
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
			c.patchShopHandler(w, r)
			return
		}
		if r.Method == http.MethodDelete {
			c.deleteShopHandler(w, r)
			return
		}
		if r.Method == http.MethodGet {
			c.getShopByIDHandler(w, r)
		}
		http.Error(w, fmt.Sprintf("expect method DELETE or PATCH or GET at /api/v1/shop/<id>, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

}

func (c *Controller) createShopHandler(w http.ResponseWriter, r *http.Request) {
	var a map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(a)
	return
}

func (c *Controller) patchShopHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	pathParts := strings.Split(path, "/")
	id := pathParts[3]
	fmt.Println(id)
	var a map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("PATCHED")
	return
}

func (c *Controller) deleteShopHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	pathParts := strings.Split(path, "/")
	id := pathParts[3]
	fmt.Println(id)
}

func (c *Controller) getAllShopHandler(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) getShopByIDHandler(w http.ResponseWriter, r *http.Request) {
	key, ok := r.URL.Query()["field"]
	if ok {
		fmt.Println("Have a key field: ", key)
	}
}

// POST /api/v1/customer/
// PATCH /api/v1/customer/{id}
// DELETE /api/v1/customer/{id}
// GET /api/v1/customer/?surname={surname}
// GET /api/v1/customer?surname={surname}&field={field}
func (c *Controller) CustController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// если без <id> запрос
	if r.URL.Path == "/api/v1/customer/" {
		if r.Method == http.MethodPost {
			c.createShopHandler(w, r)
			return
		}
		if r.Method == http.MethodGet {
			// логика на случай если параметры surname = или surname = или field =
			keys, ok := r.URL.Query()["surname"]
			if !ok || len(keys[0]) < 1 {
				log.Println("Url Param 'name' is missing")
			}
			log.Println(keys)
			c.getAllShopHandler(w, r)
			return
		}
		http.Error(w, fmt.Sprintf("expect method GET or POST at /api/v1/shop/, got %v", r.Method), http.StatusMethodNotAllowed)
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
			c.patchShopHandler(w, r)
			return
		}
		if r.Method == http.MethodDelete {
			c.deleteShopHandler(w, r)
			return
		}
		if r.Method == http.MethodGet {
			c.getShopByIDHandler(w, r)
		}
		http.Error(w, fmt.Sprintf("expect method DELETE or PATCH or GET at /api/v1/customer/<id>, got %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

}

func (c *Controller) createCustHandler(w http.ResponseWriter, r *http.Request) {
	var a map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(a)
	return
}

func (c *Controller) patchCustHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	pathParts := strings.Split(path, "/")
	id := pathParts[3]
	fmt.Println(id)
	var a map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("PATCHED")
	return
}

func (c *Controller) deleteCustHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	pathParts := strings.Split(path, "/")
	id := pathParts[3]
	fmt.Println(id)
}

func (c *Controller) getAllCustHandler(w http.ResponseWriter, r *http.Request) {

}
