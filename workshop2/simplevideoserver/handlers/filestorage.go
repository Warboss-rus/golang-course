package handlers

import "io"

type FileStorage interface {
	// Returns url to access the file
	CreateFile(id string, filename string, content io.Reader) (string, error)
}
