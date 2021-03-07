package internal

func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[target-nums[i]]; ok {
			return []int{hash[target-nums[i]], i}
		}
		hash[nums[i]] = i
	}

	return []int{}
}
