package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	videoConnector := NewMockVideosConnector()
	var fs MockFilesHandler
	r := Router(&videoConnector, &fs)

	// video list test
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

	// video details test
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

	// upload video test
	recorder = httptest.NewRecorder()
	const path = "..\\content\\d290f1ee-6c54-4b01-90e6-d701748f0851\\index.mp4"
	request, _ = newfileUploadTestRequest("/api/v1/video", path, "video/mp4")

	r.ServeHTTP(recorder, request)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if len(videoConnector.videos) != 4 {
		t.Error("handler should add a new video to the list")
	}
}
