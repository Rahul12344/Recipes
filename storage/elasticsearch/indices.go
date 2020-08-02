package elasticsearch

import "github.com/Rahul12344/skelego/services/index"

//RecipeIndex Recipe indexing system for ES
type RecipeIndex struct {
	index index.Index
}

//UserIndex User indexing system for ES
type UserIndex struct {
	index index.Index
}
