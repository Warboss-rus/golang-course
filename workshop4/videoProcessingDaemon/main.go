package main

import (
	"github.com/Warboss-rus/golang-course/workshop2/simplevideoserver/database"
	"github.com/Warboss-rus/golang-course/workshop4/videoProcessingDaemon/ffmpeg"
	"github.com/Warboss-rus/golang-course/workshop4/videoProcessingDaemon/videoprocessing"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
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

	var db database.DataBaseVideoRepository
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	killSignalChan := getKillSignalChan()
	stopChan := make(chan struct{})
	processor := ffmpeg.FfmpegVideoProcessor{}
	videosChan := videoprocessing.WatchVideos(stopChan, &db)
	wg := videoprocessing.RunWorkerPool(videosChan, &db, &processor)

	waitForKillSignal(killSignalChan)
	stopChan <- struct{}{}
	wg.Wait()
}
