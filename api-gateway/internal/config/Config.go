package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	HttpServer struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"http-server"`
	GrpcServer struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"grpc-server"`
}

func NewConfigFromFile() Config {
	f, err := os.Open("internal/config/config.yml")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer f.Close()
	cfg := new(Config)
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return *cfg
}
