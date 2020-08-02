package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Rahul12344/Recipes/models"
	"github.com/gorilla/mux"
)

//FriendService friend service
type FriendService interface {
	Follow(uuid string, friendUUID string, optionalMsg string) (bool, error)
	Unfollow(uuid string, friendUUID string, optionalMsg string) (bool, error)
	Accept(currUUID string, friendUUID string) (*models.Friends, error)
}

// FriendController controls friend actions for users
type FriendController struct {
	Friends FriendService
}

//NewFriendController creates new friend controller
func NewFriendController(friendService FriendService) *FriendController {
	return &FriendController{
		Friends: friendService,
	}
}

//Setup sets up routers
func (fs *FriendController) Setup(r *mux.Router) {
	r.HandleFunc("/follow", fs.Follow).Methods("POST")
	r.HandleFunc("/unfollow", fs.Unfollow).Methods("POST")
	r.HandleFunc("/accept", fs.AcceptFollowRequest).Methods("POST")
}

//Follow follows
func (fs *FriendController) Follow(w http.ResponseWriter, r *http.Request) {
	friendRequest := models.Friends{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&friendRequest)
	fs.Friends.Follow(friendRequest.UserID, friendRequest.FriendID, friendRequest.FriendRequestMessage)
}

//AcceptFollowRequest follows
func (fs *FriendController) AcceptFollowRequest(w http.ResponseWriter, r *http.Request) {
	friendRequest := models.Friends{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&friendRequest)
	fs.Friends.Accept(friendRequest.UserID, friendRequest.FriendID)
}

//Unfollow unfollows
func (fs *FriendController) Unfollow(w http.ResponseWriter, r *http.Request) {
	friendRequest := models.Friends{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&friendRequest)
	fs.Friends.Unfollow(friendRequest.UserID, friendRequest.FriendID, friendRequest.FriendRequestMessage)
}
