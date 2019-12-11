package api

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"

	"encoding/csv"
	"encoding/json"
	"os"

	"github.com/gorilla/mux"
)

type ApiProduct struct {
	Title    string
	DocUrl   string
	Category string
}

var jsonData []byte

func LoadMeta(router *mux.Router) {

	csvFile, err := os.Open("./api/products.csv")
	if err != nil {
		log.Error(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Error(err)
	}

	var api ApiProduct
	var allApis []ApiProduct

	for _, each := range csvData {
		api.Title = each[1]
		api.DocUrl = each[2]
		api.Category = each[4]
		allApis = append(allApis, api)
	}

	// Convert to JSON
	jsonData, err = json.Marshal(allApis)
	if err != nil {
		log.Error(err)
	}

	log.Info(string(jsonData))

	router.HandleFunc("/api", ApiList)
	router.HandleFunc("/api/enable/{apiId}", EnableApi)
	router.HandleFunc("/api/disable/{apiId}", DisableApi)
}

func ApiList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", string(jsonData))
}

func EnableApi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apiId := vars["apiId"]

	log.WithField("apiId", apiId).Info("Enabling Cloud API...")
}

func DisableApi(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apiId := vars["apiId"]

	log.WithField("apiId", apiId).Info("Disabling Cloud API...")
}

