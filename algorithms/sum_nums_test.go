package algorithms

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSumNums(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5}
	sum1 := sumInList(list1)
	assert.Equal(t, sum1, 15)

	list2 := []int{3, 0, -10, 2}
	sum2 := sumInList(list2)
	assert.Equal(t, sum2, -5)

	list3 := []int{}
	sum3 := sumInList(list3)
	assert.Equal(t, sum3, 0)

	sum4 := sumInList(nil)
	assert.Equal(t, sum4, 0)

	list5 := []int{1}
	sum5 := sumInList(list5)
	assert.Equal(t, sum5, 1)
}
