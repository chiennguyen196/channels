package channels

import (
	"sync"
)

// Parallel is used to run many worker do same thing at once.
// We can use goroutines and WaitGroup but some time got mistake when using wg.Done(), wg.Wait().
// This function provide a simple interface to wrap the WaitGroup.
func Parallel(numWorkers int, executeFunc func(workerIndex int), cleanUp func()) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(i int) {
			executeFunc(i)
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		cleanUp()
	}()
	return
}
