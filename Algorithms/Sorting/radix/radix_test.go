package radix

import (
	"reflect"
	"testing"
)

func Test_RadixSort(t *testing.T) {
	arr := []int{6, 128, 5, 24, 4, 7, 3, 1, 8, 2, 9, 12}
	RadixSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 128}) {
		t.Fatal(arr)
	}
}
