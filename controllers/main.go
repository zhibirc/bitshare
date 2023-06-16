// Package controllers contains route handlers.
package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/zhibirc/bitshare/services"
	"github.com/zhibirc/bitshare/tools"
)

// TODO: Move to env vars.
const DB_ENGINE_REDIS string = "REDIS"

var dbClient = services.GetConnection(DB_ENGINE_REDIS)

type responseUri struct {
	uri string
}

type responseId struct {
	id string
}

func RouteMain(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		log.Printf("expected GET request, but got %s\n", req.Method)
		return
	}

	ctx := context.Background()
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
		id := tools.GenerateId()
		err := dbClient.Set(ctx, id, srcValue, time.Duration(ttl)).Err()

		if err != nil {
			log.Println("error occurred while set ID:URL record")
			return
		}

		data, err := json.Marshal(responseId{id})

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

		data, err := json.Marshal(responseUri{uri})

		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		res.Write(data)
	}
}
