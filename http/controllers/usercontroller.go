package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/Recipes/util/curruser"
	"github.com/Rahul12344/Recipes/util/hash"
	"github.com/Rahul12344/Recipes/util/uuid"
	"github.com/gorilla/mux"
)

// UserService contains authorization functionality
type UserService interface {
	GetUser(username string, password string) (map[string]interface{}, string, string, time.Time, time.Time)
	SetUser(email string, uuid string, userModded *models.RecipeUser) (map[string]interface{}, error)
	NewUser(user *models.RecipeUser) (bool, error)
	DeleteUser(username string, password string) (bool, error)
	RefreshToken(uuid string) (map[string]interface{}, string, time.Time)
	AddRecipe(uuid string, recipe *models.Recipe)
	RemoveRecipe(uuid string, recipe *models.Recipe)
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
	r.HandleFunc("/addrecipe", uc.AddRecipe).Methods("POST")
	r.HandleFunc("/removerecipe", uc.RemoveRecipe).Methods("POST")
}

// Login login users and provides authentication token for user
func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var loginInfo models.RecipeUser
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&loginInfo)
	resp, token, _, _, _ := uc.User.GetUser(loginInfo.Username, loginInfo.Password)
	if token != "" {
		json.NewEncoder(w).Encode(resp)
	}
}

// Signup signs up users and provides auth token
func (uc *UserController) Signup(w http.ResponseWriter, r *http.Request) {
	var userInfo models.RecipeUser
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&userInfo)
	var resp = map[string]interface{}{"status": false, "user": userInfo}

	pass := hash.Hash(userInfo.Password)
	userInfo.Password = string(pass)
	userInfo.UserID = uuid.UUID()

	status, _ := uc.User.NewUser(&userInfo)
	if status != true {
		http.Error(w, "Error signing up", 500)
		return
	}
	resp, token, _, _, _ := uc.User.GetUser(userInfo.Username, userInfo.Password)
	if token != "" {
		json.NewEncoder(w).Encode(resp)
	}
}

// AddRecipe adds recipe
func (uc *UserController) AddRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&recipe)

	uuid := curruser.GetCurrUser(w, r)

	uc.User.AddRecipe(uuid, &recipe)
}

// RemoveRecipe removes recipe
func (uc *UserController) RemoveRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&recipe)

	uuid := curruser.GetCurrUser(w, r)

	uc.User.RemoveRecipe(uuid, &recipe)
}
