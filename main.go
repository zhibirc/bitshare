package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"controllers/main"
)

const port string = ":9870"

type ResponseUri struct {
	uri string
}

type ResponseId struct {
	id string
}

var ctx = context.Background()

func main() {
	http.HandleFunc("/", RouteMain(ctx))

	fmt.Printf("server is listening on port%s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("error running HTTP server: %s\n", err)
	}
}