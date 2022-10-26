package channels_test

import (
	"sync/atomic"
	"testing"

	"github.com/chiennguyen196/channels"
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

	channels.Wait(inputFns...)

	if int(count) != n {
		t.Errorf("After Wait() got count = %d, but want %d", count, n)
	}
}
