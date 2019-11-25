package shell

import (
	"reflect"
	"testing"
)

func Test_ShellSort(t *testing.T) {
	arr := []int{6, 7, 5, 8, 9, 3, 2, 4, 1}
	ShellSort(arr)
	if !reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Fatal(arr)
	}
}
