package videoprocessing

import "sync"

// RunWorkerPool runs 3 concurrent workers and returns a waiting group. Close the videosChan to close the workers and wait using the waiting group
func RunWorkerPool(videosChan <-chan Video, repository VideoRepository, processor VideoProcessor, contentPath string) *sync.WaitGroup {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		go func(i int) {
			wg.Add(1)
			RunWorker(videosChan, repository, processor, contentPath, i)
			wg.Done()
		}(i)
	}
	return &wg
}
