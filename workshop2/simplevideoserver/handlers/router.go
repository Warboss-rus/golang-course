package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router(vr VideosRepository, fs FilesHandler) http.Handler {
	WithRepository := func(f func(http.ResponseWriter, *http.Request, VideosRepository)) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			f(w, r, vr)
		}
	}

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/list", WithRepository(handleList)).Methods(http.MethodGet)
	s.HandleFunc("/video/{ID}", WithRepository(handleVideo)).Methods(http.MethodGet)
	s.HandleFunc("/video/{ID}/status", WithRepository(handleStatus)).Methods(http.MethodGet)
	s.HandleFunc("/video", func(writer http.ResponseWriter, request *http.Request) {
		handleVideoUpload(writer, request, vr, fs)
	}).Methods(http.MethodPost)
	return logHttpHandler(r)
}
