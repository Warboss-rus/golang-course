package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router(vc VideosConnector) http.Handler {
	WithConnector := func(f func(http.ResponseWriter, *http.Request, VideosConnector)) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			f(w, r, vc)
		}
	}

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/list", WithConnector(handleList)).Methods(http.MethodGet)
	s.HandleFunc("/video/{ID}", WithConnector(handleVideo)).Methods(http.MethodGet)
	s.HandleFunc("/video", WithConnector(handleVideoUpload)).Methods(http.MethodPost)
	return logHttpHandler(r)
}
