package postgres

import (
	"github.com/Rahul12344/skelego"
	"github.com/Rahul12344/skelego/services/storage/sqlservice"
)

/* Stores are collections of related schemas */

// RecipeStore Unit for interfacing with recipe storage use cases
type RecipeStore struct {
	client sqlservice.ORMWrapper
	logger skelego.Logging
}

//DeliveryStore Unit for interfacing with delivery storage use cases
type DeliveryStore struct {
	client sqlservice.ORMWrapper
	logger skelego.Logging
}

// FriendStore Unit for interfacing with friend storage use cases
type FriendStore struct {
	client sqlservice.ORMWrapper
	logger skelego.Logging
}

//ItemStore Unit for interfacing with item storage use cases
type ItemStore struct {
	client sqlservice.ORMWrapper
	logger skelego.Logging
}

// UserStore Unit for interfacing with user storage use cases
type UserStore struct {
	client sqlservice.ORMWrapper
	logger skelego.Logging
}
