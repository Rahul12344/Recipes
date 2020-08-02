package elasticsearch

import (
	"context"
	"encoding/json"

	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/skelego/services/index"
)

//NewUserIndex Create new User index
func NewUserIndex(index index.Index) *UserIndex {
	return &UserIndex{
		index: index,
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
	var users []models.RecipeUser
	searchResult, _ := ui.index.ElasticSearch().Search().
		Index(index).
		Type(qType).
		Query(ui.index.Query(conditional)).
		From(offset).
		Size(limit).
		Do(ctx)
	ui.index.Query(conditional)
	for _, hit := range searchResult.Hits.Hits {
		var user models.RecipeUser
		err := json.Unmarshal(hit.Source, &user)
		if err != nil {
			return nil
		}
		users = append(users, user)
	}
	return users
}

func (ui *UserIndex) verifySchema(conditional map[string]interface{}) bool {
	return true
}
