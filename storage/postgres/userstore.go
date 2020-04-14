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
	client.Exec(`set search_path='users'`)
	return &UserStore{
		client: client,
	}
}

func (us *UserStore) create() {
	/* TODO: Maybe change migration model to maybe define DB relationships */
	/*gorm.DefaultTableNameHandler = func(db *gorm.DB, tableName string) string {
		return "users." + tableName
	}*/
	us.client.AutoMigrate(&models.User{})
}

// GET gets user for login
func (us *UserStore) GET(key string, password string) (*models.User, *errors.Errors) {
	user, err := us.findUser(key, password)
	return user, err
}

func (us *UserStore) findUser(email, password string) (*models.User, *errors.Errors) {
	user := &models.User{}

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
func (us *UserStore) PUT(user *models.User) (bool, error) {
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
func (us *UserStore) GETFROMUUID(uuid string) *models.User {
	user := &models.User{}
	if err := us.client.Where("UUID = ?", uuid).First(user).Error; err != nil {
		return nil
	}

	return user
}
