package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// OrganizationSettingRepository -> Organization Setting CRUD
type OrganizationSettingRepository interface {
	GetOrganizationSetting(int) (models.OrganizationSetting, error)
	GetAllOrganizationSetting() ([]models.OrganizationSetting, error)
	AddOrganizationSetting(models.OrganizationSetting) (models.OrganizationSetting, error)
	UpdateOrganizationSetting(models.OrganizationSetting) (models.OrganizationSetting, error)
	DeleteOrganizationSetting(models.OrganizationSetting) (models.OrganizationSetting, error)
}
type organizationSettingRepository struct {
	connection *gorm.DB
}

// NewOrganizationSetting -> returns new Message template repositoty
func NewOrganizationSettingRepository() OrganizationSettingRepository {
	return &organizationSettingRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *organizationSettingRepository) GetOrganizationSetting(id int) (organizationSetting models.OrganizationSetting, err error) {
	return organizationSetting, db.connection.First(&organizationSetting, id).Error
}

func (db *organizationSettingRepository) GetAllOrganizationSetting() (organizationSetting []models.OrganizationSetting, err error) {
	return organizationSetting, db.connection.Find(&organizationSetting).Error
}

func (db *organizationSettingRepository) AddOrganizationSetting(organizationSetting models.OrganizationSetting) (models.OrganizationSetting, error) {
	return organizationSetting, db.connection.Create(&organizationSetting).Error
}

func (db *organizationSettingRepository) UpdateOrganizationSetting(organizationSetting models.OrganizationSetting) (models.OrganizationSetting, error) {
	if err := db.connection.First(&organizationSetting, organizationSetting.ID).Error; err != nil {
		return organizationSetting, err
	}
	return organizationSetting, db.connection.Model(&organizationSetting).Updates(&organizationSetting).Error
}

func (db *organizationSettingRepository) DeleteOrganizationSetting(organizationSetting models.OrganizationSetting) (models.OrganizationSetting, error) {
	if err := db.connection.First(&organizationSetting, organizationSetting.ID).Error; err != nil {
		return organizationSetting, err
	}
	return organizationSetting, db.connection.Delete(&organizationSetting).Error
}
