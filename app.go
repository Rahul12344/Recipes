package recipes

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/Rahul12344/Recipes/http/controllers"
	"github.com/Rahul12344/Recipes/middleware/authhandlers"
	"github.com/Rahul12344/Recipes/services"
	"github.com/Rahul12344/Recipes/storage/postgres"
	"github.com/jinzhu/gorm"
)

//ExecuteServer start server
func ExecuteServer(config Config) {
	var userDB *gorm.DB
	var friendDB *gorm.DB
	var recipeDB *gorm.DB
	_, ok := os.LookupEnv("DATABASE_URL")
	if ok {
		userDB = postgres.Connect(true, "DATABASE_URL", config.Storage.UserHost,
			config.Storage.UserUsername,
			config.Storage.Username,
			config.Storage.UserPassword)
	}
	if !ok {
		userDB = postgres.Connect(false, "", config.Storage.UserHost,
			config.Storage.UserUsername,
			config.Storage.Username,
			config.Storage.UserPassword)
		friendDB = postgres.Connect(false, "", config.Storage.FriendHost,
			config.Storage.FriendUsername,
			config.Storage.Friendname,
			config.Storage.FriendPassword)
		recipeDB = postgres.Connect(false, "", config.Storage.RecipeHost,
			config.Storage.RecipeUsername,
			config.Storage.RecipeName,
			config.Storage.RecipePassword)

	}

	userStore := postgres.NewUserStore(userDB)
	userRecipeStore := postgres.NewUserRecipeStore(userDB)
	friendStore := postgres.NewFriendStore(friendDB)
	recipeStore := postgres.NewRecipeStore(recipeDB)

	postgres.CreatePostgresTables(userStore, userRecipeStore, friendStore, recipeStore)

	userService := services.NewUserService(userStore, userRecipeStore)
	friendService := services.NewFriendService(friendStore)
	recipeService := services.NewRecipeService(recipeStore)

	uc := controllers.NewUserController(userService)
	fc := controllers.NewFriendController(friendService)
	rc := controllers.NewRecipeController(recipeService)
	ic := controllers.NewImageController()

	r, s := setupHandlers()
	uc.Setup(r)
	rc.Setup(r)
	uc.VerifiedSetup(s)
	fc.Setup(s)
	ic.Setup(s)

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Origin"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))
}

func setupHandlers() (*mux.Router, *mux.Router) {
	r := mux.NewRouter()
	http.Handle("/", r)
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(authhandlers.VerifyToken)

	return r, s
}
