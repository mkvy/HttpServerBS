package main

import (
	_ "github.com/lib/pq"
	"github.com/mkvy/HttpServerBS/custshopsvc/internal/config"
	"github.com/mkvy/HttpServerBS/custshopsvc/internal/utils"
	"github.com/mkvy/HttpServerBS/custshopsvc/server"
	"github.com/mkvy/HttpServerBS/custshopsvc/service"
	"log"
	"os"
	"os/signal"
)

func main() {
	cfg := config.NewConfigFromFile()
	svc := service.NewServiceImpl(cfg)
	controller := server.NewGrpcController(svc)
	gServer := server.NewGrpcServer(cfg, *controller)
	go gServer.Start()
	log.Println("Server is running")
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, os.Interrupt, os.Kill)
	<-sigTerm
	utils.DBClose()
	gServer.Stop()
}
