package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// ExtensionLogRepository -> Extension log CRUD
type ExtensionLogRepository interface {
	GetExtensionLog(int) (models.ExtensionLog, error)
	GetAllExtensionLog() ([]models.ExtensionLog, error)
	AddExtensionLog(models.ExtensionLog) (models.ExtensionLog, error)
	UpdateExtensionLog(models.ExtensionLog) (models.ExtensionLog, error)
	DeleteExtensionLog(models.ExtensionLog) (models.ExtensionLog, error)
}
type extensionLogRepository struct {
	connection *gorm.DB
}

// NewExtensionLogRepository -> returns new billing statement repositoty
func NewExtensionLogRepository() ExtensionLogRepository {
	return &extensionLogRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *extensionLogRepository) GetExtensionLog(id int) (extensionLog models.ExtensionLog, err error) {
	return extensionLog, db.connection.First(&extensionLog, id).Error
}

func (db *extensionLogRepository) GetAllExtensionLog() (extensionLog []models.ExtensionLog, err error) {
	return extensionLog, db.connection.Find(&extensionLog).Error
}

func (db *extensionLogRepository) AddExtensionLog(extensionLog models.ExtensionLog) (models.ExtensionLog, error) {
	return extensionLog, db.connection.Create(&extensionLog).Error
}

func (db *extensionLogRepository) UpdateExtensionLog(extensionLog models.ExtensionLog) (models.ExtensionLog, error) {
	if err := db.connection.First(&extensionLog, extensionLog.ID).Error; err != nil {
		return extensionLog, err
	}
	return extensionLog, db.connection.Model(&extensionLog).Updates(&extensionLog).Error
}

func (db *extensionLogRepository) DeleteExtensionLog(extensionLog models.ExtensionLog) (models.ExtensionLog, error) {
	if err := db.connection.First(&extensionLog, extensionLog.ID).Error; err != nil {
		return extensionLog, err
	}
	return extensionLog, db.connection.Delete(&extensionLog).Error
}
