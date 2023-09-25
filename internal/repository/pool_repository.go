package repository

import (
	"github.com/AbdulrahmanDaud10/wisp-manage/internal/models"
	"gorm.io/gorm"
)

// PoolRepository -> pool Repository CRUD
type PoolRepository interface {
	GetPool(int) (models.Pool, error)
	GetAllPool() ([]models.Pool, error)
	AddPool(models.Pool) (models.Pool, error)
	UpdatePool(models.Pool) (models.Pool, error)
	DeletePool(models.Pool) (models.Pool, error)
}
type poolRepository struct {
	connection *gorm.DB
}

// NewPoolRepository -> returns new pool repositoty
func NewPoolRepository() PoolRepository {
	return &poolRepository{
		connection: SetUpDatabaseConnection(),
	}
}

func (db *poolRepository) GetPool(id int) (pool models.Pool, err error) {
	return pool, db.connection.First(&pool, id).Error
}

func (db *poolRepository) GetAllPool() (pool []models.Pool, err error) {
	return pool, db.connection.Find(&pool).Error
}

func (db *poolRepository) AddPool(pool models.Pool) (models.Pool, error) {
	return pool, db.connection.Create(&pool).Error
}

func (db *poolRepository) UpdatePool(pool models.Pool) (models.Pool, error) {
	if err := db.connection.First(&pool, pool.ID).Error; err != nil {
		return pool, err
	}
	return pool, db.connection.Model(&pool).Updates(&pool).Error
}

func (db *poolRepository) DeletePool(pool models.Pool) (models.Pool, error) {
	if err := db.connection.First(&pool, pool.ID).Error; err != nil {
		return pool, err
	}
	return pool, db.connection.Delete(&pool).Error
}
