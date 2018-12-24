package handlers

// Status of the video
type Status int

const (
	// Created means the video was uploaded but not processed yet
	Created Status = 1
	// Processing means the video is being processed right now
	Processing Status = 2
	// Ready means the video is fully processed and ready to be watched
	Ready Status = 3
	// Deleted means the video was deleted
	Deleted Status = 4
	// Error means there was an error when processing the video
	Error Status = 5
)

// Video is a struct what describes a video
type Video struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
	URL       string `json:"url"`
	Status    Status `json:"status"`
}
