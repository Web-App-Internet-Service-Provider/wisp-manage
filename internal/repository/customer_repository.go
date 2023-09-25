package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// CustomerRepository -> Customer CRUD
type CustomerRepository interface {
	AddCustomer(models.Customer) (models.Customer, error)
	GetCustomer(int) (models.Customer, error)
	GetByPhoneNumber(string) (models.Customer, error)
	GetAllCustomer() ([]models.Customer, error)
	UpdateCustomer(models.Customer) (models.Customer, error)
	DeleteCustomer(models.Customer) (models.Customer, error)
	GetCustomerPlanType(int) ([]models.PlanType, error)
}

type customerRepository struct {
	connection *gorm.DB
}

// NewCustomerRepository --> returns new Customer repository
func NewCustomerRepository() CustomerRepository {
	return &customerRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *customerRepository) GetCustomer(id int) (Customer models.Customer, err error) {
	return Customer, db.connection.First(&Customer, id).Error
}

func (db *customerRepository) GetByPhoneNumber(email string) (Customer models.Customer, err error) {
	return Customer, db.connection.First(&Customer, "phone_number=?", Customer.PhoneNumber).Error
}

func (db *customerRepository) GetAllCustomer() (Customers []models.Customer, err error) {
	return Customers, db.connection.Find(&Customers).Error
}

func (db *customerRepository) AddCustomer(Customer models.Customer) (models.Customer, error) {
	return Customer, db.connection.Create(&Customer).Error
}

func (db *customerRepository) UpdateCustomer(Customer models.Customer) (models.Customer, error) {
	if err := db.connection.First(&Customer, Customer.ID).Error; err != nil {
		return Customer, err
	}
	return Customer, db.connection.Model(&Customer).Updates(&Customer).Error
}

func (db *customerRepository) DeleteCustomer(Customer models.Customer) (models.Customer, error) {
	if err := db.connection.First(&Customer, Customer.ID).Error; err != nil {
		return Customer, err
	}
	return Customer, db.connection.Delete(&Customer).Error
}

func (db *customerRepository) GetCustomerPlanType(CustomerID int) (plantype []models.PlanType, err error) {
	return plantype, db.connection.Where("Customer_id = ?", CustomerID).Set("gorm:auto_preload", true).Find(&plantype).Error
}
