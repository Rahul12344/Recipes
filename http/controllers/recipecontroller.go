package controllers

import (
	"net/http"
	"strings"

	"github.com/Rahul12344/Recipes/models"
	"github.com/gorilla/mux"
)

// RecipeService recipe service
type RecipeService interface {
	FIND(ingredients []string) [][]*models.Recipe
	INGREDIENTS(ingredients []string) []*models.Ingredients
	IMAGE(filename string) []*models.Ingredients
}

// RecipeController handles recipe-related events
type RecipeController struct {
	recipeService RecipeService
}

//NewRecipeController creates new recicpe controller
func NewRecipeController(recipeService RecipeService) *RecipeController {
	return &RecipeController{
		recipeService: recipeService,
	}
}

// Setup sets up handlers
func (rc *RecipeController) Setup(r *mux.Router) {
	r.HandleFunc("/find", rc.FindRecipe).Methods("GET")
}

// FindRecipe finds recipe
func (rc *RecipeController) FindRecipe(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("image") != "" {
		rc.recipeService.IMAGE(r.URL.Query().Get("image"))
	}

	rc.recipeService.INGREDIENTS(strings.Split(r.URL.Query().Get("ingredients"), ","))
}
