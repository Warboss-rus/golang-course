package handlers

import (
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
)

func handleVideoUpload(w http.ResponseWriter, r *http.Request, db VideosRepository, fs FilesHandler) {
	fileReader, header, err := r.FormFile("file[]")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contentType := header.Header.Get("Content-Type")
	if contentType != "video/mp4" {
		http.Error(w, "Invalid file format", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	fileName := header.Filename
	url, err := fs.CreateFile(id, fileName, fileReader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v := Video{id, fileName, 123, filepath.Join("content", id, "screen.jpg"), url, Ready}
	err = db.AddVideo(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
