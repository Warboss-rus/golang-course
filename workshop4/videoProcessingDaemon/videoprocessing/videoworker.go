package videoprocessing

import (
	log "github.com/sirupsen/logrus"
	"path/filepath"
)

// RunWorker processes videos that are pushed to videoChan channel, calculated their duration and creates a thumbnail image.
// Will stop when videosChan is closed
func RunWorker(videosChan <-chan Video, repository VideoRepository, processor VideoProcessor, contentPath string, name int) {
	log.Printf("start worker %v\n", name)
	for video := range videosChan {
		log.Printf("start processing video %s on %v\n", video.ID, name)
		if err := repository.UpdateVideoStatus(video.ID, Processing); err != nil {
			log.Printf("Cannot update status of video %s", video.ID)
		}

		pathToVideoFile := filepath.Join(contentPath, video.URL)
		duration, err := processor.GetVideoDuration(pathToVideoFile)
		if err != nil {
			log.Printf("Failed to calculate video duration of %s", video.URL)
			if err = repository.UpdateVideoStatus(video.ID, Error); err != nil {
				log.Printf("Cannot update status of video %s", video.ID)
			}
			continue
		}
		thumbnailURL := filepath.Join("content", video.ID, "screen.jpg")
		thumbnailPath := filepath.Join(contentPath, thumbnailURL)
		var thumbnailOffset int64
		if err = processor.CreateVideoThumbnail(pathToVideoFile, thumbnailPath, thumbnailOffset); err != nil {
			log.Printf("Failed to generate thumbnail for %s", video.URL)
			if err = repository.UpdateVideoStatus(video.ID, Error); err != nil {
				log.Printf("Cannot update status of video %s", video.ID)
			}
			continue
		}
		if err = repository.UpdateVideo(video.ID, int(duration), thumbnailURL, Ready); err != nil {
			log.Printf("Failed to update video %s", video.ID)
		}
		log.Printf("end processing video %s on %v\n", video.ID, name)
	}
}
