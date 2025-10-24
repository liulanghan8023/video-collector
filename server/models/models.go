package models

import "time"

// Product represents the products table
type Product struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Status    string    `json:"status"`
	VideoCount int64    `json:"video_count" gorm:"-"`
	CreatedAt time.Time `json:"created_at"`
}

// Video represents the videos table
type Video struct {
	ID        uint      `gorm:"primaryKey"`
	URL       string    `gorm:"type:text" json:"url"`
	ProductID uint      `json:"product_id"`
	Status    string    `json:"status"`
	Product   Product   `gorm:"foreignKey:ProductID"`
	CreatedAt time.Time `json:"created_at"`
}
