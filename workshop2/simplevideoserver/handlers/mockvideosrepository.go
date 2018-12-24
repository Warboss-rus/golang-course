package handlers

import (
	"errors"
	"strings"
)

// MockVideoRepository is an in-memory implementation of VideosRepository interface
type MockVideoRepository struct {
	videos        []Video
	errorToReturn error
}

// GetVideoList returns a list of videos satisfying the criteria
func (repository *MockVideoRepository) GetVideoList(search string, start *uint, count *uint) ([]Video, error) {
	videos := repository.videos
	if len(search) != 0 {
		filtered := make([]Video, 0)
		for _, v := range videos {
			if strings.Contains(v.Name, search) {
				filtered = append(filtered, v)
			}
		}
		videos = filtered
	}
	if start != nil {
		videos = videos[*start:]
	}
	if count != nil {
		videos = videos[:*count]
	}
	return videos, repository.errorToReturn
}

// GetVideoDetails returns the details of single video
func (repository *MockVideoRepository) GetVideoDetails(videoID string) (Video, error) {
	for _, v := range repository.videos {
		if v.ID == videoID {
			return v, repository.errorToReturn
		}
	}
	return Video{}, errors.New("Invalid video requested. ID=" + videoID)
}

// AddVideo adds a new video to the list
func (repository *MockVideoRepository) AddVideo(video Video) error {
	repository.videos = append(repository.videos, video)
	return repository.errorToReturn
}

// GetVideoStatus returns the status of the video
func (repository *MockVideoRepository) GetVideoStatus(videoID string) (Status, error) {
	for _, v := range repository.videos {
		if v.ID == videoID {
			return v.Status, repository.errorToReturn
		}
	}
	return Error, errors.New("Invalid video requested. ID=" + videoID)
}

// NewMockVideoRepository creates a new MockVideoRepository and fills it with initial data
func NewMockVideoRepository() MockVideoRepository {
	var repository MockVideoRepository
	repository.videos = []Video{
		{
			"d290f1ee-6c54-4b01-90e6-d701748f0851",
			"Black Retrospetive Woman",
			15,
			"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
			"/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4",
			Created,
		},
		{
			"sldjfl34-dfgj-523k-jk34-5jk3j45klj34",
			"Go Rally TEASER-HD",
			41,
			"/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/screen.jpg",
			"/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/index.mp4",
			Processing,
		},
		{
			"hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345",
			"Танцор",
			92,
			"/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/screen.jpg",
			"/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/index.mp4",
			Ready,
		},
	}
	return repository
}
