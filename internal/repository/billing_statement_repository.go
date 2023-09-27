package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// BillingStatementRepository -> BillingStatement CRUD
type BillingStatementRepository interface {
	GetBillingStatement(int) (models.BillingStatement, error)
	GetAllBillingStatement() ([]models.BillingStatement, error)
	AddBillingStatement(models.BillingStatement) (models.BillingStatement, error)
	UpdateBillingStatement(models.BillingStatement) (models.BillingStatement, error)
	DeleteBillingStatement(models.BillingStatement) (models.BillingStatement, error)
}
type billingStatementRepository struct {
	connection *gorm.DB
}

// NewBillingStatementRepository -> returns new billing statement repositoty
func NewBillingStatementRepository() BillingStatementRepository {
	return &billingStatementRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *billingStatementRepository) GetBillingStatement(id int) (billingStatement models.BillingStatement, err error) {
	return billingStatement, db.connection.First(&billingStatement, id).Error
}

func (db *billingStatementRepository) GetAllBillingStatement() (billingStatements []models.BillingStatement, err error) {
	return billingStatements, db.connection.Find(&billingStatements).Error
}

func (db *billingStatementRepository) AddBillingStatement(billingStatement models.BillingStatement) (models.BillingStatement, error) {
	return billingStatement, db.connection.Create(&billingStatement).Error
}

func (db *billingStatementRepository) UpdateBillingStatement(billingStatement models.BillingStatement) (models.BillingStatement, error) {
	if err := db.connection.First(&billingStatement, billingStatement.ID).Error; err != nil {
		return billingStatement, err
	}
	return billingStatement, db.connection.Model(&billingStatement).Updates(&billingStatement).Error
}

func (db *billingStatementRepository) DeleteBillingStatement(billingStatement models.BillingStatement) (models.BillingStatement, error) {
	if err := db.connection.First(&billingStatement, billingStatement.ID).Error; err != nil {
		return billingStatement, err
	}
	return billingStatement, db.connection.Delete(&billingStatement).Error
}
