package handlers

import (
	"io"
	"os"
	"path/filepath"
)

type FileSystemHandler struct {
}

func (fs *FileSystemHandler) CreateFile(id string, filename string, content io.Reader) (string, error) {
	const environmentDir = "C:\\Users\\Warboss-rus\\go\\src\\github.com\\Warboss-rus\\golang-course"
	const videoDir = "workshop2\\simplevideoserver\\content"
	dir := filepath.Join(environmentDir, videoDir, id)
	file, err := createFile(filename, dir)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, content)
	if err != nil {
		return "", err
	}
	return filepath.Join("content", id, filename), nil
}

func createFile(fileName string, dirPath string) (*os.File, error) {
	if err := os.Mkdir(dirPath, os.ModeDir); err != nil {
		return nil, err
	}
	filePath := filepath.Join(dirPath, fileName)
	return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
}
