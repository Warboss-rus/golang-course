package storage

import (
	"io"
	"os"
	"path/filepath"
)

// FileSystemStorage stores its files in a specified folder
type FileSystemStorage struct {
	contentPath string
}

// StoreFile Saves the file inside content folder and returns an URL that can be used to access that file
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

// NewFileSystemStorage creates a new storage on specified folder
func NewFileSystemStorage(contentPath string) *FileSystemStorage {
	return &FileSystemStorage{contentPath}
}
