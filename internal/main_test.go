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

func TestAddTwoNumbers(t *testing.T) {

	// l1 := buildListNode(2, buildListNode(4, buildListNode(3, nil)))
	l1 := buildListNode(0, nil)
	// l2 := buildListNode(5, buildListNode(6, buildListNode(4, nil)))
	l2 := buildListNode(0, nil)

	got := addTwoNumbers(l1, l2)

	assert.DeepEqual(t, got, []int{7, 0, 8})
}

func buildListNode(val int, next *ListNode) *ListNode {
	return &ListNode{Val: val, Next: next}
}
