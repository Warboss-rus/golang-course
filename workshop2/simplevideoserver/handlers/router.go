package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router(vc VideosRepository, fs FilesHandler) http.Handler {
	WithConnector := func(f func(http.ResponseWriter, *http.Request, VideosRepository)) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			f(w, r, vc)
		}
	}

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/list", WithConnector(handleList)).Methods(http.MethodGet)
	s.HandleFunc("/video/{ID}", WithConnector(handleVideo)).Methods(http.MethodGet)
	s.HandleFunc("/video/{ID}/status", WithConnector(handleStatus)).Methods(http.MethodGet)
	s.HandleFunc("/video", func(writer http.ResponseWriter, request *http.Request) {
		handleVideoUpload(writer, request, vc, fs)
	}).Methods(http.MethodPost)
	return logHttpHandler(r)
}
