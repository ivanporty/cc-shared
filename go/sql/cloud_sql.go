package sql

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func InitCloudSql(router *mux.Router) {
	router.HandleFunc("/cloudsql/instances", ListInstances)
}

func ListInstances(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	cmd := exec.Command("/usr/local/google/home/ivanporty/google-cloud-sdk/bin/gcloud", "sql", "instances", "list", "--format", "json")
	completeOutput, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("Error %s", err)
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		log.Info("Received SQL instances: %s", string(completeOutput))
		fmt.Fprintf(w, "%s\n", completeOutput)
	}
}
