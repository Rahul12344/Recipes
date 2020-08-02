package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Rahul12344/Recipes/models"
	"github.com/gorilla/mux"
)

// RecipeService recipe service
type RecipeService interface {
	Find(ingredients []string) [][]*models.Recipe
	Ingredients(ingredients []string) []*models.Ingredients
	Quantities(quantities []string) []*models.Quantities
	Instructions(instructions []string, instructionNums []string) []*models.Instructions
	Tags(tags []string) []*models.Tags
	Image(filename string) []*models.Ingredients
	Insert(recipe *models.Recipe, ingredients []*models.Ingredients, instructions []*models.Instructions, tags []*models.Tags, quantities []*models.Quantities)
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
	r.HandleFunc("/createrecipe", rc.CreateRecipe).Methods("POST")
}

// FindRecipe finds recipe
func (rc *RecipeController) FindRecipe(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("image") != "" {
		rc.recipeService.Image(r.URL.Query().Get("image"))
	}

	rc.recipeService.Ingredients(strings.Split(r.URL.Query().Get("ingredients"), ","))
}

// CreateRecipe creates new recipe
func (rc *RecipeController) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	var ingredientsCS models.Ingredients
	var instructionCS models.Instructions
	var tagsCS models.Tags
	var quantitiesCS models.Quantities

	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&recipe)
	decoder.Decode(&ingredientsCS)
	decoder.Decode(&instructionCS)
	decoder.Decode(&tagsCS)
	decoder.Decode(&quantitiesCS)

	ingredients := rc.recipeService.Ingredients(strings.Split(ingredientsCS.Ingredient, ","))
	instructions := rc.recipeService.Instructions(strings.Split(instructionCS.Instruction, ","), strings.Split(instructionCS.InstructionNum, ","))
	tags := rc.recipeService.Tags(strings.Split(tagsCS.Tag, ","))
	quantities := rc.recipeService.Quantities(strings.Split(quantitiesCS.Quantity, ","))

	rc.recipeService.Insert(&recipe, ingredients, instructions, tags, quantities)
}
