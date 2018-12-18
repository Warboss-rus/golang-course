package handlers

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

func optionalIntParam(params url.Values, name string, max int) (*uint, error) {
	if skip := params.Get(name); len(skip) != 0 {
		if val, err := strconv.Atoi(skip); err == nil {
			if val < 0 || ((max >= 0) && (val > max)) {
				return nil, errors.New("bad input parameter")
			}
			uval := uint(val)
			return &uval, nil
		} else {
			return nil, err
		}
	}
	return nil, nil
}

func handleList(w http.ResponseWriter, r *http.Request, db VideosRepository) {
	params := r.URL.Query()
	start, err := optionalIntParam(params, "skip", -1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	count, err := optionalIntParam(params, "limit", 50)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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
