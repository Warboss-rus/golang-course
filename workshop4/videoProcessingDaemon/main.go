package main

import (
	"flag"
	"github.com/Warboss-rus/golang-course/workshop2/simplevideoserver/database"
	"github.com/Warboss-rus/golang-course/workshop4/videoProcessingDaemon/ffmpeg"
	"github.com/Warboss-rus/golang-course/workshop4/videoProcessingDaemon/videoprocessing"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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

	const defaultContentDir = "workshop2\\simplevideoserver"
	contentDir := flag.String("-dir", defaultContentDir, "Specify a directory to store the videos")
	user := flag.String("-user", "root", "Specify a user for database access")
	password := flag.String("-password", "root", "Specify a password for database access")
	dbname := flag.String("-database", "db1", "Specify a database name for database access")
	flag.Parse()

	var db database.DBVideoRepository
	if err := db.Connect(*dbname, *user, *password); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	killSignalChan := getKillSignalChan()
	stopChan := make(chan struct{})
	processor := ffmpeg.VideoProcessor{}
	videosChan := videoprocessing.WatchVideos(stopChan, &db, 1*time.Second)
	wg := videoprocessing.RunWorkerPool(videosChan, &db, &processor, *contentDir)

	waitForKillSignal(killSignalChan)
	stopChan <- struct{}{}
	wg.Wait()
}
