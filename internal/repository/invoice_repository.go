package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// InvoiceRepository -> Invoice Repository CRUD
type InvoiceRepository interface {
	GetInvoice(int) (models.Invoice, error)
	GetAllInvoice() ([]models.Invoice, error)
	AddInvoice(models.Invoice) (models.Invoice, error)
	UpdateInvoice(models.Invoice) (models.Invoice, error)
	DeleteInvoice(models.Invoice) (models.Invoice, error)
}
type invoiceRepository struct {
	connection *gorm.DB
}

// NewInvoiceRepository -> returns new billing statement repositoty
func NewInvoiceRepository() InvoiceRepository {
	return &invoiceRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *invoiceRepository) GetInvoice(id int) (invoice models.Invoice, err error) {
	return invoice, db.connection.First(&invoice, id).Error
}

func (db *invoiceRepository) GetAllInvoice() (invoice []models.Invoice, err error) {
	return invoice, db.connection.Find(&invoice).Error
}

func (db *invoiceRepository) AddInvoice(invoice models.Invoice) (models.Invoice, error) {
	return invoice, db.connection.Create(&invoice).Error
}

func (db *invoiceRepository) UpdateInvoice(invoice models.Invoice) (models.Invoice, error) {
	if err := db.connection.First(&invoice, invoice.ID).Error; err != nil {
		return invoice, err
	}
	return invoice, db.connection.Model(&invoice).Updates(&invoice).Error
}

func (db *invoiceRepository) DeleteInvoice(invoice models.Invoice) (models.Invoice, error) {
	if err := db.connection.First(&invoice, invoice.ID).Error; err != nil {
		return invoice, err
	}
	return invoice, db.connection.Delete(&invoice).Error
}
