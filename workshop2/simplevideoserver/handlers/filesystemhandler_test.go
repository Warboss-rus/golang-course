package handlers

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
	const videoDir = "..\\content"
	expectedPath := filepath.Join(videoDir, id, filename)
	expectedUrl := filepath.Join("content", id, filename)
	defer os.RemoveAll(filepath.Join(videoDir, id))
	const content = "file content"
	var fs FileSystemHandler
	url, err := fs.CreateFile(id, filename, strings.NewReader(content))
	if err != nil {
		t.Error("CreateFile failed")
	}
	if url != expectedUrl {
		t.Errorf("Invalid url received. Expected %s, got %s", expectedUrl, url)
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
