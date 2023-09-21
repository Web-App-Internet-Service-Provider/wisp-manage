package models

import "gorm.io/gorm"

// OrganizationSetting is used to map your locations database table to your go code.
type OrganizationSetting struct {
	gorm.Model
	Key            string `gorm:"column:organizaion_key" json:"organizaion_key"`
	Value          string `gorm:"column:organizaion_value" json:"organizaion_value"`
	OrganizationID uint64 `gorm:"column:organization_id" json:"organization_id"`
}
