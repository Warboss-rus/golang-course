package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"path/filepath"
	"testing"
)

type MockFilesHandler struct {
	errorToReturn error
	fileCreated   bool
}

func (fs *MockFilesHandler) StoreFile(filename string, content io.Reader) (string, error) {
	fs.fileCreated = true
	return filepath.Join("content", filename), fs.errorToReturn
}

func newfileUploadTestRequest(uri string, path string, contentType string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file[]", path))
	h.Set("Content-Type", contentType)
	part, err := writer.CreatePart(h)

	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

func TestVideoUpload(t *testing.T) {
	recorder := httptest.NewRecorder()
	const path = "..\\content\\d290f1ee-6c54-4b01-90e6-d701748f0851\\index.mp4"
	request, _ := newfileUploadTestRequest("/video", path, "video/mp4")
	videoRepository := NewMockVideoRepository()
	var fs MockFilesHandler
	handleVideoUpload(recorder, request, &videoRepository, &fs)

	response := recorder.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusOK)
	}

	if len(videoRepository.videos) != 4 {
		t.Error("handleVideoUpload doesn't add a new video")
	}

	if !fs.fileCreated {
		t.Error("handleVideoUpload doesn't create a new file")
	}

	// invalid content type test
	recorder = httptest.NewRecorder()
	const path2 = "..\\content\\d290f1ee-6c54-4b01-90e6-d701748f0851\\screen.jpg"
	request, _ = newfileUploadTestRequest("/video", path2, "image/jpeg")
	handleVideoUpload(recorder, request, &videoRepository, &fs)

	response = recorder.Result()
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusBadRequest)
	}

	//invalid file test
	recorder = httptest.NewRecorder()
	request = httptest.NewRequest(http.MethodPost, "/video", nil)
	handleVideoUpload(recorder, request, &videoRepository, &fs)

	response = recorder.Result()
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusInternalServerError)
	}

	// fs error test
	fs.errorToReturn = errors.New("fs error")
	recorder = httptest.NewRecorder()
	request, _ = newfileUploadTestRequest("/video", path, "video/mp4")
	handleVideoUpload(recorder, request, &videoRepository, &fs)
	response = recorder.Result()
	fs.errorToReturn = nil
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusInternalServerError)
	}

	// bd error test
	videoRepository.errorToReturn = errors.New("bd error")
	recorder = httptest.NewRecorder()
	request, _ = newfileUploadTestRequest("/video", path, "video/mp4")
	handleVideoUpload(recorder, request, &videoRepository, &fs)
	response = recorder.Result()
	videoRepository.errorToReturn = nil
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusInternalServerError)
	}
}
