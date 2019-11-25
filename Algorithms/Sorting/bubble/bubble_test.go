package bubble

import (
	"reflect"
	"testing"
)

func Test_BubbleSort(t *testing.T) {
	arr := []int{2, 4, 6, 5, 3, 1, 7, 9, 8}
	BubbleSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Fatal(arr)
	}
}
