package schemas

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/skelego"
	"github.com/Rahul12344/skelego/services/storage/sqlservice"
)

//NewRecipeSchema New schema
func NewRecipeSchema(orm sqlservice.ORMWrapper) RecipeSchema {
	return RecipeSchema{
		client: orm,
	}
}

//Migrate Migrates RecipeSchema to database
func (rs RecipeSchema) Migrate(logger skelego.Logging) {
	if !rs.client.ORM().HasTable(&models.Recipe{}) {
		rs.client.ORM().AutoMigrate(&models.Recipe{})
		logger.LogEvent("Creating the following table: %s", rs.TableName())
	} else {
		logger.LogEvent("Already created the following table: %s", rs.TableName())
	}
}

//TableName Table name of schema
func (rs RecipeSchema) TableName() string {
	return "recipe"
}

//NewRecipeIngredientsSchema New schema
func NewRecipeIngredientsSchema(orm sqlservice.ORMWrapper) RecipeIngredientsSchema {
	return RecipeIngredientsSchema{
		client: orm,
	}
}

//Migrate Migrates RecipeIngredientsSchema to database
func (ris RecipeIngredientsSchema) Migrate(logger skelego.Logging) {
	if !ris.client.ORM().HasTable(&models.RecipeIngredients{}) {
		ris.client.ORM().AutoMigrate(&models.RecipeIngredients{})
		ris.client.ORM().Model(&models.RecipeIngredients{}).AddForeignKey("recipe_id", "recipes(recipe_id)", "RESTRICT", "RESTRICT")
		ris.client.ORM().Model(&models.RecipeIngredients{}).AddForeignKey("ingredient_id", "ingredients(ingredient_id)", "RESTRICT", "RESTRICT")
		ris.client.ORM().Model(&models.RecipeIngredients{}).AddForeignKey("quantity_id", "quantities(quantity_id)", "RESTRICT", "RESTRICT")
		logger.LogEvent("Creating the following table: %s", ris.TableName())
	} else {
		logger.LogEvent("Already created the following table: %s", ris.TableName())
	}
}

//TableName Table name of schema
func (ris RecipeIngredientsSchema) TableName() string {
	return "recipe_ingredients"
}

//NewIngredientsSchema New schema
func NewIngredientsSchema(orm sqlservice.ORMWrapper) IngredientsSchema {
	return IngredientsSchema{
		client: orm,
	}
}

//Migrate Migrates IngredientsSchema to database
func (is IngredientsSchema) Migrate(logger skelego.Logging) {
	if !is.client.ORM().HasTable(&models.Ingredients{}) {
		is.client.ORM().AutoMigrate(&models.Ingredients{})
		logger.LogEvent("Creating the following table: %s", is.TableName())
	} else {
		logger.LogEvent("Already created the following table: %s", is.TableName())
	}
}

//TableName Table name of schema
func (is IngredientsSchema) TableName() string {
	return "ingredients"
}

//NewInstructionsSchema New schema
func NewInstructionsSchema(orm sqlservice.ORMWrapper) InstructionsSchema {
	return InstructionsSchema{
		client: orm,
	}
}

//Migrate Migrates InstructionsSchema to database
func (is InstructionsSchema) Migrate(logger skelego.Logging) {
	if !is.client.ORM().HasTable(&models.Instructions{}) {
		is.client.ORM().AutoMigrate(&models.Instructions{})
		is.client.ORM().Model(&models.Instructions{}).AddForeignKey("recipe_id", "recipes(recipe_id)", "RESTRICT", "RESTRICT")
		logger.LogEvent("Creating the following table: %s", is.TableName())
	} else {
		logger.LogEvent("Already created the following table: %s", is.TableName())
	}
}

//TableName Table name of schema
func (is InstructionsSchema) TableName() string {
	return "instructions"
}

//NewQuantitiesSchema New schema
func NewQuantitiesSchema(orm sqlservice.ORMWrapper) QuantitiesSchema {
	return QuantitiesSchema{
		client: orm,
	}
}

//Migrate Migrates QuantitiesSchema to database
func (qs QuantitiesSchema) Migrate(logger skelego.Logging) {
	if !qs.client.ORM().HasTable(&models.Quantities{}) {
		qs.client.ORM().AutoMigrate(&models.Quantities{})
		logger.LogEvent("Creating the following table: %s", qs.TableName())
	} else {
		logger.LogEvent("Already created the following table: %s", qs.TableName())
	}
}

//TableName Table name of schema
func (qs QuantitiesSchema) TableName() string {
	return "quantities"
}

//NewTagsSchema New schema
func NewTagsSchema(orm sqlservice.ORMWrapper) TagsSchema {
	return TagsSchema{
		client: orm,
	}
}

//Migrate Migrates TagsSchema to database
func (ts TagsSchema) Migrate(logger skelego.Logging) {
	if !ts.client.ORM().HasTable(&models.Tags{}) {
		ts.client.ORM().AutoMigrate(&models.Tags{})
		ts.client.ORM().Model(&models.Tags{}).AddForeignKey("recipe_id", "recipes(recipe_id)", "RESTRICT", "RESTRICT")
		logger.LogEvent("Creating the following table: %s", ts.TableName())
	} else {
		logger.LogEvent("Already created the following table: %s", ts.TableName())
	}
}

//TableName Table name of schema
func (ts TagsSchema) TableName() string {
	return "tags"
}
