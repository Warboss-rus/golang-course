package handlers

import "io"

type FilesHandler interface {
	CreateFile(id string, filename string, content io.Reader) error
}
