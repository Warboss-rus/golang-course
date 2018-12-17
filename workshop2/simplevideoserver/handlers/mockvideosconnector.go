package handlers

import "errors"

type MockVideoConnector struct {
	videos        []Video
	errorToReturn error
}

func (connector *MockVideoConnector) GetVideoList() ([]Video, error) {
	return connector.videos, connector.errorToReturn
}

func (connector *MockVideoConnector) GetVideoDetails(videoId string) (Video, error) {
	for _, v := range connector.videos {
		if v.Id == videoId {
			return v, connector.errorToReturn
		}
	}
	return Video{}, errors.New("Invalid video requested. Id=" + videoId)
}

func (connector *MockVideoConnector) AddVideo(video Video) error {
	connector.videos = append(connector.videos, video)
	return connector.errorToReturn
}

func (connector *MockVideoConnector) GetVideoStatus(videoId string) (Status, error) {
	for _, v := range connector.videos {
		if v.Id == videoId {
			return v.Status, connector.errorToReturn
		}
	}
	return Error, errors.New("Invalid video requested. Id=" + videoId)
}

func NewMockVideosConnector() MockVideoConnector {
	var connector MockVideoConnector
	connector.videos = []Video{
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
	return connector
}
