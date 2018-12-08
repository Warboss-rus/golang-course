package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func handleList(w http.ResponseWriter, _ *http.Request) {
	jsonContent, err := json.Marshal(videos)
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
