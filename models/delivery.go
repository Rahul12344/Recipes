package models

import "github.com/jinzhu/gorm"

// Delivery delivery model
type Delivery struct {
	gorm.Model
	DelivererID   string  `json:"deliverer_id" gorm:"not null;primary_key"`
	DelivereeID   string  `json:"deliveree_id" gorm:"not null;primary_key"`
	ItemID        string  `json:"item_id" gorm:"not null;primary_key"`
	Quantity      int     `json:"quantity" gorm:"not null"`
	DeliveryLon   float64 `json:"delivery_lon" gorm:"not null"`
	DeliveryLat   float64 `json:"delivery_lat" gorm:"not null"`
	StartLon      float64 `json:"start_lon" gorm:"not null"`
	StartLat      float64 `json:"start_lan" gorm:"not null"`
	Timestamp     int64   `json:"timestamp"`
	TimeDelivered int64   `json:"time_delivered"`
}

//TableName name of table
func (Delivery) TableName() string {
	return "delivery"
}
