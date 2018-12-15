package handlers

import (
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func handleVideoUpload(w http.ResponseWriter, r *http.Request, db VideosConnector) {
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
	const environmentDir = "C:\\Users\\Warboss-rus\\go\\src\\github.com\\Warboss-rus\\golang-course"
	const videoDir = "workshop2\\simplevideoserver\\content"
	dir := filepath.Join(environmentDir, videoDir, id)
	fileName := header.Filename
	file, err := createFile(fileName, dir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	_, err = io.Copy(file, fileReader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v := Video{id, fileName, 123, filepath.Join(dir, "screen.jpg"), filepath.Join("content", id, fileName)}
	err = db.AddVideo(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createFile(fileName string, dirPath string) (*os.File, error) {
	if err := os.Mkdir(dirPath, os.ModeDir); err != nil {
		return nil, err
	}
	filePath := filepath.Join(dirPath, fileName)
	return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
}
