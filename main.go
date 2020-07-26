package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thomas-bousquet/startup/api"
	"github.com/thomas-bousquet/startup/client"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	mongoDB := client.NewMongoDB()
	api.RegisterRoutes(router, mongoDB)


	fmt.Printf("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
