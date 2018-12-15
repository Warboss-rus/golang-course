package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockFilesHandler struct {
}

func (fs *MockFilesHandler) CreateFile(id string, filename string, content io.Reader) error {
	return nil
}

func TestRouter(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	videoConnector := NewMockVideosConnector()
	var fs MockFilesHandler
	r := Router(&videoConnector, &fs)

	request, err := http.NewRequest(http.MethodGet, "/api/v1/list", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected, _ := json.Marshal(videoConnector.videos)
	if recorder.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}

	var videoId = videoConnector.videos[0].Id
	request, err = http.NewRequest(http.MethodGet, "/api/v1/video/"+videoId, nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder = httptest.NewRecorder()

	r.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected, _ = json.Marshal(videoConnector.videos[0])
	if recorder.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			recorder.Body.String(), expected)
	}
}
