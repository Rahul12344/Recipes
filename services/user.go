package services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/Recipes/util/errors"
	"github.com/dgrijalva/jwt-go"
)

//UserStore user store
type UserStore interface {
	GET(username string, password string) (*models.User, *errors.Errors)
	SET(email string, uuid string, userModded *models.User) (map[string]interface{}, error)
	PUT(user *models.User) (bool, error)
	DEL(username string, password string) (bool, error)
	GETFROMUUID(uuid string) *models.User
}

//UserRecipeStore user recipe store
type UserRecipeStore interface {
	ADD(uuid string, recipe *models.Recipe)
	REMOVE(uuid string, recipe *models.Recipe)
}

//UserService holds services for users
type UserService struct {
	userStore       UserStore
	userRecipeStore UserRecipeStore
}

//NewUserService constructs new user service
func NewUserService(userStore UserStore, userRecipeStore UserRecipeStore) *UserService {
	return &UserService{
		userStore:       userStore,
		userRecipeStore: userRecipeStore,
	}
}

//GET login
func (as *UserService) GET(username string, password string) (map[string]interface{}, string, string, time.Time, time.Time) {
	user, err := as.userStore.GET(username, password)
	if err != nil {

	}
	if user == nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials"}
		return resp, "", "", time.Time{}, time.Time{}
	}

	expiresAt := time.Now().Add(time.Minute * 15)

	tk := &models.Token{
		UUID:  user.UUID,
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
		UUID: user.UUID,
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

//SET sets categories
func (as *UserService) SET(email string, uuid string, userModded *models.User) (map[string]interface{}, error) {
	return as.userStore.SET(email, uuid, userModded)
}

//PUT signs user in
func (as *UserService) PUT(user *models.User) (bool, error) {
	return as.userStore.PUT(user)
}

//DEL delete user
func (as *UserService) DEL(username string, password string) (bool, error) {
	return as.userStore.DEL(username, password)
}

//REFRESH refresh token
func (as *UserService) REFRESH(uuid string) (map[string]interface{}, string, time.Time) {
	user := as.userStore.GETFROMUUID(uuid)
	if user == nil {
		var resp = map[string]interface{}{"status": false, "message": "UUID not found"}
		return resp, "", time.Time{}
	}

	expiresAt := time.Now().Add(time.Minute * 15)

	tk := &models.Token{
		UUID:  user.UUID,
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

//ADD add user recipe
func (as *UserService) ADD(uuid string, recipe *models.Recipe) {
	as.userRecipeStore.ADD(uuid, recipe)
}

//REMOVE removes user recipe
func (as *UserService) REMOVE(uuid string, recipe *models.Recipe) {
	as.userRecipeStore.REMOVE(uuid, recipe)
}
