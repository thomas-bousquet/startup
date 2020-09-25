package main

import (
	"context"
	"github.com/gorilla/mux"
	"startup/api"
	"startup/clients"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	mongoClient := clients.NewMongoClient()
	api.RegisterRoutes(router, mongoClient.Database("startup"))

	defer mongoClient.Disconnect(context.Background())

	log.Info("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
