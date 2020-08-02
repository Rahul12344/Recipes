package schemas

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/skelego"
	"github.com/Rahul12344/skelego/services/storage/sqlservice"
)

//NewUserSchema New schema
func NewUserSchema(orm sqlservice.ORMWrapper) UserSchema {
	return UserSchema{
		client: orm,
	}
}

//Migrate Migrates UserSchema to database
func (us UserSchema) Migrate(logger skelego.Logging) {
	if !us.client.ORM().HasTable(&models.RecipeUser{}) {
		us.client.ORM().AutoMigrate(&models.RecipeUser{})
		logger.LogEvent("Creating the following table: %s", us.TableName())
	} else {
		logger.LogEvent("Already created the following table: %s", us.TableName())
	}
}

//TableName Table name of schema
func (us UserSchema) TableName() string {
	return "users"
}

//NewUserRecipesSchema New schema
func NewUserRecipesSchema(orm sqlservice.ORMWrapper) UserRecipesSchema {
	return UserRecipesSchema{
		client: orm,
	}
}

//Migrate Migrates UserRecipesSchema to database
func (urs UserRecipesSchema) Migrate(logger skelego.Logging) {
	if !urs.client.ORM().HasTable(&models.UserRecipes{}) {
		urs.client.ORM().AutoMigrate(&models.UserRecipes{})
		urs.client.ORM().Model(&models.UserRecipes{}).AddForeignKey("recipe_id", "recipes(recipe_id)", "RESTRICT", "RESTRICT")
		logger.LogEvent("Creating the following table: %s", urs.TableName())
	} else {
		logger.LogEvent("Already created the following table: %s", urs.TableName())
	}
}

//TableName Table name of schema
func (urs UserRecipesSchema) TableName() string {
	return "user_recipes"
}

//NewInventorySchema New schema
func NewInventorySchema(orm sqlservice.ORMWrapper) InventorySchema {
	return InventorySchema{
		client: orm,
	}
}

//Migrate Migrates InventorySchema to database
func (is InventorySchema) Migrate(logger skelego.Logging) {
	if !is.client.ORM().HasTable(&models.Inventory{}) {
		is.client.ORM().AutoMigrate(&models.Inventory{})
		is.client.ORM().Model(&models.Inventory{}).AddForeignKey("user_id", "recipe_users(user_id)", "RESTRICT", "RESTRICT")
		is.client.ORM().Model(&models.Inventory{}).AddForeignKey("ingredient_id", "ingredients(ingredient_id)", "RESTRICT", "RESTRICT")
		logger.LogEvent("Creating the following table: %s", is.TableName())
	} else {
		logger.LogEvent("Already created the following table: %s", is.TableName())
	}
}

//TableName Table name of schema
func (is InventorySchema) TableName() string {
	return "inventory"
}
