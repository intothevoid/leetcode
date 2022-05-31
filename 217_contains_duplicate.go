package main

import "fmt"

func containsDuplicate(nums []int) bool {
	newmap := map[int]int{}

	for num := range nums {
		if _, ok := newmap[num]; ok {
			return true
		}
		newmap[num] = num
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}
