package postgres

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/jinzhu/gorm"
)

//DeliveryStore stores deliveries
type DeliveryStore struct {
	client *gorm.DB
}

// NewDeliveryStore Postgresql client
func NewDeliveryStore(client *gorm.DB) *DeliveryStore {
	const SchemaQuery = `CREATE SCHEMA IF NOT EXISTS deliveries`
	client.Exec(SchemaQuery)
	client.Exec(`set search_path='deliveries'`)
	return &DeliveryStore{
		client: client,
	}
}

func (ds *DeliveryStore) create() {
	/* TODO: Maybe change migration model to maybe define DB relationships */
	gorm.DefaultTableNameHandler = func(db *gorm.DB, tableName string) string {
		return "deliveries." + tableName
	}
	ds.client.AutoMigrate(&models.Delivery{})
}

// DELIVERY initiates delivery
func (ds *DeliveryStore) DELIVERY(deliveries ...*models.Delivery) {
	for _, delivery := range deliveries {
		ds.client.Create(delivery)
	}
}

// REMOVE removes deliveries
func (ds *DeliveryStore) REMOVE(deliveries ...*models.Delivery) {
	for _, delivery := range deliveries {
		ds.client.Delete(delivery)
	}
}
