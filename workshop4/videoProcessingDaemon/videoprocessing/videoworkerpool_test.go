package videoprocessing

import (
	"testing"
	"time"
)

func TestRunWorkerPool(t *testing.T) {
	videosChan := make(chan Video)
	var repository mockVideoRepository
	var processor mockVideoProcessor
	wg := RunWorkerPool(videosChan, &repository, &processor, "")
	// give time to start goroutimes
	time.Sleep(time.Millisecond * 100)
	close(videosChan)
	wg.Wait()
}
