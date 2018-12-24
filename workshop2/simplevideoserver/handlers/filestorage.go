package handlers

import "io"

// FileStorage is an interface that allows handlers to store data
type FileStorage interface {
	// Returns url to access the file
	StoreFile(filename string, content io.Reader) (string, error)
}
