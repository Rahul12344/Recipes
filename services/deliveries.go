package services

import "github.com/Rahul12344/Recipes/models"

//Order contains individual order info
type Order struct {
	ItemID   string
	Quantity int
}

//DeliveryStore delivery store
type DeliveryStore interface {
	AddDelivery(deliveries ...*models.Delivery)
	RemoveDelivery(deliveries ...*models.Delivery)
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

//Deliveries create deliveries
func (ds *DeliveryService) Deliveries(deliverer, recipient string, lat, lon, startLan, startLon float64, orders ...Order) {
	var deliveries []*models.Delivery
	for _, order := range orders {
		deliveries = append(deliveries, &models.Delivery{
			ItemID:      order.ItemID,
			Quantity:    order.Quantity,
			DelivereeID: recipient,
			DelivererID: deliverer,
			DeliveryLat: lat,
			DeliveryLon: lon,
			StartLat:    startLan,
			StartLon:    startLon,
		})
	}
	ds.deliveryStore.AddDelivery(deliveries...)
}

//CompleteDelivery completes deliveries
func (ds *DeliveryService) CompleteDelivery(deliverer, recipient string, ingredients ...string) {
	var deliveries []*models.Delivery
	for _, ingredient := range ingredients {
		deliveries = append(deliveries, &models.Delivery{
			ItemID:      ingredient,
			DelivereeID: recipient,
			DelivererID: deliverer,
		})
	}
	ds.deliveryStore.RemoveDelivery(deliveries...)
}

//DeleteDelivery deletes deliveries
func (ds *DeliveryService) DeleteDelivery(deliverer, recipient string, ingredients ...string) {
	var deliveries []*models.Delivery
	for _, ingredient := range ingredients {
		deliveries = append(deliveries, &models.Delivery{
			ItemID:      ingredient,
			DelivereeID: recipient,
			DelivererID: deliverer,
		})
	}
	ds.deliveryStore.RemoveDelivery(deliveries...)
}
