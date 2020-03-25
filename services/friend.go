package services

import "github.com/Rahul12344/Recipes/models"

//FriendStore friend store
type FriendStore interface {
	FOLLOW(uuid string, friendUUID string, optionalMsg string) (bool, error)
	UNFOLLOW(uuid string, friendUUID string, optionalMsg string) (bool, error)
	ACCEPT(currUUID string, friendUUID string) (*models.Friends, error)
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

//FOLLOW follows user
func (fs *FriendService) FOLLOW(uuid string, friendUUID string, optionalMsg string) (bool, error) {
	return fs.friendStore.FOLLOW(uuid, friendUUID, optionalMsg)
}

//UNFOLLOW unfollows user
func (fs *FriendService) UNFOLLOW(uuid string, friendUUID string, optionalMsg string) (bool, error) {
	return fs.friendStore.UNFOLLOW(uuid, friendUUID, optionalMsg)
}

//ACCEPT accepts user
func (fs *FriendService) ACCEPT(currUUID string, friendUUID string) (*models.Friends, error) {
	return fs.friendStore.ACCEPT(currUUID, friendUUID)
}
