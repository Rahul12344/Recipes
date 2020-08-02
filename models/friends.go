package models

import "github.com/jinzhu/gorm"

// Friends friend model
type Friends struct {
	gorm.Model
	UserID               string `json:"user_id" gorm:"not null"`
	FriendID             string `json:"friend_id" gorm:"not null"`
	FriendRequestMessage string `json:"friend_request_message" gorm:"not null"`
	TimeStamp            int64  `json:"timestamp" gorm:"not null"`
	Status               int    `json:"status" gorm:"not null"`
	TimeAccepted         int64  `json:"time_accepted" gorm:"not null"`
}

//TableName name of table
func (Friends) TableName() string {
	return "friends"
}
