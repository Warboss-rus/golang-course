package main

import (
	"context"
	"flag"
	"github.com/Warboss-rus/golang-course/workshop2/simplevideoserver/database"
	"github.com/Warboss-rus/golang-course/workshop2/simplevideoserver/handlers"
	"github.com/Warboss-rus/golang-course/workshop2/simplevideoserver/storage"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func startServer(serverURL string, videosRepository handlers.VideosRepository, fs handlers.FileStorage) *http.Server {
	router := handlers.Router(videosRepository, fs)
	srv := &http.Server{Addr: serverURL, Handler: router}
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

	const serverURL = ":8000"
	log.WithFields(log.Fields{"url": serverURL}).Info("starting the server")

	const defaultContentDir = "workshop2\\simplevideoserver\\content"
	workDir, _ := os.Getwd()
	contentDir := flag.String("-dir", filepath.Join(workDir, defaultContentDir), "Specify a directory to store the videos")
	user := flag.String("-user", "root", "Specify a user for database access")
	password := flag.String("-password", "root", "Specify a password for database access")
	dbname := flag.String("-database", "db1", "Specify a database name for database access")
	flag.Parse()

	var db database.DBVideoRepository
	if err := db.Connect(*dbname, *user, *password); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fs := storage.NewFileSystemStorage(*contentDir)

	killSignalChan := getKillSignalChan()
	srv := startServer(serverURL, &db, fs)

	waitForKillSignal(killSignalChan)
	log.Fatal(srv.Shutdown(context.Background()))
}
