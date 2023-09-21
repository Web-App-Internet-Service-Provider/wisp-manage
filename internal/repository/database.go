package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	err error
)

// SetUpDatabaseConnection opens a database, runs migrations and saves the reference to `Database` struct.
func SetUpDatabaseConnection() *gorm.DB {
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
		return nil
	}
	// Run database migrations for the models
	db.AutoMigrate(models.BillingStatement{}, models.Customer{}, models.ExtensionLog{}, models.Invoice{},
		models.Location{}, models.Messagelog{}, models.MessageTemplate{}, models.OrganizationSetting{},
		models.Organization{}, models.PaymentDetail{}, models.PaymentMethod{}, models.PlanStation{},
		models.PlanType{}, models.Plan{}, models.Pool{}, models.Service{}, models.Station{},
		models.User{})
	if err != nil {
		log.Fatal("Error running database migrations...", err)
	}

	fmt.Println("Database connection successful...")

	return db
}
