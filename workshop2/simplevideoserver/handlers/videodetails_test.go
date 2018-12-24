package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleVideo(t *testing.T) {
	videoRepository := NewMockVideoRepository()
	videoID := videoRepository.videos[0].ID
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/video/"+videoID, nil)
	request = mux.SetURLVars(request, map[string]string{"ID": videoID})
	handleVideoDetails(recorder, request, &videoRepository)
	response := recorder.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusOK)
	}

	jsonString, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	var v Video
	if err = json.Unmarshal(jsonString, &v); err != nil {
		t.Errorf("Can't parse json response with error %v", err)
	}
	if v.ID != videoID {
		t.Error("Invalid video received")
	}
	video, err := videoRepository.GetVideoDetails(videoID)
	if err != nil {
		t.Fatal(err)
	}
	if video != v {
		t.Error("Invalid video received")
	}

	videoID = "invalid"
	request = httptest.NewRequest(http.MethodGet, "/video/"+videoID, nil)
	request = mux.SetURLVars(request, map[string]string{"ID": videoID})
	recorder = httptest.NewRecorder()
	handleVideoDetails(recorder, request, &videoRepository)
	response = recorder.Result()
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusBadRequest)
	}
}
