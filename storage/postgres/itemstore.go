package postgres

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/skelego/services/storage/sqlservice"
)

// NewItemStore Postgresql client
func NewItemStore(client sqlservice.ORMWrapper) *ItemStore {
	return &ItemStore{
		client: client,
	}
}

//AddItems Adds items
func (is *ItemStore) AddItems(items ...*models.Item) {
	for item := range items {
		is.client.ORM().Create(item)
	}
}

//RemoveItems Removes items
func (is *ItemStore) RemoveItems(items *models.Item, condition ...interface{}) {
	is.client.ORM().Delete(items, condition)
}

//Find finds corresponding items
func (is *ItemStore) Find(radius, lat, lon float64, item *models.Item) []models.Item {
	stores := []models.Item{}
	is.client.ORM().Where("ItemName = ? AND (LAT - ?)^2 + (LON - ?)^2 <= ?", lat, lon, radius*radius, item).Order("PRICE").Find(&stores)
	return stores
}
