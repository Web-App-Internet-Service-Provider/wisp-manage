package models

import "gorm.io/gorm"

// Station is used to map your stations database table to your go code.
type Station struct {
	gorm.Model
	Name           string `gorm:"column:station_name" json:"station_name"`
	IPAddress      string `gorm:"column:station_ipaddress" json:"station_ipaddress"`
	NASPassword    string `gorm:"column:station_nas_password" json:"station_nas_password"`
	NasApiID       int    `gorm:"column:station_nas_api_id" json:"station_nas_api_id"`
	LocationID     uint64 `gorm:"column:station_location_id" json:"station_location_id"`
	OrganizationID uint64 `gorm:"column:station_organization_id" json:"station_organization_id"`
	PlanID         uint64 `gorm:"column:station_plan_id" json:"station_plan_id"`
	CustomerID     uint64 `gorm:"column:station_customer_id" json:"station_customer_id"`
}
