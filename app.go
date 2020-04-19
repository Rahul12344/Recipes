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
	var postgresDB *gorm.DB
	_, ok := os.LookupEnv("DATABASE_URL")
	if ok {
		postgresDB = postgres.Connect(true, "DATABASE_URL", config.Storage.Host,
			config.Storage.Name,
			config.Storage.Username,
			config.Storage.Password)
	}
	if !ok {
		postgresDB = postgres.Connect(false, "", config.Storage.Host,
			config.Storage.Name,
			config.Storage.Username,
			config.Storage.Password)
	}

	userStore := postgres.NewUserStore(postgresDB)
	recipeStore := postgres.NewRecipeStore(postgresDB)
	userRecipeStore := postgres.NewUserRecipeStore(postgresDB)
	friendStore := postgres.NewFriendStore(postgresDB)
	deliveryStore := postgres.NewDeliveryStore(postgresDB)

	postgres.CreatePostgresTables(userStore, userRecipeStore, friendStore, recipeStore, deliveryStore)

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
