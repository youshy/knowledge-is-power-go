package quick

import (
	"reflect"
	"testing"
)

func Test_QuickSort(t *testing.T) {
	arr := []int{5, 4, 6, 3, 7, 2, 8, 1, 9}
	if QuickSort(arr); !reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Fatal(arr)
	}
}
