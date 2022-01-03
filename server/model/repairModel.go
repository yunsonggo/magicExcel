package model

import (
	"gorm.io/gorm"
	"time"
)

type RepairModel struct {
	ID         int64          `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoCreateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Class      string         `gorm:"index" json:"class"`
	CarName    string         `gorm:"index" json:"car_name"`
	Pay        float64        `json:"pay"`
	Status     string         `json:"status"`
	DateString string         `gorm:"index" json:"date_string"`
}

type RepairDataModel struct {
	RepairDateString string  `json:"repair_date_string"`
	RepairPay        float64 `json:"repair_pay"`
	RepairStatus     string  `json:"repair_status"`
}
