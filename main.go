package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"net/url"
)

const port string = ":9870"

var dbClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0, // default DB
})
var ctx = context.Background()

type ResponseUri struct {
	uri string
}

type ResponseId struct {
	id string
}

func main() {
	http.HandleFunc("/services/uri-grain", processRequest)

	fmt.Printf("server is listening on port%s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("error running HTTP server: %s\n", err)
	}
}

func processRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		log.Printf("expected GET request, but got %s\n", req.Method)
		return
	}

	requestUri, err := url.Parse(req.RequestURI)

	if err != nil {
		panic(err)
	}

	query := requestUri.RawQuery

	if query == "" {
		log.Println("WARNING: query string is empty")
		return
	}

	keyValueMap, _ := url.ParseQuery(query)
	source := keyValueMap["source"][0]

	_, err = url.ParseRequestURI(source)

	if err == nil {
		id := generateId()
		err := dbClient.Set(ctx, id, source, 0).Err()
		if err != nil {
			panic(err)
		}

		data, err := json.Marshal(ResponseId{id})
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(data)
	} else {
		uri, err := dbClient.Get(ctx, source).Result()
		if err != nil {
			panic(err)
		}

		data, err := json.Marshal(ResponseUri{uri})
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(data)
	}
}

func generateId() string {
	return "123"
}
