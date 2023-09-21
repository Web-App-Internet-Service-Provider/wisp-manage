package models

import (
	"gorm.io/gorm"
)
// ExtensionLog is used to map your billing_statements database table to your go code.
type ExtensionLog struct {
	gorm.Model
	UserID    uint64 `gorm:"column:user_id" json:"user_id"`
	ServiceID uint64 `gorm:"column:service_id" json:"service_id"`
	Duration  int    `gorm:"column:duration" json:"duration"`
}
