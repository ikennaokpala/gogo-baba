package main

import (
	"encoding/json"
	"fmt"
	"github.com/ikennaokpala/gogo-baba/models"
	"net/http"
)

const (
	Host    = "localhost"
	Port    = "1051"
	Address = Host + ":" + Port
)

func main() {
	http.HandleFunc("/", serveRequests)
	http.ListenAndServe(Address, nil)
}

func serveRequests(rw http.ResponseWriter, r *http.Request) {
	response, err := getResponse()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(rw, string(response))
}

func getResponse() ([]byte, error) {
	posts := make(map[string]int)
	posts["posts"] = 61
	posts["pages"] = 99
	posts["recipes"] = 81

	media := make(map[string]int)
	media["photos"] = 205
	media["videos"] = 32
	media["pdfs"] = 72

	blog := model.Blog{posts, media}
	site := model.Site{blog}

	return json.MarshalIndent(site, "", "  ")
}
