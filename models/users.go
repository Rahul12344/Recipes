package models

import (
	"github.com/jinzhu/gorm"
)

// TODO - implement for privacy-forward implementation of this

//CountryCode country code
type CountryCode string

//Valid checks validity of country code
func (cc CountryCode) Valid() {
	if len(cc) > 2 {
		print("Invalid")
		return
	}
	print("Valid")
	return
}

// RecipeUser user model
type RecipeUser struct {
	gorm.Model
	UserID      string      `json:"user_id" gorm:"unique;not null;primary_key"`
	FirstName   string      `json:"first_name" gorm:"not null"`
	LastName    string      `json:"last_name" gorm:"not null"`
	Email       string      `json:"email" gorm:"not null"`
	Username    string      `json:"username" gorm:"not null"`
	Password    string      `json:"password" gorm:"not null"`
	PhoneNumber string      `json:"phonenumber" gorm:"not null"`
	ProfilePic  string      `json:"profilepic" gorm:"not null"`
	Country     CountryCode `json:"country" gorm:"not null"`
	Lat         float64     `json:"lat" gorm:"not null"`
	Lon         float64     `json:"lon" gorm:"not null"`
}

//TableName name of table
func (RecipeUser) TableName() string {
	return "recipe_users"
}

//UserRecipes UserRecipes model
type UserRecipes struct {
	gorm.Model
	Adder    string `json:"adder" gorm:"not null;primary_key"`
	RecipeID string `json:"recipe_id" gorm:"not null"`
	Hits     int    `json:"hits" gorm:"not null"`
}

//TableName name of table
func (UserRecipes) TableName() string {
	return "user_recipes"
}

//Inventory inventory model
type Inventory struct {
	gorm.Model
	UserID       string `json:"inventory_user" gorm:"not null;primary_key"`
	IngredientID string `json:"inventory_ingredient" gorm:"not null;primary_key" `
	Quantity     int    `json:"inventory_quantity" gorm:"not null"`
}

//TableName name of table
func (Inventory) TableName() string {
	return "inventory"
}
