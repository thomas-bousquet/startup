package main

import (
	"context"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/thomas-bousquet/startup/api"
	"github.com/thomas-bousquet/startup/clients"
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
