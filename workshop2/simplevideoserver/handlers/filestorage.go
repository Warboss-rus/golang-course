package handlers

import "io"

type FileStorage interface {
	// Returns url to access the file
	StoreFile(filename string, content io.Reader) (string, error)
}
