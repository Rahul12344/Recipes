package schemas

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/skelego"
	"github.com/Rahul12344/skelego/services/storage/sqlservice"
)

//NewDeliverySchema New schema
func NewDeliverySchema(orm sqlservice.ORMWrapper) DeliverySchema {
	return DeliverySchema{
		client: orm,
	}
}

//Migrate Migrates DeliverySchema to database
func (ds DeliverySchema) Migrate(logger skelego.Logging) {
	if !ds.client.ORM().HasTable(&models.Delivery{}) {
		ds.client.ORM().AutoMigrate(&models.Delivery{})
		ds.client.ORM().Model(&models.Delivery{}).AddForeignKey("deliverer_id", "recipe_users(user_id)", "RESTRICT", "RESTRICT")
		ds.client.ORM().Model(&models.Delivery{}).AddForeignKey("deliveree_id", "recipe_users(user_id)", "RESTRICT", "RESTRICT")
		ds.client.ORM().Model(&models.Delivery{}).AddForeignKey("item_id", "items(item_id)", "RESTRICT", "RESTRICT")
		logger.LogEvent("Creating the following table: %s", ds.TableName())
	} else {
		logger.LogEvent("Already created the following table: %s", ds.TableName())
	}
}

//TableName Table name of schema
func (ds DeliverySchema) TableName() string {
	return "deliveries"
}
