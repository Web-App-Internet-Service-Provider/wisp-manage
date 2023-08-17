package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/AbdulrahmanDaud10/wisp-manage/pkg/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	err error
)

// SetUpDatabaseConnection opens a database and saves the reference to `Database` struct.
func SetUpDatabaseConnection() (*gorm.DB, error) {
	var db *gorm.DB

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		username = os.Getenv("DB_USER")
		database = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)
	dsn := fmt.Sprintf("host=%s port=%s user=%s database=%s  password=%s sslmode=disable",
		host,
		port,
		username,
		database,
		password,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: ", err)
		return nil, err
	}
	// Run the customer migration after implemented
	db.AutoMigrate(api.Customer{
		Model: api.Model{},
	})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}

	fmt.Println("Database connection successful...")

	return db, nil
}
