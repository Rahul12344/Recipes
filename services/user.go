package services

import (
	"time"

	"github.com/Rahul12344/Recipes/models"
)

//UserStore user store
type UserStore interface {
	GET(username string, password string) (map[string]interface{}, string, string, time.Time, time.Time)
	SET(email string, uuid string, userModded *models.User) (map[string]interface{}, error)
	PUT(email string, password string, firstName string, lastName string) (bool, error)
	DEL(username string, password string) (bool, error)
	REFRESH(uuid string) (map[string]interface{}, string, time.Time)
}

//UserRecipeStore user recipe store
type UserRecipeStore interface {
	ADD(uuid string, recipe *models.Recipe)
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
	return as.userStore.GET(username, password)
}

//SET sets categories
func (as *UserService) SET(email string, uuid string, userModded *models.User) (map[string]interface{}, error) {
	return as.userStore.SET(email, uuid, userModded)
}

//PUT signs user in
func (as *UserService) PUT(email string, password string, firstName string, lastName string) (bool, error) {
	return as.userStore.PUT(email, password, firstName, lastName)
}

//DEL delete user
func (as *UserService) DEL(username string, password string) (bool, error) {
	return as.userStore.DEL(username, password)
}

//REFRESH refresh token
func (as *UserService) REFRESH(uuid string) (map[string]interface{}, string, time.Time) {
	return as.userStore.REFRESH(uuid)
}

//ADD add user recipe
func (as *UserService) ADD(uuid string, recipe *models.Recipe) {
	as.userRecipeStore.ADD(uuid, recipe)
}
