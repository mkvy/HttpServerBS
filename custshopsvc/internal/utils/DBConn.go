package utils

import (
	"database/sql"
	"fmt"
	"github.com/mkvy/HttpServerBS/custshopsvc/internal/config"
	"log"
	"sync"
)

var db *sql.DB

func GetDBConn(cfg config.Config) (*sql.DB, error) {
	once := sync.Once{}
	var err error
	once.Do(func() {
		log.Println("Creating DB connection in DB Shop repository with driver " + cfg.Database.DriverName)
		connStr := "user=" + cfg.Database.Username + " password=" + cfg.Database.Password + " dbname=" + cfg.Database.DBname + " sslmode=disable"
		fmt.Println(connStr)
		db, err = sql.Open(cfg.Database.DriverName, connStr)
		if err != nil {
			log.Println("Error while connecting to db")
			log.Println(err)
			return
		}
		if db == nil {
			log.Println("Error with database")
			return
		}
	})
	return db, err
}

func DBClose() {
	err := db.Close()
	if err != nil {
		log.Println("error closing connection with database")
	}
	return
}
