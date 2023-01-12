package main

import (
	_ "github.com/lib/pq"
	"github.com/mkvy/HttpServerBS/internal/config"
	"github.com/mkvy/HttpServerBS/internal/utils"
	"github.com/mkvy/HttpServerBS/server"
	"github.com/mkvy/HttpServerBS/service"
	"log"
	"os"
	"os/signal"
)

func main() {
	cfg := config.NewConfigFromFile()
	contr := server.NewController(service.NewServiceImpl(cfg))
	s := server.NewServer("8282", *contr)
	go s.Start()
	log.Println("Server is running")
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, os.Interrupt, os.Kill)
	<-sigTerm
	utils.DBClose()
	s.Stop()
}
