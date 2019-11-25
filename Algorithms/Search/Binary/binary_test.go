package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Search(t *testing.T) {
	assert := assert.New(t)
	sorted := []int{1, 2, 3, 4, 5, 6, 7, 9, 13, 16, 17, 19, 25, 30, 63}
	for index, target := range sorted {
		assert.Equal(search(sorted, target), index)
	}
	assert.Equal(search(sorted, 8), -1)
	assert.Equal(search(sorted, 26), -1)
	assert.Equal(search(sorted, 0), -1)
}
