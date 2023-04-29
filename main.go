package main

import (
	"fmt"
	"net/http"
	"net/url"
)

const port string = ":9870"

func main() {
	http.HandleFunc("/services/uri-grain", getUserData)

	fmt.Printf("server is listening on port%s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("error running HTTP server: %s\n", err)
	}
}

type ResponseUri struct {
	uri string
}

type ResponseId struct {
	id string
}

func getUserData(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		fmt.Printf("expected GET request, but got %s\n", req.Method)
		return
	}

	requestUri, err := url.Parse(req.RequestURI)

	if err != nil {
		panic(err)
	}

	query := requestUri.RawQuery

	if query == "" {
		fmt.Println("expected query string, but got nothing")
		return
	}

	keyValueMap, _ := url.ParseQuery(query)
	source := keyValueMap["source"][0]

	_, err = url.ParseRequestURI(source)

	if err == nil {
		w.Write([]byte(generateId()))
	} else {
		// id := uri
		w.Write([]byte("your original URL"))
	}

	// response := ResponseUri{}
	// response := ResponseId{r.URL.Path}
	// data, err := json.Marshal(response)

	//if err != nil {
	//    http.Error(w, err.Error(), 400)
	//    return
	//}
}

func generateId() string {
	return "123"
}
