package models

import "gorm.io/gorm"

// Pool is used to map your pools database table to your go code.
type Pool struct {
	gorm.Model
	Name           string `gorm:"pool_name" json:"pool_name"`
	OrganizationID uint64 `gorm:"pool_organization_id" json:"pool_organization_id"`
	PlanID         uint64 `gorm:"pool_plan_id" json:"pool_plan_id"`
}
