package videoprocessing

// VideoRepository is a database interface for videos
type VideoRepository interface {
	GetVideosByStatus(status Status) ([]Video, error)
	UpdateVideoStatus(videoID string, status Status) error
	UpdateVideo(videoID string, duration int, thumbnail string, status Status) error
}
