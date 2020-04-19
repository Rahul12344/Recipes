package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Rahul12344/Recipes/models"
	"github.com/Rahul12344/Recipes/util/curruser"
	"github.com/Rahul12344/Recipes/util/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// UserService contains authorization functionality
type UserService interface {
	GET(username string, password string) (map[string]interface{}, string, string, time.Time, time.Time)
	SET(email string, uuid string, userModded *models.RecipeUser) (map[string]interface{}, error)
	PUT(user *models.RecipeUser) (bool, error)
	DEL(username string, password string) (bool, error)
	REFRESH(uuid string) (map[string]interface{}, string, time.Time)
	ADD(uuid string, recipe *models.Recipe)
	REMOVE(uuid string, recipe *models.Recipe)
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
	resp, token, _, _, _ := uc.User.GET(loginInfo.Username, loginInfo.Password)
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

	pass, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	userInfo.Password = string(pass)
	userInfo.UserID = uuid.UUID()

	status, _ := uc.User.PUT(&userInfo)
	if status != true {
		http.Error(w, "Error signing up", 500)
		return
	}
	resp, token, _, _, _ := uc.User.GET(userInfo.Username, userInfo.Password)
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

	uc.User.ADD(uuid, &recipe)
}

// RemoveRecipe removes recipe
func (uc *UserController) RemoveRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&recipe)

	uuid := curruser.GetCurrUser(w, r)

	uc.User.REMOVE(uuid, &recipe)
}
