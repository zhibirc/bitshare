// Package main is the root application entry point.
// It starts the server and binds routes with corresponding controllers.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/zhibirc/bitshare/controllers"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := fmt.Sprintf(":%s", os.Getenv("TCP_PORT"))

	http.HandleFunc("/", controllers.RouteMain)
	http.HandleFunc("/health", controllers.RouteHealth)

	fmt.Printf("Server is listening on port%s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error running HTTP server: %s\n", err)
	}
}
