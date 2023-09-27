package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// PlanStationRepository -> Plan Station CRUD
type PlanStationRepository interface {
	GetPlanStation(int) (models.PlanStation, error)
	GetAllPlanStation() ([]models.PlanStation, error)
	AddPlanStation(models.PlanStation) (models.PlanStation, error)
	UpdatePlanStation(models.PlanStation) (models.PlanStation, error)
	DeletePlanStation(models.PlanStation) (models.PlanStation, error)
}
type planStationRepository struct {
	connection *gorm.DB
}

// NewplanStationRepository -> returns new plan Station repositoty
func NewPlanStationRepository() PlanStationRepository {
	return &planStationRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *planStationRepository) GetPlanStation(id int) (planStation models.PlanStation, err error) {
	return planStation, db.connection.First(&planStation, id).Error
}

func (db *planStationRepository) GetAllPlanStation() (planStation []models.PlanStation, err error) {
	return planStation, db.connection.Find(&planStation).Error
}

func (db *planStationRepository) AddPlanStation(planStation models.PlanStation) (models.PlanStation, error) {
	return planStation, db.connection.Create(&planStation).Error
}

func (db *planStationRepository) UpdatePlanStation(planStation models.PlanStation) (models.PlanStation, error) {
	if err := db.connection.First(&planStation, planStation.ID).Error; err != nil {
		return planStation, err
	}
	return planStation, db.connection.Model(&planStation).Updates(&planStation).Error
}

func (db *planStationRepository) DeletePlanStation(planStation models.PlanStation) (models.PlanStation, error) {
	if err := db.connection.First(&planStation, planStation.ID).Error; err != nil {
		return planStation, err
	}
	return planStation, db.connection.Delete(&planStation).Error
}
