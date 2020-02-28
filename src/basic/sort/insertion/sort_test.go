package main

import (
	"reflect"
	"testing"
)

// Put the i-th numer into its correct place in each for loop by i-th
func insertionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func Test_insertionSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Case 1", args: args{nums: []int{8}}, want: []int{8}},
		{name: "Case 2", args: args{nums: []int{32, 12, 15, 29, 24, 8, 28}}, want: []int{8, 12, 15, 24, 28, 29, 32}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
