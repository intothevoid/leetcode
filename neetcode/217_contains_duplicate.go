package neetcode

func ContainsDuplicate(nums []int) bool {
	newmap := make(map[int]int)

	for _, num := range nums {
		if _, ok := newmap[num]; ok {
			return true
		}
		newmap[num] = num
	}

	return false
}
