package videoprocessing

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func WatchVideos(stopChan <-chan struct{}, repository VideoRepository, interval time.Duration) <-chan Video {
	videosChan := make(chan Video)
	go func() {
		for {
			select {
			case <-stopChan:
				close(videosChan)
				return
			default:
			}
			videos, err := repository.GetVideosByStatus(Created)
			if err != nil {
				log.Printf("Failed to fetch videos from repository: %s", err.Error())
			}
			if len(videos) != 0 && err == nil {
				for _, v := range videos {
					log.Printf("Got the video to process %s\n", v.Id)
					videosChan <- v
				}
			} else {
				time.Sleep(interval)
			}
		}
	}()
	return videosChan

}
