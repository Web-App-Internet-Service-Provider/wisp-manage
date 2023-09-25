package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// MessagetemplateRepository -> Message template Repository CRUD
type MessageTemplateRepository interface {
	GetMessageTemplate(int) (models.MessageTemplate, error)
	GetAllMessageTemplate() ([]models.MessageTemplate, error)
	AddMessageTemplate(models.MessageTemplate) (models.MessageTemplate, error)
	UpdateMessageTemplate(models.MessageTemplate) (models.MessageTemplate, error)
	DeleteMessageTemplate(models.MessageTemplate) (models.MessageTemplate, error)
}
type messageTemplateRepository struct {
	connection *gorm.DB
}

// NewMessageTemplateRepository -> returns new Message template repositoty
func NewMessagetemplateRepository() MessageTemplateRepository {
	return &messageTemplateRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *messageTemplateRepository) GetMessageTemplate(id int) (messageTemplate models.MessageTemplate, err error) {
	return messageTemplate, db.connection.First(&messageTemplate, id).Error
}

func (db *messageTemplateRepository) GetAllMessageTemplate() (messageTemplate []models.MessageTemplate, err error) {
	return messageTemplate, db.connection.Find(&messageTemplate).Error
}

func (db *messageTemplateRepository) AddMessageTemplate(messageTemplate models.MessageTemplate) (models.MessageTemplate, error) {
	return messageTemplate, db.connection.Create(&messageTemplate).Error
}

func (db *messageTemplateRepository) UpdateMessageTemplate(messageTemplate models.MessageTemplate) (models.MessageTemplate, error) {
	if err := db.connection.First(&messageTemplate, messageTemplate.ID).Error; err != nil {
		return messageTemplate, err
	}
	return messageTemplate, db.connection.Model(&messageTemplate).Updates(&messageTemplate).Error
}

func (db *messageTemplateRepository) DeleteMessageTemplate(messageTemplate models.MessageTemplate) (models.MessageTemplate, error) {
	if err := db.connection.First(&messageTemplate, messageTemplate.ID).Error; err != nil {
		return messageTemplate, err
	}
	return messageTemplate, db.connection.Delete(&messageTemplate).Error
}
