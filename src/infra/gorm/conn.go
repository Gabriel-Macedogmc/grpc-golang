package gorm

import (
	"log"

	"github.com/Gabriel-Macedogmc/grpc-golang/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(psqlConn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(psqlConn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	RunMigrations(db)

	return db
}

func RunMigrations(db *gorm.DB) {

	db.AutoMigrate(&models.Category{})
}
