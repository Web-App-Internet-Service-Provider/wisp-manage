package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// PlanRepository -> Plan Repository CRUD
type PlanRepository interface {
	GetPlan(int) (models.Plan, error)
	GetAllPlan() ([]models.Plan, error)
	AddPlan(models.Plan) (models.Plan, error)
	UpdatePlan(models.Plan) (models.Plan, error)
	DeletePlan(models.Plan) (models.Plan, error)
}
type planRepository struct {
	connection *gorm.DB
}

// NewPlanRepository -> returns new billing statement repositoty
func NewPlanRepository() PlanRepository {
	return &planRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *planRepository) GetPlan(id int) (plan models.Plan, err error) {
	return plan, db.connection.First(&plan, id).Error
}

func (db *planRepository) GetAllPlan() (plan []models.Plan, err error) {
	return plan, db.connection.Find(&plan).Error
}

func (db *planRepository) AddPlan(plan models.Plan) (models.Plan, error) {
	return plan, db.connection.Create(&plan).Error
}

func (db *planRepository) UpdatePlan(plan models.Plan) (models.Plan, error) {
	if err := db.connection.First(&plan, plan.ID).Error; err != nil {
		return plan, err
	}
	return plan, db.connection.Model(&plan).Updates(&plan).Error
}

func (db *planRepository) DeletePlan(plan models.Plan) (models.Plan, error) {
	if err := db.connection.First(&plan, plan.ID).Error; err != nil {
		return plan, err
	}
	return plan, db.connection.Delete(&plan).Error
}
