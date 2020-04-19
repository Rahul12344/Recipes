package postgres

import (
	"time"

	"github.com/Rahul12344/Recipes/models"
	"github.com/jinzhu/gorm"
)

// FriendStore friend store
type FriendStore struct {
	client *gorm.DB
}

// NewFriendStore Postgresql client
func NewFriendStore(client *gorm.DB) *FriendStore {
	return &FriendStore{
		client: client,
	}
}

func (fs *FriendStore) create() {
	/* TODO: Maybe change migration model to maybe define DB relationships */
	/*const SchemaQuery = `CREATE SCHEMA IF NOT EXISTS users`
	fs.client.Exec(SchemaQuery)
	fs.client.Exec(`set search_path='users'`)*/
	fs.client.AutoMigrate(&models.Friends{})
}

// FOLLOW follows friend
func (fs *FriendStore) FOLLOW(uuid string, friendUUID string, optionalMsg string) (bool, error) {
	friendRequest := &models.Friends{
		UserID:               uuid,
		FriendID:             friendUUID,
		FriendRequestMessage: optionalMsg,
		TimeStamp:            time.Now().Unix(),
		Status:               0,
	}
	if !fs.client.NewRecord(friendRequest) {
		return false, nil
	}
	fs.client.Create(friendRequest)
	if fs.client.NewRecord(friendRequest) {
		return false, nil
	}
	return true, nil
}

// UNFOLLOW unfollows friend
func (fs *FriendStore) UNFOLLOW(uuid string, friendUUID string, optionalMsg string) (bool, error) {
	friendRequest := &models.Friends{
		UserID:               uuid,
		FriendID:             friendUUID,
		FriendRequestMessage: optionalMsg,
		TimeStamp:            time.Now().Unix(),
		Status:               0,
	}
	if fs.client.NewRecord(friendRequest) {
		return false, nil
	}
	fs.client.Delete(friendRequest)
	if !fs.client.NewRecord(friendRequest) {
		return false, nil
	}
	return true, nil
}

// ACCEPT accepts added friends
func (fs *FriendStore) ACCEPT(currUUID string, friendUUID string) (*models.Friends, error) {
	friendRequest := &models.Friends{}
	friendRequest.UserID = currUUID
	friendRequest.FriendID = friendUUID

	fs.acceptRequest(currUUID, friendUUID)

	return friendRequest, nil
}

func (fs *FriendStore) acceptRequest(currUUID string, friendUUID string) {
	friend := &models.Friends{}

	if err := fs.client.Where("F_UUID = ? AND UUID = ?", currUUID, friendUUID).Find(&friend); err != nil {
	}
	friend.Status = 1
	friend.TimeAccepted = time.Now().Unix()
	fs.client.Save(&friend)
}
