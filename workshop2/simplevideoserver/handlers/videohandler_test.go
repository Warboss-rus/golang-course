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
	videoId := videos[0].Id
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/video/"+videoId, nil)
	request = mux.SetURLVars(request, map[string]string{"ID": videoId})
	handleVideo(recorder, request)
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
	if v.Id != videoId {
		t.Error("Invalid video received")
	}
	video, err := findVideoById(videoId)
	if err != nil {
		t.Fatal(err)
	}
	if video != v {
		t.Error("Invalid video received")
	}

	videoId = "invalid"
	request = httptest.NewRequest(http.MethodGet, "/video/"+videoId, nil)
	request = mux.SetURLVars(request, map[string]string{"ID": videoId})
	recorder = httptest.NewRecorder()
	handleVideo(recorder, request)
	response = recorder.Result()
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusInternalServerError)
	}
}
