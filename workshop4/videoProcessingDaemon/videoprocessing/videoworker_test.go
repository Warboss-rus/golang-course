package videoprocessing

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"
)

type mockVideoRepository struct {
	videos    []Video
	videoID   string
	status    Status
	duration  int
	thumbnail string
}

type mockVideoProcessor struct {
	duration                 float64
	videoPath                string
	thumbnailPath            string
	videoDurationError       error
	thumbnailGenerationError error
}

func (mock *mockVideoRepository) GetVideosByStatus(status Status) ([]Video, error) {
	videos := mock.videos
	mock.videos = nil
	return videos, nil
}
func (mock *mockVideoRepository) UpdateVideoStatus(videoID string, status Status) error {
	mock.videoID = videoID
	mock.status = status
	return nil
}
func (mock *mockVideoRepository) UpdateVideo(videoID string, duration int, thumbnail string, status Status) error {
	mock.videoID = videoID
	mock.duration = duration
	mock.thumbnail = thumbnail
	mock.status = status
	return nil
}

func (mock *mockVideoProcessor) GetVideoDuration(videoPath string) (float64, error) {
	mock.videoPath = videoPath
	return mock.duration, mock.videoDurationError
}
func (mock *mockVideoProcessor) CreateVideoThumbnail(videoPath string, thumbnailPath string, thumbnailOffset int64) error {
	mock.videoPath = videoPath
	mock.thumbnailPath = thumbnailPath
	return mock.thumbnailGenerationError
}

func TestRunWorker(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	videoChan := make(chan Video)
	var repository mockVideoRepository
	repository.status = Created
	var processor mockVideoProcessor
	go RunWorker(videoChan, &repository, &processor, "", 0)

	//check initial state
	if len(repository.videos) != 0 || repository.videoID != "" || repository.status != Created || repository.duration != 0 || repository.thumbnail != "" || processor.videoPath != "" || processor.duration != 0 || processor.thumbnailPath != "" {
		t.Error("Inital state has been modified")
	}

	// add video to process
	video := Video{"videoId", "videoUrl"}
	thumbnail := filepath.Join("content", video.Id, "screen.jpg")
	processor.duration = 987
	videoChan <- video
	if repository.videoID != video.Id || repository.duration != int(processor.duration) || repository.thumbnail != thumbnail || repository.status != Ready {
		t.Error("Repository was not updated correctly")
	}
	if processor.thumbnailPath != thumbnail || processor.videoPath != video.Url {
		t.Error("Incorrect arguments supplied to videoprocessor")
	}

	// video duration error
	video = Video{"videoId2", "videoUrl2"}
	processor.videoDurationError = errors.New("cannot get video duration")
	processor.thumbnailGenerationError = nil
	videoChan <- video
	time.Sleep(time.Second)
	if repository.videoID != video.Id || repository.status != Error {
		t.Error("Worker did not report an error to repository")
	}

	// thumbnailError
	video = Video{"videoId3", "videoUrl3"}
	processor.videoDurationError = nil
	processor.thumbnailGenerationError = errors.New("cannot generate thumbnail")
	videoChan <- video
	time.Sleep(time.Second)
	if repository.videoID != video.Id || repository.status != Error {
		t.Error("Worker did not report an error to repository")
	}
}
