package handlers

type VideosRepository interface {
	GetVideoList(search string, start *uint, count *uint) ([]Video, error)
	GetVideoDetails(videoId string) (Video, error)
	AddVideo(video Video) error
	GetVideoStatus(videoId string) (Status, error)
}
