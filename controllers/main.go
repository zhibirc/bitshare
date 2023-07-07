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

	"github.com/zhibirc/bitshare/models"
	"github.com/zhibirc/bitshare/tools"
)

type responseUri struct {
	uri string
}

type responseId struct {
	id string
}

func RouteMain(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		log.Printf("expected GET request, but got %s\n", request.Method)
		response.Header().Set("Allow", "GET")
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()
	requestUri, err := url.Parse(request.RequestURI)

	if err != nil {
		log.Println(err)
		http.Error(response, "Bad Request", http.StatusBadRequest)
		return
	}

	query := requestUri.RawQuery

	if query == "" {
		log.Println("WARNING: query string is empty")
		http.Error(response, "Bad Request", http.StatusBadRequest)
		return
	}

	keyValueMap, _ := url.ParseQuery(query)

	var ttl int
	srcKey, isSrcKeyExists := keyValueMap["src"]
	ttlKey, isTtlKeyExists := keyValueMap["ttl"]

	if !isSrcKeyExists {
		log.Println("WARNING: required \"src\" query parameter is absent")
		http.Error(response, "Bad Request", http.StatusBadRequest)
		return
	}

	if isTtlKeyExists {
		value, err := strconv.Atoi(ttlKey[0])

		if err != nil {
			log.Println("WARNING: \"ttl\" query parameter has invalid format, integer expected")
			http.Error(response, "Bad Request", http.StatusBadRequest)
			return
		}

		ttl = value
	}

	srcValue := srcKey[0]

	if len(srcValue) == 0 {
		log.Println("WARNING: \"src\" field is empty")
		http.Error(response, "Bad Request", http.StatusBadRequest)
		return
	}

	_, err = url.ParseRequestURI(srcValue)

	if err == nil {
		id := tools.GenerateId()
		err := models.Record.Create(ctx, id, srcValue, time.Duration(ttl))

		if err != nil {
			log.Println("error occurred while set ID:URL record")
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(responseId{id})

		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Write(data)
	} else {
		uri, err := models.Record.GetOne(ctx, srcValue).Result()

		if err != nil {
			//log.Println("WARNING: any URI not found by given ID")
			uri = ""
		}

		data, err := json.Marshal(responseUri{uri})

		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Write(data)
	}
}
