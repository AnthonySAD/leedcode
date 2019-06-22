package main
import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1,2,2,0,-1,4}
	target := 3
	ans := fourSum(nums, target)
	fmt.Println(ans)
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)

	count := len(nums)
	ans := new([][]int)
	if count < 4 {
		return *ans
	}

	for i := 0; i < count - 3; i ++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		if nums[i] + nums[i + 1] + nums[i + 2] + nums[i + 3] > target {
			break
		}

		if nums[i] + nums[count - 1] + nums[count - 2] + nums[count - 3] < target {
			continue
		}

		for j := i + 1; j < count - 2; j ++ {
			if j - i > 1 && nums[j] == nums[j - 1] {
				continue
			}

			if nums[i] + nums[j] + nums[j + 1] + nums[j + 2] > target {
				break
			}
	
			if nums[i] + nums[j] + nums[count - 1] + nums[count - 2] < target {
				continue
			}

			left := j + 1
			right := count - 1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				fmt.Println(nums[i], nums[j], nums[left], nums[right], sum)
				if sum == target {
					*ans = append(*ans, []int{nums[i], nums[j], nums[left], nums[right]})
				}

				if sum <= target {
					for left < right && nums[left + 1] == nums[left]{
						left ++
					}
					left ++
				}

				if sum >= target {
					for left < right && nums[right - 1] == nums[right]{
						right --
					}
					right --
				}
			}

		}
	}

	return *ans
}