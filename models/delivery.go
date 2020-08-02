package models

import (
	"github.com/Rahul12344/Recipes/util/uuid"
	"github.com/jinzhu/gorm"
)

// Delivery delivery model
type Delivery struct {
	gorm.Model
	DeliveryID    string  `json:"delivery_id" gorm:"not null;primary_key;unique"`
	DelivererID   string  `json:"deliverer_id" gorm:"not null"`
	DelivereeID   string  `json:"deliveree_id" gorm:"not null"`
	ItemID        string  `json:"item_id" gorm:"not null"`
	Quantity      int     `json:"quantity" gorm:"not null"`
	DeliveryLon   float64 `json:"delivery_lon" gorm:"not null"`
	DeliveryLat   float64 `json:"delivery_lat" gorm:"not null"`
	StartLon      float64 `json:"start_lon" gorm:"not null"`
	StartLat      float64 `json:"start_lan" gorm:"not null"`
	Timestamp     int64   `json:"timestamp" gorm:"not null"`
	TimeDelivered int64   `json:"time_delivered" gorm:"not null"`
}

//TableName name of table
func (Delivery) TableName() string {
	return "deliveries"
}

//NewDelivery creates new delivery
func NewDelivery(delivereeID, delivererID, itemID string, quantity int, deliveryLon, deliveryLat, startLon, startLat float64, timestamp int64) *Delivery {
	return &Delivery{
		DeliveryID:  uuid.UUID(),
		DelivereeID: delivereeID,
		DelivererID: delivererID,
		ItemID:      itemID,
		Quantity:    quantity,
		DeliveryLon: deliveryLon,
		DeliveryLat: deliveryLat,
		StartLon:    startLon,
		StartLat:    startLat,
		Timestamp:   timestamp,
	}
}
