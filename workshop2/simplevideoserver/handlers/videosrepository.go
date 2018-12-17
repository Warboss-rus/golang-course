package handlers

type VideosRepository interface {
	GetVideoList() ([]Video, error)
	GetVideoDetails(videoId string) (Video, error)
	AddVideo(video Video) error
	GetVideoStatus(videoId string) (Status, error)
}
