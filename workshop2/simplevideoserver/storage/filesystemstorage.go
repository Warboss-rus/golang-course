package storage

import (
	"io"
	"os"
	"path/filepath"
)

type FileSystemStorage struct {
	contentPath string
}

func (fs *FileSystemStorage) StoreFile(filename string, content io.Reader) (string, error) {
	file, err := createFile(filename, fs.contentPath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, content)
	if err != nil {
		return "", err
	}
	return filepath.Join("content", filename), nil
}

func createFile(fileName string, dirPath string) (*os.File, error) {
	filePath := filepath.Join(dirPath, fileName)
	if err := os.Mkdir(filepath.Dir(filePath), os.ModeDir); err != nil {
		return nil, err
	}
	return os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
}

func NewFileSystemStorage(contentPath string) *FileSystemStorage {
	return &FileSystemStorage{contentPath}
}
