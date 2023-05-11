package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
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

func processRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		log.Printf("expected GET request, but got %s\n", req.Method)
		return
	}

	requestUri, err := url.Parse(req.RequestURI)

	if err != nil {
		log.Println(err)
		return
	}

	query := requestUri.RawQuery

	if query == "" {
		log.Println("WARNING: query string is empty")
		return
	}

	keyValueMap, _ := url.ParseQuery(query)

	var ttl int
	srcKey, isSrcKeyExists := keyValueMap["src"]
	ttlKey, isTtlKeyExists := keyValueMap["ttl"]

	if !isSrcKeyExists {
		log.Println("WARNING: required \"src\" query parameter is absent")
		return
	}

	if isTtlKeyExists {
		value, err := strconv.Atoi(ttlKey[0])

		if err != nil {
			log.Println("WARNING: \"ttl\" query parameter has invalid format, integer expected")
			return
		}

		ttl = value
	}

	srcValue := srcKey[0]

	if len(srcValue) == 0 {
		log.Println("WARNING: \"src\" field is empty")
		return
	}

	_, err = url.ParseRequestURI(srcValue)

	if err == nil {
		id := generateId()
		err := dbClient.Set(ctx, id, srcValue, time.Duration(ttl)).Err()

		if err != nil {
			log.Println("error occurred while set ID:URL record")
			return
		}

		data, err := json.Marshal(ResponseId{id})

		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		res.Write(data)
	} else {
		uri, err := dbClient.Get(ctx, srcValue).Result()

		if err != nil {
			log.Println("WARNING: any URI not found by given ID")
			uri = ""
		}

		data, err := json.Marshal(ResponseUri{uri})

		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		res.Write(data)
	}
}

func generateId() string {
	return "123"
}
