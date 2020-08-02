package controllers

//DeliveryService delivery service
type DeliveryService interface {
	Deliveries(deliverer, recipient string, lat, lon, startLan, startLon float64, ingredients ...string)
	CompleteDelivery(deliverer, recipient string, ingredients ...string)
	DeleteDelivery(deliverer, recipient string, ingredients ...string)
}

// DeliveryController controls delivery actions for users
type DeliveryController struct {
	deliveryService DeliveryService
}

//NewDeliveryController creates new delivery controller
func NewDeliveryController(deliveryService DeliveryService) *DeliveryController {
	return &DeliveryController{
		deliveryService: deliveryService,
	}
}
