package models

import "gorm.io/gorm"

// Location is used to map your locations database table to your go code.
type Location struct {
	gorm.Model
	Name           string `gorm:"colum:location_name" json:"location_name"`
	OrganizationID uint64 `gorm:"colum:organization_id" json:"organization_id"`
}
