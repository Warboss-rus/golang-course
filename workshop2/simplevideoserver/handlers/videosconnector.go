package handlers

type VideosConnector interface {
	GetVideoList() ([]Video, error)
	GetVideoDetails(videoId string) (Video, error)
	AddVideo(video Video) error
}
