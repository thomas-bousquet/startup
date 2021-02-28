package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/thomas-bousquet/user-service/api"
	"github.com/thomas-bousquet/user-service/clients"
	"github.com/thomas-bousquet/user-service/utils/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
)

var (
	version string
)

func main() {
	customLogger := logger.NewLogger()
	customLogger.Infof("Starting application version %q", version)
	router := mux.NewRouter()
	mongoClient := clients.NewMongoClient(customLogger)

	api.RegisterRoutes(router, mongoClient, customLogger)

	defer disconnect(mongoClient, customLogger)

	//TODO: Make this configurable
	port := os.Getenv("APP_PORT")
	customLogger.Infof("Starting server on port %q", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
		customLogger.Fatal(err)
	}
}

func disconnect(client *mongo.Client, logger *logrus.Logger) {
	if err := client.Disconnect(context.Background()); err != nil {
		logger.Fatal(err)
	}
}
