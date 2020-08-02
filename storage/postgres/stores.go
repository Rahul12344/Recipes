package postgres

import (
	"github.com/Rahul12344/skelego/services/storage/sqlservice"
)

/* Stores are collections of related schemas */

// RecipeStore Unit for interfacing with recipe storage use cases
type RecipeStore struct {
	client sqlservice.ORMWrapper
}

//DeliveryStore Unit for interfacing with delivery storage use cases
type DeliveryStore struct {
	client sqlservice.ORMWrapper
}

// FriendStore Unit for interfacing with friend storage use cases
type FriendStore struct {
	client sqlservice.ORMWrapper
}

//ItemStore Unit for interfacing with item storage use cases
type ItemStore struct {
	client sqlservice.ORMWrapper
}

// UserStore Unit for interfacing with user storage use cases
type UserStore struct {
	client sqlservice.ORMWrapper
}
