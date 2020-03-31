package controllers

import (
	"net/http"
	"strings"

	"github.com/Rahul12344/Recipes/models"
	"github.com/gorilla/mux"
)

// RecipeService recipe service
type RecipeService interface {
	FIND(image bool, matches *models.Recipe) []*models.Recipe
	INGREDIENTS(ingredients []string) *models.Recipe
	IMAGE(filename string) *models.Recipe
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
		rc.recipeService.FIND(true, rc.recipeService.IMAGE(r.URL.Query().Get("image")))
	}

	rc.recipeService.FIND(false, rc.recipeService.INGREDIENTS(strings.Split(r.URL.Query().Get("ingredients"), ",")))
}
