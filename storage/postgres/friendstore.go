package postgres

import (
	"time"

	"github.com/Rahul12344/skelego"

	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/skelego/services/storage/sqlservice"
)

// NewFriendStore Postgresql client
func NewFriendStore(client sqlservice.ORMWrapper, logger skelego.Logging) *FriendStore {
	return &FriendStore{
		client: client,
		logger: logger,
	}
}

// Follow Follows friend
func (fs *FriendStore) Follow(uuid string, friendUUID string, optionalMsg string) (bool, error) {
	friendRequest := &models.Friends{
		UserID:               uuid,
		FriendID:             friendUUID,
		FriendRequestMessage: optionalMsg,
		TimeStamp:            time.Now().Unix(),
		Status:               0,
	}
	if !fs.client.ORM().NewRecord(friendRequest) {
		return false, nil
	}
	fs.client.ORM().Create(friendRequest)
	if fs.client.ORM().NewRecord(friendRequest) {
		return false, nil
	}
	return true, nil
}

// Unfollow Unfollows friend
func (fs *FriendStore) Unfollow(uuid string, friendUUID string, optionalMsg string) (bool, error) {
	friendRequest := &models.Friends{
		UserID:               uuid,
		FriendID:             friendUUID,
		FriendRequestMessage: optionalMsg,
		TimeStamp:            time.Now().Unix(),
		Status:               0,
	}
	if fs.client.ORM().NewRecord(friendRequest) {
		return false, nil
	}
	fs.client.ORM().Delete(friendRequest)
	if !fs.client.ORM().NewRecord(friendRequest) {
		return false, nil
	}
	return true, nil
}

// Accept Accepts added friends
func (fs *FriendStore) Accept(currUUID string, friendUUID string) (*models.Friends, error) {
	friendRequest := &models.Friends{}
	friendRequest.UserID = currUUID
	friendRequest.FriendID = friendUUID

	fs.acceptRequest(currUUID, friendUUID)

	return friendRequest, nil
}

func (fs *FriendStore) acceptRequest(currUUID string, friendUUID string) {
	friend := &models.Friends{}

	if err := fs.client.ORM().Where("F_UUID = ? AND UUID = ?", currUUID, friendUUID).Find(&friend); err != nil {
	}
	friend.Status = 1
	friend.TimeAccepted = time.Now().Unix()
	fs.client.ORM().Save(&friend)
}
