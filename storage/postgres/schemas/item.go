package schemas

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/skelego"
	"github.com/Rahul12344/skelego/services/storage/sqlservice"
)

//NewItemSchema New schema
func NewItemSchema(orm sqlservice.ORMWrapper) ItemSchema {
	return ItemSchema{
		client: orm,
	}
}

//Migrate Migrates ItemSchema to database
func (is ItemSchema) Migrate(logger skelego.Logging) {
	if !is.client.ORM().HasTable(&models.Item{}) {
		is.client.ORM().AutoMigrate(&models.Item{})
		logger.LogEvent("Creating the following table: %s", is.TableName())
	} else {
		logger.LogEvent("Already created the following table: %s", is.TableName())
	}
}

//TableName Table name of schema
func (is ItemSchema) TableName() string {
	return "items"
}
