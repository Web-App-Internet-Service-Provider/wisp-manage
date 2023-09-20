package models

import (
	"gorm.io/gorm"
)

// Planstation is used to map your planstations database table to your go code.
type PlanStation struct {
	gorm.Model
	StationID uint64 `gorm:"column:plan_station_id" json:"plan_station_id"`
	PlanID    uint64 `gorm:"column:plan_id" json:"plan_id"`
}
