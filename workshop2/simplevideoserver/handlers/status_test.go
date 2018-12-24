package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	videoRepository := NewMockVideoRepository()
	videoID := videoRepository.videos[0].ID
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/video/"+videoID+"/status", nil)
	request = mux.SetURLVars(request, map[string]string{"ID": videoID})
	handleStatus(recorder, request, &videoRepository)
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
	var sStruct statusStruct
	if err = json.Unmarshal(jsonString, &sStruct); err != nil {
		t.Errorf("Can't parse json response with error %v", err)
	}
	if sStruct.Status != videoRepository.videos[0].Status {
		t.Error("Invalid status received")
	}

	videoID = "invalid"
	request = httptest.NewRequest(http.MethodGet, "/video/"+videoID+"/status", nil)
	request = mux.SetURLVars(request, map[string]string{"ID": videoID})
	recorder = httptest.NewRecorder()
	handleStatus(recorder, request, &videoRepository)
	response = recorder.Result()
	if response.StatusCode != http.StatusNotFound {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusNotFound)
	}
}
