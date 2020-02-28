package main

import (
	"reflect"
	"testing"
)

// Max heap
func procolateUp(nums []int) []int {
	heap := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		hole := i + 1
		for heap[0] = nums[i]; heap[0] > heap[hole/2]; hole /= 2 {
			heap[hole] = heap[hole/2]
		}
		heap[hole] = heap[0]
	}

	return heap[1:]
}

func Test_procolateUp(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "Case 1", args: args{nums: []int{8}}, want: []int{8}},
		{name: "Case 2", args: args{nums: []int{32, 12, 15, 29, 24, 8, 28}}, want: []int{32, 29, 28, 12, 24, 8, 15}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := procolateUp(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("percDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

// DeleteMax
func procolateDown(nums []int) []int {
	nums = append([]int{0}, nums...)

	for i := 1; i < len(nums); i++ {
		nums[0] = nums[i]
		hole := i
		for ; nums[0] > nums[hole/2]; hole /= 2 {
			nums[hole] = nums[hole/2]
		}
		nums[hole] = nums[0]
	}

	return nums[1:]
}

func Test_procolateDown(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "Case 1", args: args{nums: []int{69, 38, 29, 16, 27, 14, 12, 8}}, want: []int{69, 38, 29, 16, 27, 14, 12, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := procolateDown(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("procolateDown() = %v, want %v", got, tt.want)
			}
		})
	}
}
