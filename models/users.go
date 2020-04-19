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
	PhoneNumber string      `json:"phonenumber"`
	ProfilePic  string      `json:"profilepic"`
	Country     CountryCode `json:"country"`
	Lat         float64     `json:"lat"`
	Lon         float64     `json:"lon"`
}

//TableName name of table
func (RecipeUser) TableName() string {
	return "recipe_user"
}

//UserRecipes UserRecipes model
type UserRecipes struct {
	gorm.Model
	Adder   string `json:"adder" gorm:"not null;primary_key"`
	Recipes string `json:"recipes"`
	Hits    int    `json:"hits"`
}

//TableName name of table
func (UserRecipes) TableName() string {
	return "user_recipes"
}
