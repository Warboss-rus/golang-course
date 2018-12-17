package handlers

import "testing"

func TestDataBaseConnector(t *testing.T) {
	var db DataBaseConnector
	defer func() {
		if err := db.Close(); err != nil {
			t.Error("Cannot close database")
		}

	}()

	if _, err := db.GetVideoList(); err == nil {
		t.Error("Database connector should return error if not connected")
	}

	if err := db.ConnectTestDatabase(); err == nil {
		// test database is available, if its not that's OK too, don't fail the test
		defer func() {
			if err := db.ClearVideos(); err != nil {
				t.Error("Cannot clear database")
			}

		}()

		videoList := NewMockVideosConnector()
		for _, v := range videoList.videos {
			err = db.AddVideo(v)
			if err != nil {
				t.Error("Adding video failed")
			}
		}

		videos, err := db.GetVideoList()
		if err != nil {
			t.Error("Getting list of videos failed")
		}
		if len(videos) != len(videoList.videos) {
			t.Error("Invalid number of videos received")
		}
		for _, v := range videos {
			video, err := videoList.GetVideoDetails(v.Id)
			if err != nil || video != v {
				t.Error("Invalid video received")
			}
		}

		for _, v := range videoList.videos {
			video, err := db.GetVideoDetails(v.Id)
			if err != nil || video != v {
				t.Error("Invalid video received")
			}
		}
	}
}
