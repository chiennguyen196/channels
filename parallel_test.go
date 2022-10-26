package channels_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/chiennguyen196/channels"
)

func TestParallel(t *testing.T) {
	n := 100
	giveArr := make([]int, n)
	for i := range giveArr {
		giveArr[i] = i
	}
	inputChan := channels.AsChan(giveArr...)
	outChan := make(chan int)

	channels.Parallel(
		4, // num workers
		func(_ int) { // execute function
			for v := range inputChan {
				outChan <- v
			}
		},
		func() { // clean up function
			close(outChan)
		},
	)

	var output []int
	for v := range outChan {
		output = append(output, v)
	}
	sort.Ints(output)

	if !reflect.DeepEqual(giveArr, output) {
		t.Errorf("Parallel() = %v, want %v", output, giveArr)
	}
}
