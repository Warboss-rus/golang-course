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
	videoConnector := NewMockVideosConnector()
	videoId := videoConnector.videos[0].Id
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/video/"+videoId+"/status", nil)
	request = mux.SetURLVars(request, map[string]string{"ID": videoId})
	handleStatus(recorder, request, &videoConnector)
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
	if sStruct.Status != videoConnector.videos[0].Status {
		t.Error("Invalid status received")
	}

	videoId = "invalid"
	request = httptest.NewRequest(http.MethodGet, "/video/"+videoId+"/status", nil)
	request = mux.SetURLVars(request, map[string]string{"ID": videoId})
	recorder = httptest.NewRecorder()
	handleStatus(recorder, request, &videoConnector)
	response = recorder.Result()
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusInternalServerError)
	}
}
