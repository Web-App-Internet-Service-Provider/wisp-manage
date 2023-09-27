package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// MessagelogRepository -> Message log Repository CRUD
type MessagelogRepository interface {
	GetMessageLog(int) (models.MessageLog, error)
	GetAllMessageLog() ([]models.MessageLog, error)
	AddMessageLog(models.MessageLog) (models.MessageLog, error)
	UpdateMessageLog(models.MessageLog) (models.MessageLog, error)
	DeleteMessageLog(models.MessageLog) (models.MessageLog, error)
}
type messagelogRepository struct {
	connection *gorm.DB
}

// NewMessagelogRepository -> returns new Message log repositoty
func NewMessagelogRepository() MessagelogRepository {
	return &messagelogRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *messagelogRepository) GetMessageLog(id int) (messageLog models.MessageLog, err error) {
	return messageLog, db.connection.First(&messageLog, id).Error
}

func (db *messagelogRepository) GetAllMessageLog() (messageLog []models.MessageLog, err error) {
	return messageLog, db.connection.Find(&messageLog).Error
}

func (db *messagelogRepository) AddMessageLog(messageLog models.MessageLog) (models.MessageLog, error) {
	return messageLog, db.connection.Create(&messageLog).Error
}

func (db *messagelogRepository) UpdateMessageLog(messageLog models.MessageLog) (models.MessageLog, error) {
	if err := db.connection.First(&messageLog, messageLog.ID).Error; err != nil {
		return messageLog, err
	}
	return messageLog, db.connection.Model(&messageLog).Updates(&messageLog).Error
}

func (db *messagelogRepository) DeleteMessageLog(messageLog models.MessageLog) (models.MessageLog, error) {
	if err := db.connection.First(&messageLog, messageLog.ID).Error; err != nil {
		return messageLog, err
	}
	return messageLog, db.connection.Delete(&messageLog).Error
}
