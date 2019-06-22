package main
import "sort"
import "math"
import "fmt"

func main() {
	var nums []int = []int{0, 2, 1, -3}
	var target int = 1
	ans := threeSumClosest(nums, target)
	fmt.Println(ans)
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	fmt.Println(nums)

	ans := nums[0] + nums[1] + nums[2]
	diff := int(math.Abs(float64(target - ans)))
	count := len(nums)
	for i := 0; i < count - 2; i++ {
		left := i + 1
		right := count - 1
		for left < right{
			sum := nums[i] + nums[left] + nums[right]
			fmt.Printf("i is %d, left is %d, right is %d, sum is %d\n", i, left, right, sum)
			minus := sum - target
			tempDiff := int(math.Abs(float64(minus)))
			if tempDiff < diff{
				ans = sum
				diff = tempDiff
			}

			if minus > 0{
				right --

			}else if minus < 0{
				left ++

			}else{
				return ans
			}
		}
	}

	return ans
}

func wrong(nums []int, target int) int {
	sort.Ints(nums)
	sum := nums[0] + nums[1] + nums[2]
	diff := int(math.Abs(float64(target - sum)))
	fmt.Println(nums)
	count := len(nums)
	if count == 3 {
		return sum
	}
	for i := 1; i < count - 2; i++ {
		tempSum := nums[i] + nums[i + 1] + nums[i + 2];
		tempDiff := int(math.Abs(float64(target - tempSum))) 
		if tempDiff == 0{
			return tempSum
		}
		if tempDiff < diff {
			sum = tempSum
			diff = tempDiff
		}else{
			return sum
		}
	}

	return sum
}