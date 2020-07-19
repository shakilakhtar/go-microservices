package main

import (
	"context"
	"shakilakhtar/go-microservices/assets"
	"shakilakhtar/go-microservices/utils"
	"net/http"
	"os"
)

const (
	DEFAULT_PORT = "8080"
)

func initialize() {
	//load application resources
	utils.Init()
}

func main() {

	port := os.Getenv("PORT")

	if len(port) == 0 {

		port = DEFAULT_PORT

	}

	initialize()

	var service assets.Service
	service = assets.NewService()

	ctx := context.Background()
	//gets a router
	router := assets.NewRouter(ctx, service)
	//start serving http requests
	http.ListenAndServe(":"+port, router)

}
