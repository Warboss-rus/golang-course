package handlers

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func handleVideoUpload(w http.ResponseWriter, r *http.Request, repository VideosRepository, fs FileStorage) {
	fileReader, header, err := r.FormFile("file[]")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	contentType := header.Header.Get("Content-Type")
	if contentType != "video/mp4" {
		http.Error(w, "Invalid file format", http.StatusBadRequest)
		log.Error(err)
		return
	}

	id := uuid.New().String()
	fileName := header.Filename
	url, err := fs.CreateFile(id, fileName, fileReader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	v := Video{id, fileName, 0, "", url, Created}
	err = repository.AddVideo(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}
}
