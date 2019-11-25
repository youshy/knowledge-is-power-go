package merge

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MergeSort(t *testing.T) {
	assert := assert.New(t)

	arr := []int{1, 5, 4, 6, 3, 7, 2, 8, 9, 19, 16, 13, 12, 14, 17, 15, 18, 11, 20, 10}
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

	assert.Equal([]int{1, 5, 4, 6, 3, 7, 2, 8, 9, 19, 16, 13, 12, 14, 17, 15, 18, 11, 20, 10}, arr)
	assert.Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, sorted)
	assert.Equal(sorted, MergeSort(arr))
}
