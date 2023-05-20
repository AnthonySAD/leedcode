package solution_test

import "testing"
import "fmt"

func lengthOfLISOld(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	for k, v := range nums {
		num := 0
		for i := 0; i < k; i++ {
			if nums[i] < v {
				if dp[i] > num {
					num = dp[i]
				}
			}
		}
		dp[k] = num + 1
	}

	max := 0
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	max := 0
	for _, v := range nums {
		i := 0
		j := max
		//二分法查找小于元素v的最右侧元素
		for i < j {
			m := (i + j) / 2
			if dp[m] < v {
				i = m + 1
			} else {
				j = m
			}
		}
		dp[i] = v
		if j == max {
			max++
		}
	}
	return max
}

func TestFunc(t *testing.T) {
	t.Log(lengthOfLIS([]int{7, 7, 7, 7, 7}))

}
