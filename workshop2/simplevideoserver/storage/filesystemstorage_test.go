package storage

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFileSystemHandler(t *testing.T) {
	const filename = "test.mp4"
	const id = "test-id"
	videoDir, _ := ioutil.TempDir(os.TempDir(), "TestFileSystemHandler")
	filePath := filepath.Join(id, filename)
	expectedPath := filepath.Join(videoDir, filePath)
	expectedURL := filepath.Join("content", filePath)
	defer os.RemoveAll(videoDir)
	const content = "file content"
	fs := NewFileSystemStorage(videoDir)
	url, err := fs.StoreFile(filePath, strings.NewReader(content))
	if err != nil {
		t.Error("StoreFile failed")
	}
	if url != expectedURL {
		t.Errorf("Invalid url received. Expected %s, got %s", expectedURL, url)
	}
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Error("File does not exists")
	}
	fileContent, err := ioutil.ReadFile(expectedPath)
	if err != nil {
		t.Error("Cannot read file content")
	}
	if string(fileContent) != content {
		t.Error("File content does not match")
	}
}
