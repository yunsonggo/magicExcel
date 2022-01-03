package model

import (
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID        int64          `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoCreateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name      string         `gorm:"index" json:"name"`
	Password  string         `gorm:"index" json:"password"`
}
