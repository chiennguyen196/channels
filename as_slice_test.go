package channels

import (
	"reflect"
	"testing"
)

func TestAsSlice(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual := AsSlice(AsChan(input...))
	if !reflect.DeepEqual(actual, input) {
		t.Errorf("AsSlice() = %v, want %v", actual, input)
	}
}
