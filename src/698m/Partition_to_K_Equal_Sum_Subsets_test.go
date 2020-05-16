/*
698. Partition to K Equal Sum Subsets

Given an array of integers nums and a positive integer k, find whether it's possible to divide this array into k non-empty subsets whose sums are all equal.

Example 1:

Input: nums = [4, 3, 2, 3, 5, 2, 1], k = 4
Output: True
Explanation: It's possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal sums.

*/

package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func canPartitionKSubsets(nums []int, k int) bool {
	total := 0
	for _, num := range nums {
		total += num
	}
	if total%k > 0 {
		return false
	}
	subTotal := total / k

	sort.Ints(nums)
	idx := len(nums) - 1
	if nums[idx] > subTotal {
		return false
	}

	for idx >= 0 && nums[idx] == subTotal {
		idx--
		k--
	}

	return search(make([]int, k), idx, nums, subTotal)
}

func search(groups []int, idx int, nums []int, subTotal int) bool {
	if idx < 0 {
		return true
	}
	v := nums[idx]
	idx--
	for i := 0; i < len(groups); i++ {
		if groups[i] + v <= subTotal {
			groups[i] += v
			if search(groups, idx, nums, subTotal) {
				return true
			}
			groups[i] -= v
		}
		if groups[i] == 0 {
			return false
		}
	}

	return false
}


func TestCase5(t *testing.T) {
	mockedNums := []int{5, 5, 5, 5, 5, 5}
	mockedK := 6

	assert.Equal(t, true, canPartitionKSubsets(mockedNums, mockedK))
}

func TestCase4(t *testing.T) {
	mockedNums := []int{5, 2, 5, 5, 3}
	mockedK := 4

	assert.Equal(t, true, canPartitionKSubsets(mockedNums, mockedK))
}

func TestCase3(t *testing.T) {
	mockedNums := []int{4, 16, 5, 3, 10, 4, 4, 4, 10}
	mockedK := 3

	assert.Equal(t, true, canPartitionKSubsets(mockedNums, mockedK))
}

func TestCase2(t *testing.T) {
	mockedNums := []int{2, 2, 2, 2, 3, 4, 5}
	mockedK := 4

	assert.Equal(t, false, canPartitionKSubsets(mockedNums, mockedK))
}

func TestCase1(t *testing.T) {
	mockedNums := []int{4, 3, 2, 3, 5, 2, 1}
	mockedK := 4

	assert.Equal(t, true, canPartitionKSubsets(mockedNums, mockedK))
}
