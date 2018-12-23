package videoprocessing

type Status int

const (
	Created    Status = 1
	Processing Status = 2
	Ready      Status = 3
	Deleted    Status = 4
	Error      Status = 5
)

type Video struct {
	Id  string
	Url string
}
