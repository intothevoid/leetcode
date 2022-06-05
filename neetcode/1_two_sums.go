package neetcode

func TwoSum(nums []int, target int) []int {

	numlen := len(nums)
	for idx, val := range nums {
		for i := 1; i < numlen; i++ {
			if val+nums[i] == target {
				return []int{idx, i}
			}
		}
	}

	return []int{}
}
