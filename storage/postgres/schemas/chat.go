package schemas

import (
	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/skelego"
	"github.com/Rahul12344/skelego/services/storage/sqlservice"
)

//NewChatRoomSchema New schema
func NewChatRoomSchema(orm sqlservice.ORMWrapper) ChatRoomSchema {
	return ChatRoomSchema{
		client: orm,
	}
}

//Migrate Migrates ChatRoomSchema to db
func (crs ChatRoomSchema) Migrate(logger skelego.Logging) {
	if !crs.client.ORM().HasTable(&models.ChatRoom{}) {
		crs.client.ORM().AutoMigrate(&models.ChatRoom{})
		crs.client.ORM().Model(&models.ChatRoom{}).AddForeignKey("user_id", "recipe_users(user_id)", "RESTRICT", "RESTRICT")
		logger.LogEvent("Creating the following table: %s", crs.TableName())
	} else {
		logger.LogEvent("Already created the following table: %s", crs.TableName())
	}
}

//TableName Table name of schema
func (crs ChatRoomSchema) TableName() string {
	return "chats"
}
