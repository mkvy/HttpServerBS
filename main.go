package main

import (
	"github.com/mkvy/HttpServerBS/server"
	"github.com/mkvy/HttpServerBS/service"
	"log"
	"os"
	"os/signal"
)

func main() {
	contr := server.NewController(service.Service{})
	s := server.NewServer("8282", *contr)
	go s.Start()
	log.Println("Server is running")
	//graceful shutdown на прерывание
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, os.Interrupt, os.Kill)
	<-sigTerm
	s.Stop()
}