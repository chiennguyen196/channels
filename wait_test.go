package channels

import (
	"sync/atomic"
	"testing"
)

func TestWait(t *testing.T) {
	var count uint32
	n := 100

	execFn := func() {
		atomic.AddUint32(&count, 1)
	}

	inputFns := make([]func(), n)
	for i := 0; i < n; i++ {
		inputFns[i] = execFn
	}

	Wait(inputFns...)

	if int(count) != n {
		t.Errorf("After Wait() got count = %d, but want %d", count, n)
	}
}
