package handlers

import (
	"errors"
	"strings"
)

type MockVideoRepository struct {
	videos        []Video
	errorToReturn error
}

func (repository *MockVideoRepository) GetVideoList(search string, start *uint, count *uint) ([]Video, error) {
	videos := repository.videos
	if len(search) != 0 {
		filtered := []Video{}
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

func (repository *MockVideoRepository) GetVideoDetails(videoId string) (Video, error) {
	for _, v := range repository.videos {
		if v.Id == videoId {
			return v, repository.errorToReturn
		}
	}
	return Video{}, errors.New("Invalid video requested. Id=" + videoId)
}

func (repository *MockVideoRepository) AddVideo(video Video) error {
	repository.videos = append(repository.videos, video)
	return repository.errorToReturn
}

func (repository *MockVideoRepository) GetVideoStatus(videoId string) (Status, error) {
	for _, v := range repository.videos {
		if v.Id == videoId {
			return v.Status, repository.errorToReturn
		}
	}
	return Error, errors.New("Invalid video requested. Id=" + videoId)
}

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
