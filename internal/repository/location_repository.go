package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// LocationRepository -> Location Repository CRUD
type LocationRepository interface {
	GetLocation(int) (models.Location, error)
	GetAllLocation() ([]models.Location, error)
	AddLocation(models.Location) (models.Location, error)
	UpdateLocation(models.Location) (models.Location, error)
	DeleteLocation(models.Location) (models.Location, error)
}
type locationRepository struct {
	connection *gorm.DB
}

// NewLocationRepository -> returns new billing statement repositoty
func NewLocationRepository() LocationRepository {
	return &locationRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *locationRepository) GetLocation(id int) (location models.Location, err error) {
	return location, db.connection.First(&location, id).Error
}

func (db *locationRepository) GetAllLocation() (location []models.Location, err error) {
	return location, db.connection.Find(&location).Error
}

func (db *locationRepository) AddLocation(location models.Location) (models.Location, error) {
	return location, db.connection.Create(&location).Error
}

func (db *locationRepository) UpdateLocation(location models.Location) (models.Location, error) {
	if err := db.connection.First(&location, location.ID).Error; err != nil {
		return location, err
	}
	return location, db.connection.Model(&location).Updates(&location).Error
}

func (db *locationRepository) DeleteLocation(location models.Location) (models.Location, error) {
	if err := db.connection.First(&location, location.ID).Error; err != nil {
		return location, err
	}
	return location, db.connection.Delete(&location).Error
}
