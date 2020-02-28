package main

import (
	"reflect"
	"testing"
)

func shellSort(nums []int) []int {

	for gap := len(nums) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(nums); i++ {
			tmp := nums[i]
			j := i
			for ; j >= gap && nums[j-gap] > tmp; j -= gap {
				nums[j] = nums[j-gap]
			}
			nums[j] = tmp
		}
	}

	return nums
}

func Test_shellSort(t *testing.T) {
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
		{name: "Case 2", args: args{nums: []int{32, 12, 15, 29, 24, 8, 28}}, want: []int{8, 12, 15, 24, 28, 29, 32}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shellSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shellSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 1. hk should be an increment sequence, such as h0 = N/2 and hk = h(k-1)/2: Must start from N/2 so that can assure the sequence is ordered within 2 elements
// 2. In each hk for loop, all the sequence with hk gap should be ordered
//
