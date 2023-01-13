package server

import (
	"context"
	"github.com/mkvy/HttpServerBS/api-gateway/internal/config"
	"log"
	"net/http"
)

type Server struct {
	srv  *http.Server
	addr string
}

func NewServer(cfg config.Config, c Controller) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/shop/", c.ShopController)
	mux.HandleFunc("/api/v1/customer/", c.CustController)
	server := &http.Server{Addr: ":" + cfg.HttpServer.Port, Handler: mux}
	return &Server{server, cfg.HttpServer.Port}
}

func (s *Server) Start() {
	if err := s.srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func (s *Server) Stop() {
	s.srv.Shutdown(context.Background())
}
