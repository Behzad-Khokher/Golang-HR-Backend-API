package database

import (
	"GOHR/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	// host := os.Getenv("DB_HOST")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_NAME")
	// port := os.Getenv("DB_PORT")
	// sslmode := os.Getenv("DB_SSLMODE")

	// dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode
	dsn := "host=localhost user=behzad password=123 dbname=golang port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	// Create the tables
	db.AutoMigrate(&models.Channel{}, &models.Post{})

	// Singleton
	Database = DbInstance{
		Db: db,
	}
}
