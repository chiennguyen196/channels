package channels_test

import (
	"reflect"
	"testing"

	"github.com/chiennguyen196/channels"
)

func TestAsChan(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	gotChan := channels.AsChan(values...)

	actualValues := make([]int, len(values))
	i := 0
	for v := range gotChan {
		actualValues[i] = v
		i++
	}
	if !reflect.DeepEqual(actualValues, values) {
		t.Errorf("AsChan() = %v, want %v", actualValues, values)
	}
}
