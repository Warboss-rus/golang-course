package handlers

type Status int

const (
	Created    Status = 1
	Processing Status = 2
	Ready      Status = 3
	Deleted    Status = 4
	Error      Status = 5
)

type Video struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
	Url       string `json:"url"`
	Status    Status `json:"status"`
}
