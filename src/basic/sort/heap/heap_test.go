package main

import (
	"reflect"
	"testing"
)

func heapSort(nums []int) []int {

	nums = append([]int{0}, nums...)

	for i := 1; i < len(nums); i++ {
		nums[0] = nums[i]
		hole := i
		for ; nums[0] < nums[hole/2]; hole /= 2 {
			nums[hole] = nums[hole/2]
		}
		nums[hole] = nums[0]
	}

	nums = nums[1:]
	sorted := make([]int, len(nums))

	for i := 0; i < len(sorted); i++ {
		sorted[i] = nums[0]

		hole := 0
		child := 0
		for ; 2*hole+2 < len(nums); hole = child {
			child = 2*hole + 1
			if child+1 != len(nums) && nums[child] > nums[child+1] {
				child++
			}
			nums[hole] = nums[child]

		}
		nums[hole] = nums[len(nums)-1]
		nums = nums[:len(nums)-1]
	}

	return sorted
}

func Test_heapSort(t *testing.T) {
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
		{name: "Case 2", args: args{nums: []int{12, 32, 15, 29, 24, 8, 28}}, want: []int{8, 12, 15, 24, 28, 29, 32}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heapSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
