package main

import(
	"fmt"
)

func main(){
	num := 3
	ans := generateParenthesis(num)
	fmt.Println(ans)
}

var ans []string = []string{}
func generateParenthesis(n int) []string {
	make(0, 0, n, "")
	return ans
}

func make(left, right, max int, str string){
	if right == max {
		ans = append(ans, str)
	}
	if(left < max){
		make(left + 1, right, max, str + "(")
	}
	if(right < left){
		make(left, right + 1, max, str + ")")
	}
}