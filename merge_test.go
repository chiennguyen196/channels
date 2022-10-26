package channels_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/chiennguyen196/channels"
)

func TestMerge(t *testing.T) {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	arr3 := []int{7, 8, 9}
	expectValues := append(append(arr1, arr2...), arr3...)

	outChan := channels.Merge(
		channels.AsChan(arr1...),
		channels.AsChan(arr2...),
		channels.AsChan(arr3...),
	)

	actualValues := make([]int, 0)
	for v := range outChan {
		actualValues = append(actualValues, v)
	}

	sort.Ints(actualValues)
	sort.Ints(expectValues)

	if !reflect.DeepEqual(actualValues, expectValues) {
		t.Errorf("Merge() = %v, want %v", actualValues, expectValues)
	}
}
