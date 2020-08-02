package services

import "github.com/Rahul12344/Recipes/models"

//FriendStore friend store
type FriendStore interface {
	Follow(uuid string, friendUUID string, optionalMsg string) (bool, error)
	Unfollow(uuid string, friendUUID string, optionalMsg string) (bool, error)
	Accept(currUUID string, friendUUID string) (*models.Friends, error)
}

//FriendService holds services for friends
type FriendService struct {
	friendStore FriendStore
}

//NewFriendService constructs new friend service
func NewFriendService(friendStore FriendStore) *FriendService {
	return &FriendService{
		friendStore: friendStore,
	}
}

//Follow follows user
func (fs *FriendService) Follow(uuid string, friendUUID string, optionalMsg string) (bool, error) {
	return fs.friendStore.Follow(uuid, friendUUID, optionalMsg)
}

//Unfollow unfollows user
func (fs *FriendService) Unfollow(uuid string, friendUUID string, optionalMsg string) (bool, error) {
	return fs.friendStore.Unfollow(uuid, friendUUID, optionalMsg)
}

//Accept accepts user
func (fs *FriendService) Accept(currUUID string, friendUUID string) (*models.Friends, error) {
	return fs.friendStore.Accept(currUUID, friendUUID)
}
