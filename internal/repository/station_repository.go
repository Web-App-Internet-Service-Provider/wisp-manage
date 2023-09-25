package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// StationRepository -> station Repository CRUD
type StationRepository interface {
	GetStation(int) (models.Station, error)
	GetAllStation() ([]models.Station, error)
	AddStation(models.Station) (models.Station, error)
	UpdateStation(models.Station) (models.Station, error)
	DeleteStation(models.Station) (models.Station, error)
}
type stationRepository struct {
	connection *gorm.DB
}

// NewstationRepository -> returns new station repositoty
func NewStationRepository() StationRepository {
	return &stationRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *stationRepository) GetStation(id int) (station models.Station, err error) {
	return station, db.connection.First(&station, id).Error
}

func (db *stationRepository) GetAllStation() (station []models.Station, err error) {
	return station, db.connection.Find(&station).Error
}

func (db *stationRepository) AddStation(station models.Station) (models.Station, error) {
	return station, db.connection.Create(&station).Error
}

func (db *stationRepository) UpdateStation(station models.Station) (models.Station, error) {
	if err := db.connection.First(&station, station.ID).Error; err != nil {
		return station, err
	}
	return station, db.connection.Model(&station).Updates(&station).Error
}

func (db *stationRepository) DeleteStation(station models.Station) (models.Station, error) {
	if err := db.connection.First(&station, station.ID).Error; err != nil {
		return station, err
	}
	return station, db.connection.Delete(&station).Error
}
