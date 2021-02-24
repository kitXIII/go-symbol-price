package server

import (
	"log"
	"net/http"
)

type storage interface {
	GetJsonData() ([]byte, error)
}

// RunServer starts the prices server
func RunServer(s storage) {
	handler := getHandler(s)

	http.HandleFunc("/", handler)
	log.Print(http.ListenAndServe(":8080", nil))
}

// getHandler return http request handler
func getHandler(s storage) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		body, err := s.GetJsonData()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		_, err = w.Write(body)

		if err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}
