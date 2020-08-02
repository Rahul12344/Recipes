package elasticsearch

import (
	"context"
	"encoding/json"

	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/skelego/services/index"
)

//NewRecipeIndex Create new Recipe index
func NewRecipeIndex(index index.Index) *RecipeIndex {
	return &RecipeIndex{
		index: index,
	}
}

//AddRecipes Adds recipes to index
func (ri *RecipeIndex) AddRecipes(recipes ...models.TotalRecipe) error {
	return nil
}

//GetRecipes Searches on Recipe index
func (ri *RecipeIndex) GetRecipes(ctx context.Context, index, qType string, conditional map[string]interface{}, offset, limit int) []models.TotalRecipe {
	if !ri.verifySchema(conditional) {
		return nil
	}
	var recipes []models.TotalRecipe
	searchResult, _ := ri.index.ElasticSearch().Search().
		Index(index).
		Type(qType).
		Query(ri.index.Query(conditional)).
		From(offset).
		Size(limit).
		Do(ctx)
	ri.index.Query(conditional)
	for _, hit := range searchResult.Hits.Hits {
		var recipe models.TotalRecipe
		err := json.Unmarshal(hit.Source, &recipe)
		if err != nil {
			return nil
		}
		recipes = append(recipes, recipe)
	}
	return recipes
}

//QueryBuilder Builds query
func (ri *RecipeIndex) QueryBuilder(ingredients ...string) map[string]interface{} {
	return nil
}

//VerifySchema Verifies query
func (ri *RecipeIndex) verifySchema(conditional map[string]interface{}) bool {
	return true
}
