package models

// Plan is used to map your plans database table to your go code.
type Plan struct {
	Name           string `gorm:"column:plan_name" json:"plan_name"`
	PlanTypeID     uint64 `gorm:"column:plan_type_id" json:"plan_type_id"`
	Price          int    `gorm:"column:plan_price" json:"plan_price"`
	ExpirationType string `gorm:"column:plan_expirationtype" json:"plan_expirationtype"`
	Speed          string `gorm:"column:plan_speed" json:"plan_speed"`
	PoolID         uint64 `gorm:"column:plan_pool_id" json:"plan_pool_id"`
	Device         string `gorm:"column:plan_device" json:"plan_device" `
	OrganizationID uint64 `gorm:"column:plan_organization_id" json:"plan_organization_id"`
}
