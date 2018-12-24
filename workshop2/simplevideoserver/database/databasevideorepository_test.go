package database

import (
	"github.com/Warboss-rus/golang-course/workshop4/videoProcessingDaemon/videoprocessing"
	"testing"
)
import . "github.com/Warboss-rus/golang-course/workshop2/simplevideoserver/handlers"

func TestDataBaseVideoRepository(t *testing.T) {
	var repository DBVideoRepository
	defer func() {
		if err := repository.Close(); err != nil {
			t.Error("Cannot close database")
		}
	}()

	if _, err := repository.GetVideoList("", nil, nil); err == nil {
		t.Error("Database connector should return error if not connected")
	}

	const user = "root"
	const password = "root"
	const dbname = "dbtest"
	if err := repository.Connect(dbname, user, password); err == nil {
		// test database is available, if its not that's OK too, don't fail the test
		defer func() {
			if err := repository.removeTable(); err != nil {
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
			video, err := mockVideoRepository.GetVideoDetails(v.ID)
			if err != nil || video != v {
				t.Error("Invalid video received")
			}
		}

		for _, v := range mockVideos {
			video, err := repository.GetVideoDetails(v.ID)
			if err != nil || video != v {
				t.Error("Invalid video received")
			}
		}

		for _, v := range mockVideos {
			status, err := repository.GetVideoStatus(v.ID)
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
		var invalidID = "invalid"
		_, err = repository.GetVideoDetails(invalidID)
		if _, ok := err.(*VideoNotFound); !ok {
			t.Error("VideoNotFound error expected")
		}

		_, err = repository.GetVideoStatus(invalidID)
		if _, ok := err.(*VideoNotFound); !ok {
			t.Error("VideoNotFound error expected")
		}

		// videos by status test
		processingVideos, err := repository.GetVideosByStatus(videoprocessing.Created)
		if err != nil {
			t.Error("Getting list of videos failed")
		}
		if len(processingVideos) != 1 {
			t.Error("Invalid number of videos received")
		}
		if processingVideos[0].ID != mockVideos[0].ID || processingVideos[0].URL != mockVideos[0].URL {
			t.Error("Invalid video received")
		}

		// empty result test
		processingVideos, err = repository.GetVideosByStatus(videoprocessing.Deleted)
		if err != nil {
			t.Error("Getting list of videos failed")
		}
		if len(processingVideos) != 0 {
			t.Error("Invalid number of videos received")
		}

		// update video status test
		videoID := mockVideos[0].ID
		err = repository.UpdateVideoStatus(videoID, videoprocessing.Processing)
		if err != nil {
			t.Error("Failed to update video status")
		}
		status, err := repository.GetVideoStatus(videoID)
		if err != nil || status != Processing {
			t.Error("Status was not correctly updated")
		}
		processingVideos, err = repository.GetVideosByStatus(videoprocessing.Created)
		if err != nil || len(processingVideos) != 0 {
			t.Error("Status was not correctly updated")
		}

		// update video test
		const duration = 987
		const thumbnail = "content\\preview.jpg"
		err = repository.UpdateVideo(videoID, duration, thumbnail, videoprocessing.Ready)
		if err != nil {
			t.Error("Video updating failed")
		}
		v, err := repository.GetVideoDetails(videoID)
		if err != nil || v.Duration != duration || v.Thumbnail != thumbnail || v.Status != Ready {
			t.Error("Invalid data received after updating")
		}
	}
}
