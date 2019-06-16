package main
import "sort"
import "fmt"
func main() {
	nums := []int {1,-1,-1,0}
	ans := threeSum(nums)
	fmt.Println(ans)
}

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	count := len(nums)
	sort.Ints(nums)
	fmt.Println(nums)
	left := 0
	right := 0
	tempSum := 0
	tempLeft := 0
	tempRight := 0
	tempI := 0
	i:= 0
	for i < count - 2 && nums[i] <= 0 {
		left = i + 1
		right = count - 1
		for left < right && nums[right] >= 0{
			tempSum = nums[left] + nums[right] + nums[i]
			// fmt.Println(i, left, right, tempSum)
			if tempSum == 0 {
				res = append(res, []int{nums[left], nums[i], nums[right]})
			}

			if tempSum <= 0 {
				tempLeft = nums[left]
				left ++
				for left < right && tempLeft == nums[left]{
					left ++
				}
			}

			if tempSum >= 0 {
				tempRight = nums[right]
				right --
				for left < right && tempRight == nums[right]{
					right --
				}
			}
		}
		tempI = nums[i]
			i ++
		for i < count - 2 && tempI == nums[i]{
			i ++
		}
		
	}
	return res
}