package postgres

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/Recipes/util/errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// UserStore handles authorization flow for users and user functionality
type UserStore struct {
	client *gorm.DB
}

// NewUserStore Postgresql client
func NewUserStore(client *gorm.DB) *UserStore {
	const SchemaQuery = `CREATE SCHEMA IF NOT EXISTS users`
	client.Exec(SchemaQuery)
	//client.Exec(`set search_path='users'`)
	return &UserStore{
		client: client,
	}
}

func (us *UserStore) create() {
	/* TODO: Maybe change migration model to maybe define DB relationships */
	us.client.AutoMigrate(&models.RecipeUser{})
}

// GET gets user for login
func (us *UserStore) GET(key string, password string) (*models.RecipeUser, *errors.Errors) {
	user, err := us.findUser(key, password)
	return user, err
}

func (us *UserStore) findUser(email, password string) (*models.RecipeUser, *errors.Errors) {
	user := &models.RecipeUser{}

	if err := us.client.Where("Email = ?", email).First(user).Error; err != nil {
		return nil, nil
	}

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		return nil, nil
	}

	return user, nil
}

// PUT puts user into postgres
func (us *UserStore) PUT(user *models.RecipeUser) (bool, error) {
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
func (us *UserStore) SET(email string, uuid string, userModded *models.RecipeUser) (map[string]interface{}, error) {
	var user models.RecipeUser
	if err := us.client.Where("EMAIL = ? AND UUID = ?", email, uuid).Find(&user); err != nil {
	}

	fname := userModded.FirstName
	lname := userModded.LastName
	lat := userModded.Lat
	lon := userModded.Lon

	if userModded.FirstName == "" {
		fname = user.FirstName
	}
	if userModded.LastName == "" {
		lname = user.LastName
	}
	if userModded.Lat == 0 {
		lat = user.Lat
	}
	if userModded.Lon == 0 {
		lon = user.Lon
	}

	user.FirstName = fname
	user.LastName = lname
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

//GETFROMUUID gets user from uuid
func (us *UserStore) GETFROMUUID(uuid string) *models.RecipeUser {
	user := &models.RecipeUser{}
	if err := us.client.Where("UUID = ?", uuid).First(user).Error; err != nil {
		return nil
	}

	return user
}
