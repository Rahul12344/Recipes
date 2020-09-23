package elasticsearch

import (
	"github.com/Rahul12344/skelego"
	"github.com/Rahul12344/skelego/services/index"
)

//RecipeIndex Recipe indexing system for ES
type RecipeIndex struct {
	index  index.Index
	logger skelego.Logging
}

//UserIndex User indexing system for ES
type UserIndex struct {
	index  index.Index
	logger skelego.Logging
}
