package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/scientist-v08/bphs/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var SK string

func ConnectToDb() {
    var err error

    dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	SK = os.Getenv("SYMMETRIC_KEY")

    if err != nil {
        log.Fatal("Failed to connect to database")
    }

	// AutoMigrate creates the table if it doesn't exist
	err = DB.AutoMigrate(
		&models.User{}, &models.Session{},
	)
	if err != nil {
    	log.Fatal("Failed to migrate database: ", err)
	}
}