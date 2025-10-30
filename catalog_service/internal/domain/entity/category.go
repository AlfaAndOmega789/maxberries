package entity

import "time"

type Category struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
