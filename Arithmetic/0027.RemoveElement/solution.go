package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 2, 3}
	res := removeElement(nums, 3)
	fmt.Print(res)
	fmt.Print(nums)
}

func removeElement(nums []int, val int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {
		if nums[left] == val {
			nums[left] = nums[right]
			right--
		} else {
			left++
		}
	}

	return left
}
