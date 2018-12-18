package handlers

type VideoNotFound struct{}

func (v *VideoNotFound) Error() string {
	return "Video not found"
}

type VideosRepository interface {
	GetVideoList(search string, start *uint, count *uint) ([]Video, error)
	GetVideoDetails(videoId string) (Video, error)
	AddVideo(video Video) error
	GetVideoStatus(videoId string) (Status, error)
}
