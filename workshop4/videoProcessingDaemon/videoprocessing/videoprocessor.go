package videoprocessing

type VideoProcessor interface {
	GetVideoDuration(videoPath string) (float64, error)
	CreateVideoThumbnail(videoPath string, thumbnailPath string, thumbnailOffset int64) error
}
