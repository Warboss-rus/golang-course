package videoprocessing

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"testing"
	"time"
)

func getVideoFromChan(ch <-chan Video) *Video {
	select {
	case v := <-ch:
		return &v
	default:
		return nil
	}
}

func TestWatchVideos(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	stopChan := make(chan struct{})
	var repository mockVideoRepository
	videoChan := WatchVideos(stopChan, &repository, 100*time.Millisecond)

	if getVideoFromChan(videoChan) != nil {
		t.Error("should not return video yet")
	}

	// get one video from repository
	expectedVideo := Video{"videoId", "videoUrl"}
	repository.videos = append(repository.videos, expectedVideo)
	time.Sleep(time.Second)
	gotVideo := getVideoFromChan(videoChan)
	if gotVideo == nil {
		t.Error("Video expected, got nil")
	}
	if *gotVideo != expectedVideo {
		t.Error("Invalid video received")
	}
	if v2 := getVideoFromChan(videoChan); v2 != nil {
		t.Error("Only one video expected")
	}

	// multiple video test
	expectedVideos := []Video{
		{"video1", "videourl1"},
		{"video2", "videourl2"},
		{"video3", "videourl3"},
		{"video4", "videourl4"},
		{"video5", "videourl5"},
	}
	repository.videos = expectedVideos
	time.Sleep(time.Millisecond * 500)
	for _, v := range expectedVideos {
		gotVideo := <-videoChan
		if gotVideo != v {
			t.Error("did not recieve expected video")
		}
	}
	if v2 := getVideoFromChan(videoChan); v2 != nil {
		t.Error("No more videos expected")
	}

	// should close channel after stop channel fires
	stopChan <- struct{}{}
	time.Sleep(time.Second)
	_, open := <-videoChan
	if open {
		t.Error("Didn't close video channel after stop channel fired")
	}
}
