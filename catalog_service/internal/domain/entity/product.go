package entity

import "time"

type Product struct {
	ID           string    `gorm:"primaryKey;autoIncrement"`
	Name         string    `gorm:"type:varchar(255);not null"`
	Descriptions string    `gorm:"type:text"`
	Price        float64   `gorm:"not null"`
	CategoryID   string    `gorm:"default:null"`
	CreateAt     time.Time `gorm:"autoCreateTime"`
}
