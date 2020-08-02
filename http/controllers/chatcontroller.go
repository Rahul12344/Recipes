package controllers

import (
	"log"
	"net/http"

	"github.com/Rahul12344/Recipes/services/chats"
	"github.com/Rahul12344/Recipes/util/websockets"
	"github.com/gorilla/mux"
)

//ChatController controls chat
type ChatController struct {
	Rooms chats.Rooms
}

//NewChatController cc
func NewChatController() *ChatController {
	return &ChatController{
		Rooms: chats.NewRooms(),
	}
}

//Setup sets up controller to route
func (cc *ChatController) Setup(r *mux.Router) {
	r.HandleFunc("/chat", cc.HandleChat).Methods("GET")
}

//HandleChat handles chats
func (cc *ChatController) HandleChat(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("room")
	socket, err := websockets.Upgrade(w, r)
	if err != nil {
		log.Println(err.Error())
	}
	subscription := websockets.NewSubscription(socket, roomID)
	cc.Rooms.Register <- subscription
	go subscription.SocketWriter()
	subscription.SocketReader(cc.Rooms.Broadcast, cc.Rooms.Unregister)
}
