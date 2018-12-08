package handlers

import "testing"

func TestVideoList(t *testing.T) {
	for _, v := range videos {
		video, err := findVideoById(v.Id)
		if err != nil {
			t.Fatal(err)
		}
		if video != v {
			t.Error("Invalid video")
		}
	}
	_, err := findVideoById("invalid id")
	if err == nil {
		t.Error("Error expected when requesting video with invalid id, got nil")
	}
}
