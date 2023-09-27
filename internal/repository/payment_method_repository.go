package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// paymentMethodRepository -> Message log Repository CRUD
type PaymentMethodRepository interface {
	GetPaymenMethod(int) (models.PaymentMethod, error)
	GetAllPaymenMethod() ([]models.PaymentMethod, error)
	AddPaymenMethod(models.PaymentMethod) (models.PaymentMethod, error)
	UpdatePaymenMethod(models.PaymentMethod) (models.PaymentMethod, error)
	DeletePaymenMethod(models.PaymentMethod) (models.PaymentMethod, error)
}
type paymentMethodRepository struct {
	connection *gorm.DB
}

// NewpaymentMethodRepository -> returns new Message log repositoty
func NewPaymentMethodRepository() PaymentMethodRepository {
	return &paymentMethodRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *paymentMethodRepository) GetPaymenMethod(id int) (paymentMethod models.PaymentMethod, err error) {
	return paymentMethod, db.connection.First(&paymentMethod, id).Error
}

func (db *paymentMethodRepository) GetAllPaymenMethod() (paymentMethod []models.PaymentMethod, err error) {
	return paymentMethod, db.connection.Find(&paymentMethod).Error
}

func (db *paymentMethodRepository) AddPaymenMethod(paymentMethod models.PaymentMethod) (models.PaymentMethod, error) {
	return paymentMethod, db.connection.Create(&paymentMethod).Error
}

func (db *paymentMethodRepository) UpdatePaymenMethod(paymentMethod models.PaymentMethod) (models.PaymentMethod, error) {
	if err := db.connection.First(&paymentMethod, paymentMethod.ID).Error; err != nil {
		return paymentMethod, err
	}
	return paymentMethod, db.connection.Model(&paymentMethod).Updates(&paymentMethod).Error
}

func (db *paymentMethodRepository) DeletePaymenMethod(paymentMethod models.PaymentMethod) (models.PaymentMethod, error) {
	if err := db.connection.First(&paymentMethod, paymentMethod.ID).Error; err != nil {
		return paymentMethod, err
	}
	return paymentMethod, db.connection.Delete(&paymentMethod).Error
}
