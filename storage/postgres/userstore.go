package postgres

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/skelego"
	"github.com/Rahul12344/skelego/services/storage/sqlservice"
	"golang.org/x/crypto/bcrypt"
)

// NewUserStore Postgresql client
func NewUserStore(client sqlservice.ORMWrapper, logger skelego.Logging) *UserStore {
	return &UserStore{
		client: client,
		logger: logger,
	}
}

// Get gets user for login
func (us *UserStore) Get(key string, password string) (*models.RecipeUser, error) {
	user, err := us.findUser(key, password)
	return user, err
}

func (us *UserStore) findUser(email, password string) (*models.RecipeUser, error) {
	user := &models.RecipeUser{}

	if err := us.client.ORM().Where("Email = ?", email).First(user).Error; err != nil {
		us.logger.LogError("Error retrieving user: %s", email)
		return nil, nil
	}

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		return nil, nil
	}

	return user, nil
}

// Add puts user into postgres
func (us *UserStore) Add(user *models.RecipeUser) (bool, error) {
	if !us.client.ORM().NewRecord(user) {
		return false, nil
	}
	us.client.ORM().Create(user)
	if us.client.ORM().NewRecord(user) {
		return false, nil
	}
	return true, nil
}

// Set sets updated fields
func (us *UserStore) Set(uuid string, userModded *models.RecipeUser) (*models.RecipeUser, error) {
	var user models.RecipeUser
	if err := us.client.ORM().Where("UUID = ?", uuid).Find(&user); err != nil {
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
	us.client.ORM().Save(&user)

	if err := us.client.ORM().Where("UUID = ?", uuid).Find(&user); err != nil {
		us.logger.LogError("Error modifying user: %s", uuid)
		return nil, nil
	}

	return &user, nil
}

// Remove dels clients
func (us *UserStore) Remove(key string, password string) (bool, error) {
	return true, nil
}

//AddRecipe adds recipe to user
func (us *UserStore) AddRecipe(userID string, recipe *models.Recipe) {
	us.client.ORM().Create(models.UserRecipes{
		Adder:    userID,
		RecipeID: recipe.RecipeID,
		Hits:     0,
	})
}

//RemoveRecipe removes recipe from user
func (us *UserStore) RemoveRecipe(userID string, recipe *models.Recipe) {
	us.client.ORM().Delete(models.UserRecipes{
		Adder: userID,
	})
}

//GetUserFromID gets user from uuid
func (us *UserStore) GetUserFromID(uuid string) *models.RecipeUser {
	user := &models.RecipeUser{}
	if err := us.client.ORM().Where("UUID = ?", uuid).First(user).Error; err != nil {
		us.logger.LogError("Error getting user from id: %s", uuid)
		return nil
	}

	return user
}
