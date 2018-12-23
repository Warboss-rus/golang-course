package database

import "testing"
import . "github.com/Warboss-rus/golang-course/workshop2/simplevideoserver/handlers"

func TestDataBaseVideoRepository(t *testing.T) {
	var repository DataBaseVideoRepository
	defer func() {
		if err := repository.Close(); err != nil {
			t.Error("Cannot close database")
		}
	}()

	if _, err := repository.GetVideoList("", nil, nil); err == nil {
		t.Error("Database connector should return error if not connected")
	}

	if err := repository.ConnectTestDatabase(); err == nil {
		// test database is available, if its not that's OK too, don't fail the test
		defer func() {
			if err := repository.ClearVideos(); err != nil {
				t.Error("Cannot clear database")
			}
		}()

		mockVideoRepository := NewMockVideoRepository()
		mockVideos, _ := mockVideoRepository.GetVideoList("", nil, nil)
		for _, v := range mockVideos {
			err = repository.AddVideo(v)
			if err != nil {
				t.Error("Adding video failed")
			}
		}

		videos, err := repository.GetVideoList("", nil, nil)
		if err != nil {
			t.Error("Getting list of videos failed")
		}
		if len(videos) != len(mockVideos) {
			t.Error("Invalid number of videos received")
		}
		for _, v := range videos {
			video, err := mockVideoRepository.GetVideoDetails(v.Id)
			if err != nil || video != v {
				t.Error("Invalid video received")
			}
		}

		for _, v := range mockVideos {
			video, err := repository.GetVideoDetails(v.Id)
			if err != nil || video != v {
				t.Error("Invalid video received")
			}
		}

		for _, v := range mockVideos {
			status, err := repository.GetVideoStatus(v.Id)
			if err != nil || status != v.Status {
				t.Error("Invalid status received")
			}
		}

		//search test
		videos, err = repository.GetVideoList("Rally", nil, nil)
		if err != nil {
			t.Error("Getting list of videos failed with: ", err)
		}
		if len(videos) != 1 {
			t.Errorf("Expected 1 video, got %v", len(videos))
		}
		if videos[0] != mockVideos[1] {
			t.Error("Invalid video received")
		}

		// skip and limit test
		var start uint = 2
		var count uint = 1
		videos, err = repository.GetVideoList("", &start, &count)
		if err != nil {
			t.Error("Getting list of videos failed")
		}
		if len(videos) != 1 {
			t.Errorf("Expected 1 video, got %v", len(videos))
		}
		if videos[0] != mockVideos[2] {
			t.Error("Invalid video received")
		}

		// invalid id test
		var invalidId = "invalid"
		_, err = repository.GetVideoDetails(invalidId)
		if _, ok := err.(*VideoNotFound); !ok {
			t.Error("VideoNotFound error expected")
		}

		_, err = repository.GetVideoStatus(invalidId)
		if _, ok := err.(*VideoNotFound); !ok {
			t.Error("VideoNotFound error expected")
		}
	}
}
