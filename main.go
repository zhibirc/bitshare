// Package main is the root application entry point.
// It starts the server and binds routes with corresponding controllers.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"controllers/main"
)

const port string = ":9870"

var ctx = context.Background()

func main() {
	http.HandleFunc("/", RouteMain(ctx))

	fmt.Printf("server is listening on port%s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("error running HTTP server: %s\n", err)
	}
}