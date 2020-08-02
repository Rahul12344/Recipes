package services

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/Rahul12344/Recipes/models"
	"github.com/dgrijalva/jwt-go"
)

//UserStore user store
type UserStore interface {
	Get(username string, password string) (*models.RecipeUser, error)
	Set(email string, uuid string, userModded *models.RecipeUser) (map[string]interface{}, error)
	Add(user *models.RecipeUser) (bool, error)
	Remove(username string, password string) (bool, error)
	GetUserFromID(uuid string) *models.RecipeUser
	AddRecipe(uuid string, recipe *models.Recipe)
	RemoveRecipe(uuid string, recipe *models.Recipe)
}

//UserIndex Index for users
type UserIndex interface {
	AddUsers(recipes ...models.RecipeUser) error
	GetUsers(ctx context.Context, index, qType string, conditional map[string]interface{}, offset, limit int) []models.RecipeUser
}

//UserService holds services for users
type UserService struct {
	userStore UserStore
	userIndex UserIndex
}

//NewUserService constructs new user service
func NewUserService(userStore UserStore, userIndex UserIndex) *UserService {
	return &UserService{
		userStore: userStore,
		userIndex: userIndex,
	}
}

//GetUser login
func (as *UserService) GetUser(username string, password string) (map[string]interface{}, string, string, time.Time, time.Time) {
	user, err := as.userStore.Get(username, password)
	if err != nil {

	}
	if user == nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials"}
		return resp, "", "", time.Time{}, time.Time{}
	}

	expiresAt := time.Now().Add(time.Minute * 15)

	tk := &models.Token{
		UUID:  user.UserID,
		Email: user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	refreshString, stringErr := generateRandomString(32)
	if stringErr != nil {

	}

	refreshExpiresAt := time.Now().Add(time.Hour * 144)

	refreshTk := &models.RefreshToken{
		UUID: user.UserID,
		ID:   refreshString,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: refreshExpiresAt.Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), refreshTk)

	refreshTokenString, error := refreshToken.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString
	resp["refresh"] = refreshTokenString
	resp["user"] = user

	return resp, tokenString, refreshTokenString, expiresAt, refreshExpiresAt
}

func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

//SetUser sets categories
func (as *UserService) SetUser(email string, uuid string, userModded *models.RecipeUser) (map[string]interface{}, error) {
	return as.userStore.Set(email, uuid, userModded)
}

//NewUser signs user in
func (as *UserService) NewUser(user *models.RecipeUser) (bool, error) {
	return as.userStore.Add(user)
}

//DeleteUser delete user
func (as *UserService) DeleteUser(username string, password string) (bool, error) {
	return as.userStore.Remove(username, password)
}

//RefreshToken refresh token
func (as *UserService) RefreshToken(uuid string) (map[string]interface{}, string, time.Time) {
	user := as.userStore.GetUserFromID(uuid)
	if user == nil {
		var resp = map[string]interface{}{"status": false, "message": "UUID not found"}
		return resp, "", time.Time{}
	}

	expiresAt := time.Now().Add(time.Minute * 15)

	tk := &models.Token{
		UUID:  user.UserID,
		Email: user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString
	resp["user"] = user

	return resp, tokenString, expiresAt
}

//AddRecipe add user recipe
func (as *UserService) AddRecipe(uuid string, recipe *models.Recipe) {
	as.userStore.AddRecipe(uuid, recipe)
}

//RemoveRecipe removes user recipe
func (as *UserService) RemoveRecipe(uuid string, recipe *models.Recipe) {
	as.userStore.RemoveRecipe(uuid, recipe)
}
