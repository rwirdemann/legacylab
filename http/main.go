package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rwirdemann/legacylab/complexity"
	"github.com/rwirdemann/legacylab/git"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/frequency", MakeChangeFrequencyHandler()).Methods("POST")
	log.Printf("Service listening on http://localhost:8090...")
	handler := cors.AllowAll().Handler(r)
	_ = http.ListenAndServe(":8090", handler)
}

func MakeChangeFrequencyHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		b, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Printf("error: %s", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var body struct {
			URL string `json:"url"`
		}
		if err := json.Unmarshal(b, &body); err != nil {
			log.Printf("error: %s", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var response struct {
			Files []git.File `json:"files"`
		}

		path := git.Checkout(body.URL)
		files := git.ChangeFrequency(body.URL, 30)
		for _, f := range files {
			var found bool
			f.Lines, f.Complextiy, found = complexity.Calc(fmt.Sprintf("%s/%s", path, f.Name))
			if found {
				response.Files = append(response.Files, f)
			}
		}

		data, err := json.Marshal(response)
		if err != nil {
			log.Printf("error: %s", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(data)
		if err != nil {
			log.Printf("error: %s", err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
