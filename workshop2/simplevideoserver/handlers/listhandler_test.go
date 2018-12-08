package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleList(t *testing.T) {
	recorder := httptest.NewRecorder()
	handleList(recorder, nil)
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
	if len(items) != len(videos) {
		t.Error("Invalid number of videos received")
	}
	for _, v := range items {
		video, err := findVideoById(v.Id)
		if err != nil {
			t.Fatal(err)
		}
		if video != v {
			t.Error("Invalid video received")
		}
	}
}
