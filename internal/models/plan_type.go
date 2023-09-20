package models

import "gorm.io/gorm"

// Plantype is used to map your plantypes database table to your go code.
type PlanType struct {
	gorm.Model
	Name   string `gorm:"column:plan_typde_name" json:"plan_typde_name"`
	PlanID uint64 `gorm:"column:plan_id" json:"plan_id"`
}
