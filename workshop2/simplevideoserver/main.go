package main

import (
	"github.com/Warboss-rus/golang-course/workshop2/simplevideoserver/handlers"
	"net/http"
)

func main() {
	router := handlers.Router()
	_ = http.ListenAndServe(":8000", router)
}
