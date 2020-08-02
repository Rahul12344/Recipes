package schemas

import (
	"github.com/Rahul12344/skelego/services/storage/sqlservice"
)

//DeliverySchema Schema for delivery object
type DeliverySchema struct {
	client sqlservice.ORMWrapper
}

//ItemSchema Schema for item object
type ItemSchema struct {
	client sqlservice.ORMWrapper
}

//UserSchema Schema for user object
type UserSchema struct {
	client sqlservice.ORMWrapper
}

//UserRecipesSchema Schema for user recipes object
type UserRecipesSchema struct {
	client sqlservice.ORMWrapper
}

//InventorySchema Schema for user inventory object
type InventorySchema struct {
	client sqlservice.ORMWrapper
}

//ChatRoomSchema Schema for chat object
type ChatRoomSchema struct {
	client sqlservice.ORMWrapper
}

//RecipeSchema Schema for recipe object
type RecipeSchema struct {
	client sqlservice.ORMWrapper
}

//FriendSchema Schema for friend object
type FriendSchema struct {
	client sqlservice.ORMWrapper
}

//RecipeIngredientsSchema Schema for recipe ingredient object
type RecipeIngredientsSchema struct {
	client sqlservice.ORMWrapper
}

//IngredientsSchema Schema for ingredients object
type IngredientsSchema struct {
	client sqlservice.ORMWrapper
}

//TagsSchema Schema for tags object
type TagsSchema struct {
	client sqlservice.ORMWrapper
}

//InstructionsSchema Schema for recipe instructions object
type InstructionsSchema struct {
	client sqlservice.ORMWrapper
}

//QuantitiesSchema Schema for recipe quantities object
type QuantitiesSchema struct {
	client sqlservice.ORMWrapper
}
