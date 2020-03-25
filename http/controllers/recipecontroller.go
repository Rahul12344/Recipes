package controllers

import "github.com/Rahul12344/Recipes/models"

// RecipeService recipe service
type RecipeService interface {
	FIND(matches *models.Recipe) []*models.Recipe
}

// RecipeController handles recipe-related events
type RecipeController struct {
	Service RecipeService
}
