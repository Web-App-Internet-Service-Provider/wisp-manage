package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// PlanTypeRepository -> Plan Station CRUD
type PlanTypeRepository interface {
	GetPlanType(int) (models.PlanType, error)
	GetAllPlanType() ([]models.PlanType, error)
	AddPlanType(models.PlanType) (models.PlanType, error)
	UpdatePlanType(models.PlanType) (models.PlanType, error)
	DeletePlanType(models.PlanType) (models.PlanType, error)
}
type planTypeRepository struct {
	connection *gorm.DB
}

// NewPlanTypeRepository -> returns new Plan Type repositoty
func NewPlanTypeRepository() PlanTypeRepository {
	return &planTypeRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *planTypeRepository) GetPlanType(id int) (planType models.PlanType, err error) {
	return planType, db.connection.First(&planType, id).Error
}

func (db *planTypeRepository) GetAllPlanType() (planType []models.PlanType, err error) {
	return planType, db.connection.Find(&planType).Error
}

func (db *planTypeRepository) AddPlanType(planType models.PlanType) (models.PlanType, error) {
	return planType, db.connection.Create(&planType).Error
}

func (db *planTypeRepository) UpdatePlanType(planType models.PlanType) (models.PlanType, error) {
	if err := db.connection.First(&planType, planType.ID).Error; err != nil {
		return planType, err
	}
	return planType, db.connection.Model(&planType).Updates(&planType).Error
}

func (db *planTypeRepository) DeletePlanType(planType models.PlanType) (models.PlanType, error) {
	if err := db.connection.First(&planType, planType.ID).Error; err != nil {
		return planType, err
	}
	return planType, db.connection.Delete(&planType).Error
}
