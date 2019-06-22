package main

import (
	"fmt"
)

func main(){
	digits := "23"
	ans := letterCombinations(digits)
	fmt.Println(ans)
}

func letterCombinations(digits string) []string {
	
	value := map[rune]string{
		'2' : "abc",
		'3' : "def",
		'4' : "ghi",
		'5' : "jkl",
		'6' : "mno",
		'7' : "pqrs",
		'8' : "tuv",
		'9' : "wxyz",
	}
	ans := new([]string)
	if len(digits) == 0 {
		return *ans
	}

	for i := 0; i < len(value[rune(digits[0])]); i ++ {
		*ans = append(*ans, string(value[rune(digits[0])][i]))
	}

	digits = digits[1:]
	fmt.Println(ans)

	for _, str := range digits {
		temp := ans
		tempAns := new([]string)
		for i := 0; i < len(value[str]); i ++ {
			for _, tempValue := range *temp{
				*tempAns = append(*tempAns, tempValue + string(value[str][i]))
			}
		}
		ans = tempAns
		fmt.Println(ans)
	}

	return *ans

}