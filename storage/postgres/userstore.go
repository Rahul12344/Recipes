package postgres

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/Recipes/util/uuid"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// UserStore handles authorization flow for users and user functionality
type UserStore struct {
	client *gorm.DB
}

// NewUserStore Postgresql client
func NewUserStore(client *gorm.DB) *UserStore {
	return &UserStore{
		client: client,
	}
}

func (us *UserStore) create() {
	/* TODO: Maybe change migration model to maybe define DB relationships */
	us.client.AutoMigrate(&models.User{})
}

// GET gets user for login
func (us *UserStore) GET(key string, password string) (map[string]interface{}, string, string, time.Time, time.Time) {
	resp, token, refreshToken, expirationTime, refreshTokenExpirationTime := us.findUser(key, password)
	return resp, token, refreshToken, expirationTime, refreshTokenExpirationTime
}

func (us *UserStore) findUser(email, password string) (map[string]interface{}, string, string, time.Time, time.Time) {
	user := &models.User{}

	if err := us.client.Where("Email = ?", email).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp, "", "", time.Time{}, time.Time{}
	}

	expiresAt := time.Now().Add(time.Minute * 15)

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials"}
		return resp, "", "", time.Time{}, time.Time{}
	}

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

	refreshString, err := generateRandomString(32)
	if err != nil {

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

// PUT puts user into postgres
func (us *UserStore) PUT(email string, password string, firstname string, lastname string) (bool, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	user := &models.User{}
	user.Password = string(pass)
	user.Email = email
	user.FirstName = firstname
	user.LastName = lastname
	user.UUID = uuid.UUID()
	if !us.client.NewRecord(user) {
		return false, nil
	}
	us.client.Create(user)
	if us.client.NewRecord(user) {
		return false, nil
	}
	return true, nil
}

// SET sets updated fields
func (us *UserStore) SET(email string, uuid string, userModded *models.User) (map[string]interface{}, error) {
	var user models.User
	if err := us.client.Where("EMAIL = ? AND UUID = ?", email, uuid).Find(&user); err != nil {
	}

	fname := userModded.FirstName
	lname := userModded.LastName
	interests := userModded.Interests
	bio := userModded.Bio
	clubs := userModded.Clubs
	lat := userModded.Lat
	lon := userModded.Lon

	if userModded.FirstName == "" {
		fname = user.FirstName
	}
	if userModded.LastName == "" {
		lname = user.LastName
	}
	if userModded.Interests == "" {
		interests = user.Interests
	}
	if userModded.Bio == "" {
		bio = user.Bio
	}
	if userModded.Clubs == "" {
		bio = user.Clubs
	}
	if userModded.Lat == 0 {
		lat = user.Lat
	}
	if userModded.Lon == 0 {
		lon = user.Lon
	}

	user.FirstName = fname
	user.LastName = lname
	user.Interests = interests
	user.Bio = bio
	user.Clubs = clubs
	user.Lat = lat
	user.Lon = lon
	us.client.Save(&user)

	if err := us.client.Where("EMAIL = ? AND UUID = ?", email, uuid).Find(&user); err != nil {
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["mod_user"] = user

	return resp, nil
}

// DEL dels clients
func (us *UserStore) DEL(key string, password string) (bool, error) {
	return true, nil
}

//REFRESH generates a new refresh token for the authenticated client
func (us *UserStore) REFRESH(uuid string) (map[string]interface{}, string, time.Time) {
	user := &models.User{}
	if err := us.client.Where("UUID = ?", uuid).First(user).Error; err != nil {
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
