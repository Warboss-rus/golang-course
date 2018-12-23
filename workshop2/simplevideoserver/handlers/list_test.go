package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleList(t *testing.T) {
	recorder := httptest.NewRecorder()
	videoRepository := NewMockVideoRepository()
	request := httptest.NewRequest(http.MethodGet, "/list", nil)
	handleList(recorder, request, &videoRepository)
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
	items := make([]Video, 10)
	if err = json.Unmarshal(jsonString, &items); err != nil {
		t.Errorf("Can't parse json response with error %v", err)
	}
	if len(items) != len(videoRepository.videos) {
		t.Error("Invalid number of videos received")
	}
	for _, v := range items {
		video, err := videoRepository.GetVideoDetails(v.Id)
		if err != nil {
			t.Fatal(err)
		}
		if video != v {
			t.Error("Invalid video received")
		}
	}

	// bd error test
	videoRepository.errorToReturn = errors.New("bd error")
	recorder = httptest.NewRecorder()
	handleList(recorder, request, &videoRepository)
	response = recorder.Result()
	videoRepository.errorToReturn = nil
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusOK)
	}

	// search test
	recorder = httptest.NewRecorder()
	request = httptest.NewRequest(http.MethodGet, "/video?searchString=Black", nil)
	handleList(recorder, request, &videoRepository)
	response = recorder.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusOK)
	}
	jsonString, err = ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	items = make([]Video, 10)
	if err = json.Unmarshal(jsonString, &items); err != nil {
		t.Errorf("Can't parse json response with error %v", err)
	}
	if len(items) != 1 {
		t.Errorf("Expected 1 video, got %v", len(items))
	}
	if items[0] != videoRepository.videos[0] {
		t.Error("Invalid video received")
	}

	// skip and limit test
	recorder = httptest.NewRecorder()
	request = httptest.NewRequest(http.MethodGet, "/video?skip=1&limit=1", nil)
	handleList(recorder, request, &videoRepository)
	response = recorder.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusOK)
	}
	jsonString, err = ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	items = make([]Video, 10)
	if err = json.Unmarshal(jsonString, &items); err != nil {
		t.Errorf("Can't parse json response with error %v", err)
	}
	if len(items) != 1 {
		t.Errorf("Expected 1 video, got %v", len(items))
	}
	if items[0] != videoRepository.videos[1] {
		t.Error("Invalid video received")
	}

	// invalid skip
	recorder = httptest.NewRecorder()
	request = httptest.NewRequest(http.MethodGet, "/video?skip=-1", nil)
	handleList(recorder, request, &videoRepository)
	response = recorder.Result()
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusBadRequest)
	}

	// invalid limit
	recorder = httptest.NewRecorder()
	request = httptest.NewRequest(http.MethodGet, "/video?limit=51", nil)
	handleList(recorder, request, &videoRepository)
	response = recorder.Result()
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusBadRequest)
	}
}
