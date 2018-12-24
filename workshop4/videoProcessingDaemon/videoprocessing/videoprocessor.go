package videoprocessing

// VideoProcessor is an interface for video processing routines
type VideoProcessor interface {
	GetVideoDuration(videoPath string) (float64, error)
	CreateVideoThumbnail(videoPath string, thumbnailPath string, thumbnailOffset int64) error
}
