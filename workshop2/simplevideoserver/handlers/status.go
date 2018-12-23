package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type statusStruct struct {
	Status Status `json:"status"`
}

func handleStatus(w http.ResponseWriter, r *http.Request, repository VideosRepository) {
	vars := mux.Vars(r)
	id := vars["ID"]
	status, err := repository.GetVideoStatus(id)
	if err != nil {
		if _, ok := err.(*VideoNotFound); !ok {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	sStruct := statusStruct{status}
	jsonContent, err := json.Marshal(sStruct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(jsonContent)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}
