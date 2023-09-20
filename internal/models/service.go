package models

import (
	"time"

	"gorm.io/gorm"
)

// Service is used to map your services database table to your go code.
type Service struct {
	gorm.Model
	CustomerID uint64    `gorm:"column:service_customer_id" json:"service_customer_id"`
	PlanID     uint64    `gorm:"column:service_plan_id" json:"service_plan_id"`
	Expiry     time.Time `gorm:"column:service_expirytime" json:"service_expirytime"`
	Status     string    `gorm:"column:service_status" json:"service_status"`
	Type       string    `gorm:"column:service_type" json:"service_type"`
}
