package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// TODO - implement for privacy-forward implementation of this

// Token token model
type Token struct {
	jwt.Claims
	UUID           string
	Email          string
	StandardClaims *jwt.StandardClaims
}

// RefreshToken refresh token model
type RefreshToken struct {
	jwt.Claims
	UUID           string
	ID             string
	StandardClaims *jwt.StandardClaims
}

// User user model
type User struct {
	gorm.Model
	FirstName   string  `json:"fname"`
	LastName    string  `json:"lname"`
	Email       string  `json:"email"`
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	PhoneNumber string  `json:"phonenumber"`
	ProfilePic  string  `json:"profilepic"`
	UUID        string  `json:"uuid"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Interests   string  `json:"interests"`
	Clubs       string  `json:"clubs"`
	Bio         string  `json:"bio"`
	Recipes     string  `json:"recipes"`
}

//RecipePointer recipepointer model
type RecipePointer struct {
	Adder  string `json:"adder"`
	Recipe string `json:"recipe"`
	Hits   int    `json:"hits"`
}

// Friends friend model
type Friends struct {
	gorm.Model
	UUID         string `json:"uuid"`
	FUUID        string `json:"fuuid"`
	FReqMess     string `json:"msg"`
	TimeStamp    int64  `json:"timestamp"`
	Status       int    `json:"status"`
	TimeAccepted int64  `json:"timeaccepted"`
}

// FriendRequest friendrequest model
type FriendRequest struct {
	gorm.Model
	Sender    string `json:"uuid"`
	Receiver  string `json:"fuuid"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

// Recipe recipe model
type Recipe struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Ingredients  []string `gorm:"type:varchar(64)[]" json:"ingredients"`
	Instructions []string `gorm:"type:varchar(64)[]" json:"instructions"`
	Quantities   string   `json:"quantities"`
	Materials    string   `json:"materials"`
	Tag          string   `json:"tag"`
	UUID         string   `json:"uuid"`
}
