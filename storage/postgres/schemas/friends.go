package schemas

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/skelego"
	"github.com/Rahul12344/skelego/services/storage/sqlservice"
)

//NewFriendSchema New schema
func NewFriendSchema(orm sqlservice.ORMWrapper) FriendSchema {
	return FriendSchema{
		client: orm,
	}
}

//Migrate Migrates FriendSchema to database
func (fs FriendSchema) Migrate(logger skelego.Logging) {
	if !fs.client.ORM().HasTable(&models.Friends{}) {
		fs.client.ORM().AutoMigrate(&models.Friends{})
		fs.client.ORM().Model(&models.Friends{}).AddForeignKey("user_id", "recipe_users(user_id)", "RESTRICT", "RESTRICT")
		fs.client.ORM().Model(&models.Friends{}).AddForeignKey("friend_id", "recipe_users(user_id)", "RESTRICT", "RESTRICT")
		logger.LogEvent("Creating the following table: %s", fs.TableName())
	} else {
		logger.LogEvent("Already created the following table: %s", fs.TableName())
	}
}

//TableName Table name of schema
func (fs FriendSchema) TableName() string {
	return "friends"
}
