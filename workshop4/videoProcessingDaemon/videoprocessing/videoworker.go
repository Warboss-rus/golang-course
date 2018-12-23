package videoprocessing

import (
	log "github.com/sirupsen/logrus"
	"path/filepath"
)

func RunWorker(videosChan <-chan Video, repository VideoRepository, processor VideoProcessor, contentPath string, name int) {
	log.Printf("start worker %v\n", name)
	for video := range videosChan {
		log.Printf("start processing video %s on %v\n", video.Id, name)
		if err := repository.UpdateVideoStatus(video.Id, Processing); err != nil {
			log.Printf("Cannot update status of video %s", video.Id)
		}

		pathToVideoFile := filepath.Join(contentPath, video.Url)
		duration, err := processor.GetVideoDuration(pathToVideoFile)
		if err != nil {
			log.Printf("Failed to calculate video duration of %s", video.Url)
			if err = repository.UpdateVideoStatus(video.Id, Error); err != nil {
				log.Printf("Cannot update status of video %s", video.Id)
			}
			continue
		}
		thumbnailUrl := filepath.Join("content", video.Id, "screen.jpg")
		thumbnailPath := filepath.Join(contentPath, thumbnailUrl)
		var thumbnailOffset int64
		if err = processor.CreateVideoThumbnail(pathToVideoFile, thumbnailPath, thumbnailOffset); err != nil {
			log.Printf("Failed to generate thumbnail for %s", video.Url)
			if err = repository.UpdateVideoStatus(video.Id, Error); err != nil {
				log.Printf("Cannot update status of video %s", video.Id)
			}
			continue
		}
		if err = repository.UpdateVideo(video.Id, int(duration), thumbnailUrl, Ready); err != nil {
			log.Printf("Failed to update video %s", video.Id)
		}
		log.Printf("end processing video %s on %v\n", video.Id, name)
	}
}
