package recipes

import (
	"context"
	"log"
	"net/http"

	"github.com/Rahul12344/skelego/services/index"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/Rahul12344/Recipes/http/controllers"
	"github.com/Rahul12344/Recipes/middleware/authhandlers"
	"github.com/Rahul12344/Recipes/services"
	"github.com/Rahul12344/Recipes/storage/postgres"
	"github.com/Rahul12344/Recipes/storage/postgres/schemas"
	"github.com/Rahul12344/skelego"
	"github.com/Rahul12344/skelego/services/storage/sqlservice"
)

//ExecuteServer start server
func ExecuteServer(config Config) {
	logger := skelego.NewLogger()
	conf := skelego.NewConfig("config", logger, ".", "..")

	pDB := sqlservice.NewORM()
	pDB.Configurifier(conf)
	es := index.NewIndex()
	es.Configurifier(conf)

	pDB.Connect(context.Background(), conf, logger)
	es.Connect(context.Background(), conf, logger)
	pDB.Start(context.Background(), logger)

	postgres.Migrate(logger, schemas.NewUserSchema(pDB), schemas.NewRecipeSchema(pDB), schemas.NewUserRecipesSchema(pDB),
		schemas.NewItemSchema(pDB), schemas.NewIngredientsSchema(pDB), schemas.NewInventorySchema(pDB),
		schemas.NewInstructionsSchema(pDB), schemas.NewTagsSchema(pDB), schemas.NewQuantitiesSchema(pDB),
		schemas.NewRecipeIngredientsSchema(pDB), schemas.NewDeliverySchema(pDB), schemas.NewChatRoomSchema(pDB))

	defer pDB.Stop(context.Background(), logger)

	userStore := postgres.NewUserStore(pDB)
	userService = services.NewUserService(userStore)

	/*
		userStore := postgres.NewUserStore(postgresDB)
		recipeStore := postgres.NewRecipeStore(postgresDB)
		userRecipeStore := postgres.NewUserRecipeStore(postgresDB)
		friendStore := postgres.NewFriendStore(postgresDB)
		deliveryStore := postgres.NewDeliveryStore(postgresDB)
		itemStore := postgres.NewItemStore(postgresDB)

		postgres.CreatePostgresTables(userStore, userRecipeStore, friendStore, recipeStore, deliveryStore, itemStore)
		logger.LogEvent("Migrated models to database.")

		userService := services.NewUserService(userStore, userRecipeStore)
		friendService := services.NewFriendService(friendStore)
		recipeService := services.NewRecipeService(recipeStore)

		logger.LogEvent("Starting internal services.")

		uc := controllers.NewUserController(userService)
		fc := controllers.NewFriendController(friendService)
		rc := controllers.NewRecipeController(recipeService)
		ic := controllers.NewImageController()
	*/
	cc := controllers.NewChatController()

	logger.LogEvent("Chat rooms opened on web sockets")
	go cc.Rooms.Run()

	r, _ := setupHandlers()
	/*uc.Setup(r)
	rc.Setup(r)
	uc.VerifiedSetup(s)
	fc.Setup(s)
	ic.Setup(s)*/

	logger.LogEvent("Starting controllers.")

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
