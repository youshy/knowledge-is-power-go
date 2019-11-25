package selection

import (
	"reflect"
	"testing"
)

func Test_SelectionSort(t *testing.T) {
	arr := []int{6, 4, 5, 7, 3, 8, 9, 2, 1}
	SelectionSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Fatal(arr)
	}
}
