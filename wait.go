package channels

import (
	"sync"
)

// Wait is used to execute many functions parallel and wait all done.
func Wait(execFns ...func()) {
	var wg sync.WaitGroup
	wg.Add(len(execFns))
	for _, fn := range execFns {
		fn := fn
		go func() {
			fn()
			wg.Done()
		}()
	}
	wg.Wait()
}
