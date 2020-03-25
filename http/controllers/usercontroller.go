package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Rahul12344/Recipes/models"
	"github.com/gorilla/mux"
)

// UserService contains authorization functionality
type UserService interface {
	GET(username string, password string) (map[string]interface{}, string, string, time.Time, time.Time)
	SET(email string, uuid string, userModded *models.User) (map[string]interface{}, error)
	PUT(email string, password string, firstName string, lastName string) (bool, error)
	DEL(username string, password string) (bool, error)
	REFRESH(uuid string) (map[string]interface{}, string, time.Time)
}

// UserController controls user actions
type UserController struct {
	User UserService
}

//NewUserController creates new user controller
func NewUserController(userService UserService) *UserController {
	return &UserController{
		User: userService,
	}
}

// Setup sets up handlers
func (uc *UserController) Setup(r *mux.Router) {
	r.HandleFunc("/login", uc.Login).Methods("POST")
	r.HandleFunc("/signup", uc.Signup).Methods("POST")
}

// VerifiedSetup sets up handlers
func (uc *UserController) VerifiedSetup(r *mux.Router) {
	r.HandleFunc("/login", uc.Login).Methods("POST")
	r.HandleFunc("/signup", uc.Signup).Methods("POST")
}

// Login login users and provides authentication token for user
func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var loginInfo models.User
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&loginInfo)
	resp, token, _, _, _ := uc.User.GET(loginInfo.Username, loginInfo.Password)
	if token != "" {
		json.NewEncoder(w).Encode(resp)
	}
}

// Signup signs up users and provides auth token
func (uc *UserController) Signup(w http.ResponseWriter, r *http.Request) {
	var userInfo models.User
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&userInfo)
	var resp = map[string]interface{}{"status": false, "user": userInfo}

	status, _ := uc.User.PUT(userInfo.Username, userInfo.Password, userInfo.FirstName, userInfo.LastName)
	if status != true {
		http.Error(w, "Error signing up", 500)
		return
	}
	resp, token, _, _, _ := uc.User.GET(userInfo.Username, userInfo.Password)
	if token != "" {
		json.NewEncoder(w).Encode(resp)
	}
}
