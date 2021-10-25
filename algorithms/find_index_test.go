package algorithms

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestFindIndex(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5}
	index1, exists1 := findIndex(list1, 5)
	assert.Equal(t, index1, 4)
	assert.Equal(t, exists1, true)

	list2 := []int{1, 5, 3, 4, 5}
	index2, exists2 := findIndex(list2, 5)
	assert.Equal(t, index2, 1)
	assert.Equal(t, exists2, true)

	list3 := []int{}
	index3, exists3 := findIndex(list3, 2)
	assert.Equal(t, index3, 0)
	assert.Equal(t, exists3, false)

	list4 := []int{}
	index4, exists4 := findIndex(list4, 1)
	assert.Equal(t, index4, 0)
	assert.Equal(t, exists4, false)
}
