package elasticsearch

import (
	"context"

	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/Recipes/storage/elasticsearch/queries"
	"github.com/Rahul12344/skelego"
	"github.com/Rahul12344/skelego/services/index"
)

//NewUserIndex Create new User index
func NewUserIndex(index index.Index, logger skelego.Logging) *UserIndex {
	return &UserIndex{
		index:  index,
		logger: logger,
	}
}

//AddUsers Adds users to index
func (ui *UserIndex) AddUsers(recipes ...models.RecipeUser) error {
	return nil
}

//GetUsers Searches on User index
func (ui *UserIndex) GetUsers(ctx context.Context, index, qType string, conditional map[string]interface{}, offset, limit int) []models.RecipeUser {
	if !ui.verifySchema(conditional) {
		return nil
	}
	hits := ui.index.SearchIndex(ctx, "users", ui.index.Query(queries.CreateRecipeQueryForArgs(conditional), ui.logger), ui.logger)
	if hits == nil {
		return nil
	}
	return nil
}

func (ui *UserIndex) verifySchema(conditional map[string]interface{}) bool {
	return true
}
