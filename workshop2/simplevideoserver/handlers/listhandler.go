package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func optionalIntParam(params url.Values, name string) *uint {
	if skip := params.Get(name); len(skip) != 0 {
		if val, err := strconv.Atoi(skip); err == nil {
			uval := uint(val)
			return &uval
		}
	}
	return nil
}

func handleList(w http.ResponseWriter, r *http.Request, db VideosRepository) {
	params := r.URL.Query()
	start := optionalIntParam(params, "skip")
	count := optionalIntParam(params, "limit")
	search := params.Get("searchString")

	videos, err := db.GetVideoList(search, start, count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
