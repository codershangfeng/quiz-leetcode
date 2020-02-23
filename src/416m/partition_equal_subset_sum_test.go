/*
# 416. Partition Equal Subset Sum

Given a non-empty array containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

Note:

Each of the array element will not exceed 100.
The array size will not exceed 200.


Example 1:

Input: [1, 5, 11, 5]

Output: true

Explanation: The array can be partitioned as [1, 5, 5] and [11].


Example 2:

Input: [1, 2, 3, 5]

Output: false

Explanation: The array cannot be partitioned into equal sum subsets.
*/

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func canPartition(nums []int) bool {

	total := 0
	max := 0
	for _, num := range nums {
		total += num
		if max < num {
			max = num
			num = 0
		}
	}

	if total%2 != 0 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		threshold := total/2 - max

		if threshold == 0 {
			return true
		}

		for j := i; j < len(nums); j++ {
			t := threshold - nums[j]
			if t < 0 {
				continue
			}
			if t == 0 {
				return true
			}
			threshold = t
		}
	}

	return false
}

func TestCase6(t *testing.T) {
	mocked := []int{2, 2, 3, 5}

	assert.Equal(t, false, canPartition(mocked))
}

func TestCase5(t *testing.T) {
	mocked := []int{1, 2, 5}

	assert.Equal(t, false, canPartition(mocked))
}

func TestCase1(t *testing.T) {
	mocked := []int{1, 5, 11, 5}

	assert.Equal(t, true, canPartition(mocked))
}

func TestCase2(t *testing.T) {
	mocked := []int{1, 5, 5, 13, 6, 4}

	assert.Equal(t, true, canPartition(mocked))
}

func TestCase3(t *testing.T) {
	mocked := []int{1, 2, 3, 5}

	assert.Equal(t, false, canPartition(mocked))
}

func TestCase4(t *testing.T) {
	mocked := []int{11, 5, 1, 5}

	assert.Equal(t, true, canPartition(mocked))
}
