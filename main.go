package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/thomas-bousquet/startup/api"
	"github.com/thomas-bousquet/startup/clients"
	"github.com/thomas-bousquet/startup/utils/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	customLogger := logger.NewLogger()
	mongoClient := clients.NewMongoClient(customLogger)

	api.RegisterRoutes(router, mongoClient, customLogger)

	defer disconnect(mongoClient, customLogger)

	//TODO: Make this configurable
	port := "8080"
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
