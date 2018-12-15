package main

import (
	"context"
	"github.com/Warboss-rus/golang-course/workshop2/simplevideoserver/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func startServer(serverUrl string, db *handlers.DataBaseConnector) *http.Server {
	var fs handlers.FileSystemHandler
	router := handlers.Router(db, &fs)
	srv := &http.Server{Addr: serverUrl, Handler: router}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	return srv
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
		defer file.Close()
	}

	const serverUrl = ":8000"
	log.WithFields(log.Fields{"url": serverUrl}).Info("starting the server")

	var db handlers.DataBaseConnector
	db.Connect()
	defer db.Close()

	killSignalChan := getKillSignalChan()
	srv := startServer(serverUrl, &db)

	waitForKillSignal(killSignalChan)
	log.Fatal(srv.Shutdown(context.Background()))
}
