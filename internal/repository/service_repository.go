package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// ServiceRepository -> service Repository CRUD
type ServiceRepository interface {
	GetService(int) (models.Service, error)
	GetAllService() ([]models.Service, error)
	AddService(models.Service) (models.Service, error)
	UpdateService(models.Service) (models.Service, error)
	DeleteService(models.Service) (models.Service, error)
}
type serviceRepository struct {
	connection *gorm.DB
}

// NewServiceRepository -> returns new service repositoty
func NewServiceRepository() ServiceRepository {
	return &serviceRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *serviceRepository) GetService(id int) (service models.Service, err error) {
	return service, db.connection.First(&service, id).Error
}

func (db *serviceRepository) GetAllService() (service []models.Service, err error) {
	return service, db.connection.Find(&service).Error
}

func (db *serviceRepository) AddService(service models.Service) (models.Service, error) {
	return service, db.connection.Create(&service).Error
}

func (db *serviceRepository) UpdateService(service models.Service) (models.Service, error) {
	if err := db.connection.First(&service, service.ID).Error; err != nil {
		return service, err
	}
	return service, db.connection.Model(&service).Updates(&service).Error
}

func (db *serviceRepository) DeleteService(service models.Service) (models.Service, error) {
	if err := db.connection.First(&service, service.ID).Error; err != nil {
		return service, err
	}
	return service, db.connection.Delete(&service).Error
}
