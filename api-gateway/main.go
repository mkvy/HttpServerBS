package main

import (
	"github.com/mkvy/HttpServerBS/api-gateway/client"
	"github.com/mkvy/HttpServerBS/api-gateway/internal/config"
	"github.com/mkvy/HttpServerBS/api-gateway/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	cfg := config.NewConfigFromFile()
	svc := client.NewGrpcClient(cfg)
	controller := server.NewController(svc)
	s := server.NewServer(cfg, *controller)
	go s.Start()
	log.Println("Http Server is running")
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, os.Interrupt, os.Kill)
	<-sigTerm
	s.Stop()
}
