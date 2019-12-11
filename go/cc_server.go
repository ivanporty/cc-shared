package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"

	api "google.com/cc-shared/api"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)

	api.LoadMeta(r)

	log.Info("Starting Cloud Code Shared Server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
