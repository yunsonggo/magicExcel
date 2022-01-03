package model

import (
	"gorm.io/gorm"
	"time"
)

type OilModel struct {
	ID         int64          `gorm:"primarykey" json:"id"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoCreateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Class      string         `gorm:"index" json:"class"`
	CarName    string         `gorm:"index" json:"car_name"`
	DateString string         `gorm:"index" json:"date_string"`
	BackupNum  string         `json:"backup_num"`
	NowNum     string         `json:"now_num"`
	OilType    string         `json:"oil_type"`
	OilNum     float64        `json:"oil_num"`
	Pay        float64        `json:"pay"`
	Status     string         `gorm:"index" json:"status"`
}

type OilDataModel struct {
	DateString string  `json:"date_string"`
	BackupNum  string  `json:"backup_num"`
	NowNum     string  `json:"now_num"`
	OilType    string  `json:"oil_type"`
	OilNum     float64 `json:"oil_num"`
	Pay        float64 `json:"pay"`
	Status     string  `json:"status"`
}
