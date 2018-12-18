package handlers

import "testing"

func TestDataBaseConnector(t *testing.T) {
	var db DataBaseConnector
	defer func() {
		if err := db.Close(); err != nil {
			t.Error("Cannot close database")
		}
	}()

	if _, err := db.GetVideoList("", nil, nil); err == nil {
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

		videos, err := db.GetVideoList("", nil, nil)
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

		for _, v := range videoList.videos {
			status, err := db.GetVideoStatus(v.Id)
			if err != nil || status != v.Status {
				t.Error("Invalid status received")
			}
		}

		//search test
		videos, err = db.GetVideoList("Rally", nil, nil)
		if err != nil {
			t.Error("Getting list of videos failed with: ", err)
		}
		if len(videos) != 1 {
			t.Errorf("Expected 1 video, got %v", len(videos))
		}
		if videos[0] != videoList.videos[1] {
			t.Error("Invalid video received")
		}

		// skip and limit test
		var start uint = 2
		var count uint = 1
		videos, err = db.GetVideoList("", &start, &count)
		if err != nil {
			t.Error("Getting list of videos failed")
		}
		if len(videos) != 1 {
			t.Errorf("Expected 1 video, got %v", len(videos))
		}
		if videos[0] != videoList.videos[2] {
			t.Error("Invalid video received")
		}

		// invalid id test
		var invalidId = "invalid"
		_, err = db.GetVideoDetails(invalidId)
		if _, ok := err.(*VideoNotFound); !ok {
			t.Error("VideoNotFound error expected")
		}

		_, err = db.GetVideoStatus(invalidId)
		if _, ok := err.(*VideoNotFound); !ok {
			t.Error("VideoNotFound error expected")
		}
	}
}
