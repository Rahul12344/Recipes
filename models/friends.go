package models

import "github.com/jinzhu/gorm"

// Friends friend model
type Friends struct {
	gorm.Model
	UserID               string `json:"user_id"`
	FriendID             string `json:"friend_id"`
	FriendRequestMessage string `json:"friend_request_message"`
	TimeStamp            int64  `json:"timestamp"`
	Status               int    `json:"status"`
	TimeAccepted         int64  `json:"time_accepted"`
}

//TableName name of table
func (Friends) TableName() string {
	return "friends"
}
