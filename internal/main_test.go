package internal

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	got := twoSum(nums, target)

	assert.DeepEqual(t, got, []int{0, 1})

	nums = []int{3, 2, 4}
	target = 6

	got = twoSum(nums, target)

	assert.DeepEqual(t, got, []int{1, 2})
}
