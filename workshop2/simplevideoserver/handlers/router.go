package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/list", handleList).Methods(http.MethodGet)
	s.HandleFunc("/video/{ID}", handleVideo).Methods(http.MethodGet)
	return logHttpHandler(r)
}
