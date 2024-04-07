package main

import (
	"context"
	"encoding/gob"
	"log"

	mongodbstore "github.com/2-72/gorilla-sessions-mongodb"
	handlerServices "github.com/Sri2103/services/internal/ui/allServices"
	"github.com/Sri2103/services/internal/ui/config"
	"github.com/Sri2103/services/internal/ui/handlers"
	admin_handlers "github.com/Sri2103/services/internal/ui/handlers/admin"
	user_handlers "github.com/Sri2103/services/internal/ui/handlers/user"
	user_service "github.com/Sri2103/services/internal/ui/services/user"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	server := echo.New()
	initateGobModels()
	// configuration
	appConfig, err := config.GetAppConfig()
	if err != nil {
		log.Fatalf("Error while getting app configurations: %s", err)
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Printf("error connecting to mongodb %v \n", err)
		return
	}
	defer client.Disconnect(context.TODO())
	coll := client.Database("app_db").Collection("sessions")
	store, err := mongodbstore.NewMongoDBStore(coll, []byte("secret"))
	if err != nil {
		log.Printf("error initializing mongodb store %v \n", err)
		return
	}

	server.Use(middleware.Logger())
	// server.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	server.Use(session.Middleware(store))

	server.Static("/dist", "./dist")
	server.Static("/static", "./static")

	server.Use(
		middleware.Recover(),
	)
	server.Use(user_handlers.AddUserContext)
	svc := handlerServices.New(appConfig)
	initiateProducts(svc, server)
	initiateUsers(svc, server)
	initiateAdmin(server)
	log.Fatal(server.Start(":1102"))
}

func initiateAdmin(e *echo.Echo) {
	adminHandler := admin_handlers.NewHandler()
	adminHandler.SetAdminRoutes(e)
}

func initiateUsers(s *handlerServices.Services, e *echo.Echo) {
	userHandler := user_handlers.NewHandler(s)
	userHandler.SetUserRoutes(e)
}

func initiateProducts(s *handlerServices.Services, server *echo.Echo) {
	productHandlers := handlers.NewProductHandlers(s)
	productHandlers.SetProductRoutes(server)
}

func initateGobModels() {
	gob.Register(user_service.User{})
}
