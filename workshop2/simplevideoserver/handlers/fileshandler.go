package handlers

import "io"

type FilesHandler interface {
	// Returns url to access the file
	CreateFile(id string, filename string, content io.Reader) (string, error)
}
