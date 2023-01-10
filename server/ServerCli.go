package server

import (
	"context"
	"log"
	"net/http"
)

type Server struct {
	srv  *http.Server
	addr string
}

func NewServer(port string, c Controller) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/shop/", c.ShopController)
	mux.HandleFunc("/api/v1/customer/", c.CustController)
	server := &http.Server{Addr: ":" + port, Handler: mux}
	return &Server{server, port}
}

func (s *Server) Start() {
	if err := s.srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func (s *Server) Stop() {
	s.srv.Shutdown(context.Background())
}
