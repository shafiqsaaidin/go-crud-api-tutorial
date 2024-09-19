package initializers

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("crudapi.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}
