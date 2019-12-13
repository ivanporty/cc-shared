package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"google.com/cc-shared/sql"
	"net/http"

	api "google.com/cc-shared/api"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)

	api.InitApis(r)
	sql.InitCloudSql(r)

	log.Info("Starting Cloud Code Shared Server")
	log.Fatal(http.ListenAndServe(":1205", nil))
}
