package videoprocessing

import "sync"

func RunWorkerPool(videosChan <-chan Video, repository VideoRepository, processor VideoProcessor) *sync.WaitGroup {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		go func(i int) {
			wg.Add(1)
			RunWorker(videosChan, repository, processor, i)
			wg.Done()
		}(i)
	}
	return &wg
}
