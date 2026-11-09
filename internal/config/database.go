package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnectionDB(dsn string) (*gorm.DB, error) {

	//Open db connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	log.Println("Connected to database")

	return db, err

}
