package handlers

// VideoNotFound is a custom error type that is return if video does not exists
type VideoNotFound struct{}

func (v *VideoNotFound) Error() string {
	return "Video not found"
}

// VideosRepository is a database interface for this package
type VideosRepository interface {
	GetVideoList(search string, start *uint, count *uint) ([]Video, error)
	GetVideoDetails(videoID string) (Video, error)
	AddVideo(video Video) error
	GetVideoStatus(videoID string) (Status, error)
}
