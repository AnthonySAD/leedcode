package main
import (
	"fmt"
)

func main(){
	height := []int{1,8,6,2,5,4,8,3,7}
	ans := maxArea(height)
	fmt.Println(ans)
}

func maxArea(height []int) int {
	length := len(height)
	i := 0
	j := length - 1
	max := 0
	tempA := 0
	for i < j {
		if height[i] < height[j] {
			tempA = height[i] * (j - i)
			i ++
		}else{
			tempA = height[j] * (j - i)
			j --
		}
		if tempA > max  {
			max = tempA
		}
	}
	return max
}

func maxArea1(height []int) int {
	length := len(height)
	max := 0
	tempH := 0
	tempA := 0
	for i := 0; i < length - 1; i++ {
		for j := i + 1; j < length; j++ {
			if height[i] < height[j] {
				tempH = height[i]
			}else{
				tempH = height[j]
			}
			
			tempA = tempH * (j - i)
			if tempA > max {
				max = tempA
			}
		}
	}
	return max
}