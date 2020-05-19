package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 2, 2, 2, 3}
	res = removeDuplicates(nums)
	fmt.Print(res)
	fmt.Print(nums)
}

func removeDuplicates(nums []int) int {

	key := 0
	if len(nums) < 2 {
		return len
	}
	for i := 1; i < len(nums); i++ {
		if nums[key] != nums[i] {
			key++
			nums[key] = nums[i]
		}
	}

	return key + 1
}
