package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// paymentDetailRepository -> Message log Repository CRUD
type PaymentDetailRepository interface {
	GetPaymenDetail(int) (models.PaymentDetail, error)
	GetAllPaymenDetail() ([]models.PaymentDetail, error)
	AddPaymenDetail(models.PaymentDetail) (models.PaymentDetail, error)
	UpdatePaymenDetail(models.PaymentDetail) (models.PaymentDetail, error)
	DeletePaymenDetail(models.PaymentDetail) (models.PaymentDetail, error)
}
type paymentDetailRepository struct {
	connection *gorm.DB
}

// NewPaymentDetailRepository -> returns new Message log repositoty
func NewPaymentDetailRepository() PaymentDetailRepository {
	return &paymentDetailRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *paymentDetailRepository) GetPaymenDetail(id int) (paymentDetail models.PaymentDetail, err error) {
	return paymentDetail, db.connection.First(&paymentDetail, id).Error
}

func (db *paymentDetailRepository) GetAllPaymenDetail() (paymentDetail []models.PaymentDetail, err error) {
	return paymentDetail, db.connection.Find(&paymentDetail).Error
}

func (db *paymentDetailRepository) AddPaymenDetail(paymentDetail models.PaymentDetail) (models.PaymentDetail, error) {
	return paymentDetail, db.connection.Create(&paymentDetail).Error
}

func (db *paymentDetailRepository) UpdatePaymenDetail(paymentDetail models.PaymentDetail) (models.PaymentDetail, error) {
	if err := db.connection.First(&paymentDetail, paymentDetail.ID).Error; err != nil {
		return paymentDetail, err
	}
	return paymentDetail, db.connection.Model(&paymentDetail).Updates(&paymentDetail).Error
}

func (db *paymentDetailRepository) DeletePaymenDetail(paymentDetail models.PaymentDetail) (models.PaymentDetail, error) {
	if err := db.connection.First(&paymentDetail, paymentDetail.ID).Error; err != nil {
		return paymentDetail, err
	}
	return paymentDetail, db.connection.Delete(&paymentDetail).Error
}
