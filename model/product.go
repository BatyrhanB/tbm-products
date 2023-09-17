package model

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID           uint           `gorm:"primaryKey;autoIncrement"`
	Title        string         `gorm:"type:varchar(155);index;not null"`
	Description  string         `gorm:"type:text;not null"`
	Price        int64          `gorm:"type:bigint;not null"` // Note: GORM doesn't have a direct Decimal type, consider using integer types for money
	InitialPrice float64        `gorm:"type:decimal(20,2);not null"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
