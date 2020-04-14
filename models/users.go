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

// User user model
type User struct {
	gorm.Model
	FirstName   string      `json:"fname" gorm:"not null"`
	LastName    string      `json:"lname" gorm:"not null"`
	Email       string      `json:"email" gorm:"not null"`
	Username    string      `json:"username" gorm:"not null"`
	Password    string      `json:"password" gorm:"not null"`
	PhoneNumber string      `json:"phonenumber"`
	ProfilePic  string      `json:"profilepic"`
	UUID        string      `json:"uuid" gorm:"not null;primary_key"`
	Country     CountryCode `json:"country"`
	Lat         float64     `json:"lat"`
	Lon         float64     `json:"lon"`
}

//RecipePointer recipepointer model
type RecipePointer struct {
	gorm.Model
	Adder  string `json:"adder" gorm:"not null;primary_key"`
	Recipe string `json:"recipe" gorm:"not null;primary_key"`
	Hits   int    `json:"hits"`
}
