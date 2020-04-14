package models

import "github.com/jinzhu/gorm"

// Delivery delivery model
type Delivery struct {
	gorm.Model
	DelivererID string  `json:"DelivererID" gorm:"not null;primary_key"`
	DelivereeID string  `json:"DelivereeID" gorm:"not null;primary_key"`
	ItemID      string  `json:"ItemID" gorm:"not null;primary_key"`
	Quantity    int     `json:"Quantity" gorm:"not null"`
	DeliveryLon float64 `json:"DeliveryLon" gorm:"not null"`
	DeliveryLat float64 `json:"DeliveryLat" gorm:"not null"`
	StartLon    float64 `json:"StartLon" gorm:"not null"`
	StartLat    float64 `json:"StartLat" gorm:"not null"`
}
