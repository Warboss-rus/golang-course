package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func handleVideo(w http.ResponseWriter, r *http.Request, repository VideosRepository) {
	vars := mux.Vars(r)
	id := vars["ID"]
	video, err := repository.GetVideoDetails(id)
	if err != nil {
		if _, ok := err.(*VideoNotFound); !ok {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	jsonContent, err := json.Marshal(video)
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
