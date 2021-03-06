package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

// Router redirects the requests to their respective handlers
func Router(vr VideosRepository, fs FileStorage) http.Handler {
	WithRepository := func(f func(http.ResponseWriter, *http.Request, VideosRepository)) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			f(w, r, vr)
		}
	}

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/list", WithRepository(handleList)).Methods(http.MethodGet)
	s.HandleFunc("/video/{ID}", WithRepository(handleVideoDetails)).Methods(http.MethodGet)
	s.HandleFunc("/video/{ID}/status", WithRepository(handleStatus)).Methods(http.MethodGet)
	s.HandleFunc("/video", func(writer http.ResponseWriter, request *http.Request) {
		handleVideoUpload(writer, request, vr, fs)
	}).Methods(http.MethodPost)
	return logHTTPHandler(r)
}
