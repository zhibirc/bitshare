package uri_grain

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const port string = ":9870"

func main() {
	http.HandleFunc("/services/uri-grain", getUserData)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("error occurred on server start")
		return
	}

	fmt.Printf("server is listening on port%s", port)
}

type ResponseUrl struct {
	url string
}

type ResponseId struct {
	id string
}

func getUserData(w http.ResponseWriter, r *http.Request) {
	response := ResponseUrl{r.URL.Path}
	data, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Write(data)
}
