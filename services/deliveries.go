package services

import "github.com/Rahul12344/Recipes/models"

//DeliveryStore delivery store
type DeliveryStore interface {
	DELIVERY(deliveries ...*models.Delivery)
	REMOVE(deliveries ...*models.Delivery)
}

//IngredientStore ingredient store
type IngredientStore interface {
}

//DeliveryService delivery service
type DeliveryService struct {
	deliveryStore DeliveryStore
}

//NewDeliveryService constructs new delivery service
func NewDeliveryService(deliveryStore DeliveryStore) *DeliveryService {
	return &DeliveryService{
		deliveryStore: deliveryStore,
	}
}

//DELIVERIES create deliveries
func (ds *DeliveryService) DELIVERIES(deliverer, recipient string, lat, lon, startLan, startLon float64, ingredients ...string) {
	var deliveries []*models.Delivery
	for _, ingredient := range ingredients {
		deliveries = append(deliveries, &models.Delivery{
			ItemID:      ingredient,
			DelivereeID: recipient,
			DelivererID: deliverer,
			DeliveryLat: lat,
			DeliveryLon: lon,
			StartLat:    startLan,
			StartLon:    startLon,
		})
	}
	ds.deliveryStore.DELIVERY(deliveries...)
}

//COMPLETE completes deliveries
func (ds *DeliveryService) COMPLETE(deliverer, recipient string, ingredients ...string) {
	var deliveries []*models.Delivery
	for _, ingredient := range ingredients {
		deliveries = append(deliveries, &models.Delivery{
			ItemID:      ingredient,
			DelivereeID: recipient,
			DelivererID: deliverer,
		})
	}
	ds.deliveryStore.REMOVE(deliveries...)
}

//DELETE deletes deliveries
func (ds *DeliveryService) DELETE(deliverer, recipient string, ingredients ...string) {
	var deliveries []*models.Delivery
	for _, ingredient := range ingredients {
		deliveries = append(deliveries, &models.Delivery{
			ItemID:      ingredient,
			DelivereeID: recipient,
			DelivererID: deliverer,
		})
	}
	ds.deliveryStore.REMOVE(deliveries...)
}
