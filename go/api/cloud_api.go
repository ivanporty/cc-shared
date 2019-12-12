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

var apiProductList []byte

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

	// skip first line - header.
	for i := 1; i < len(csvData); i++ {
		api.Title = csvData[i][1]
		api.DocUrl = csvData[i][2]
		api.Category = csvData[i][4]
		allApis = append(allApis, api)
	}

	// Convert to JSON
	apiProductList, err = json.Marshal(allApis)
	if err != nil {
		log.Error(err)
	}

	log.WithField("loaded APIs", len(apiProductList)).Info(string(apiProductList))

	router.HandleFunc("/api/list", ApiList)
	router.HandleFunc("/api/enable/{apiId}", EnableApi)
	router.HandleFunc("/api/disable/{apiId}", DisableApi)
}

func ApiList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", string(apiProductList))

	log.Info("Requested API list, served successfully")
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

