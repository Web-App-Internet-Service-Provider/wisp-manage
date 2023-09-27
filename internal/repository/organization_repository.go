package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// OrganizationRepository -> Organization Setting CRUD
type OrganizationRepository interface {
	GetOrganization(int) (models.Organization, error)
	GetAllOrganization() ([]models.Organization, error)
	AddOrganization(models.Organization) (models.Organization, error)
	UpdateOrganization(models.Organization) (models.Organization, error)
	DeleteOrganization(models.Organization) (models.Organization, error)
}
type organizationRepository struct {
	connection *gorm.DB
}

// NewOrganizationSetting -> returns new Message template repositoty
func NewOrganizationRepository() OrganizationRepository {
	return &organizationRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *organizationRepository) GetOrganization(id int) (organization models.Organization, err error) {
	return organization, db.connection.First(&organization, id).Error
}

func (db *organizationRepository) GetAllOrganization() (organization []models.Organization, err error) {
	return organization, db.connection.Find(&organization).Error
}

func (db *organizationRepository) AddOrganization(organization models.Organization) (models.Organization, error) {
	return organization, db.connection.Create(&organization).Error
}

func (db *organizationRepository) UpdateOrganization(organization models.Organization) (models.Organization, error) {
	if err := db.connection.First(&organization, organization.ID).Error; err != nil {
		return organization, err
	}
	return organization, db.connection.Model(&organization).Updates(&organization).Error
}

func (db *organizationRepository) DeleteOrganization(organization models.Organization) (models.Organization, error) {
	if err := db.connection.First(&organization, organization.ID).Error; err != nil {
		return organization, err
	}
	return organization, db.connection.Delete(&organization).Error
}
