package models

//ChatRoom chat room
type ChatRoom struct {
	UserID     string `json:"userid" gorm:"not null"`
	ChatRoomID string `json:"chatid" gorm:"not null"`
}

// Message message object
type Message struct {
	Message     []byte `json:"message" gorm:"not null"`
	MessageRoom string `gorm:"not null"`
	MessageID   string `gorm:"not null"`
}
