package model

import (
	"gorm.io/gorm"
	"time"
)

type FileNameModel struct {
	ID            int64          `gorm:"primarykey" json:"id"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoCreateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name string `gorm:"index" json:"name"`
	Type string `gorm:"index" json:"type"`
	Path string `json:"path"`
}
