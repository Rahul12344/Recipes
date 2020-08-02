package postgres

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/skelego/services/storage/sqlservice"
)

// NewDeliveryStore Postgresql client
func NewDeliveryStore(client sqlservice.ORMWrapper) *DeliveryStore {
	return &DeliveryStore{
		client: client,
	}
}

// AddDelivery Initiates delivery
func (ds *DeliveryStore) AddDelivery(deliveries ...*models.Delivery) {
	for _, delivery := range deliveries {
		ds.client.ORM().Create(delivery)
	}
}

// RemoveDelivery Removes deliveries
func (ds *DeliveryStore) RemoveDelivery(deliveries ...*models.Delivery) {
	for _, delivery := range deliveries {
		ds.client.ORM().Delete(delivery)
	}
}
