package user_test

import (
	"testing"

	"github.com/Rahul12344/Recipes/storage/postgres"
	"github.com/Rahul12344/Recipes/util/hash"
	"github.com/Rahul12344/Recipes/util/uuid"
	"github.com/jinzhu/gorm"

	"github.com/Rahul12344/Recipes/models"
)

/* UNIT TESTS FOR USER FUNCTIONS */

func UserStoreTest(t *testing.T) {
	users := models.User{
		UUID:      uuid.UUID(),
		FirstName: "TestA",
		LastName:  "TestB",
		Email:     "test1@gmail.com",
		Username:  "TestA",
		Password:  hash.Hash("password"),
	}

	testDB := CreateTestDB()

	store := postgres.NewUserStore(testDB)
	store.PUT(&users)
	testGET, _ := store.GET(users.Email, "password")

	if users != *testGET {
		t.Errorf("Obtaining user record failed at insertion or retrieval by email and password")
	}

	testUUID := store.GETFROMUUID(users.UUID)
	if users != *testUUID {
		t.Errorf("Obtaining user record failed at insertion or retrieval by UUID")
	}

	DropTestDB(testDB)
}

func CreateTestDB() *gorm.DB {
	return nil
}

func DropTestDB(db *gorm.DB) {

}
