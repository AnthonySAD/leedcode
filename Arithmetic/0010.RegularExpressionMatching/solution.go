package main
import (
	"fmt"
	"strconv"
)

func main(){
	// text := "mississippi"
	// pattern := "mis*is*p*."	
	text := "aa"
	pattern := "a*"
	res := isMatch(text, pattern)
	fmt.Println(res)
}

var count int
func isMatch2(s string, p string) bool {
	count ++
	fmt.Println("第" + strconv.Itoa(count) + "次,s is " + s + " ,p is " + p)
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		}else{
			return false
		}
	}

	baseMatch := len(s) != 0 && (s[0] == p[0] || p[0] == '.')

	if len(p) > 1 && p[1] == '*' {
		return isMatch(s, p[2:]) || (baseMatch && isMatch(s[1:], p))
	}else{
		return  baseMatch && isMatch(s[1:], p[1:])
	}
}

func isMatch3(s string, p string) bool {
	h := len(s)
	w := len(p)
	res := make([][]int8, h + 1)

	for index := 0; index < h + 1; index++ {
		res[index] = make([]int8, w + 1)
	}
	return dp(res, 0, 0, s, p, h, w)
	
}

func dp(res [][]int8, i, j int, s, p string, h, w int) bool{
	count ++
	fmt.Println("第" + strconv.Itoa(count) + "次,i is " + strconv.Itoa(i) + " ,j is " + strconv.Itoa(j))
	if j == w {
		if i == h {
			return true
		}else{
			return false
		}
	}

	if(res[i][j] != 0 ){
		if(res[i][j] == 2){
			return true
		}else{
			return false
		}
	}
	
	baseMatch := (i < h) && (s[i] == p[j] || p[j] == '.')

	var ans bool

	if j + 1 < w && p[j + 1] == '*' {
		ans = dp(res, i, j + 2, s, p, h ,w) || (baseMatch && dp(res, i + 1, j, s, p, h ,w))
	}else{
		ans = baseMatch && dp(res, i + 1, j + 1, s, p, h ,w)
	}

	if ans {
		res[i][j] = 2
	}else{
		res[i][j] = 1
	}

	return ans
}

func isMatch(s, p string) bool {
	h := len(s)
	w := len(p)
	dp := make([][]bool, h + 1)
	for i := 0; i < h + 1; i++ {
		dp[i] = make([]bool, w + 1)
	}

	dp[h][w] = true

	for i := h; i >= 0; i -- {
		for j := w - 1; j >= 0; j -- {
			
			baseMatch := i < h && (s[i] == p[j] || p[j] == '.')
			if j + 1 < w && p[j + 1] == '*' {
				dp[i][j] = dp[i][j + 2] || (baseMatch && dp[i + 1][j])
			}else{
				dp[i][j] = baseMatch && dp[i + 1][j + 1]
			}
			count ++
			fmt.Println("第" + strconv.Itoa(count) + "次,i is " + strconv.Itoa(i) + " ,j is " + strconv.Itoa(j) + " ,result is", dp[i][j])
		}
	}

	return dp[0][0]
}

// boolean[][] dp = new boolean[text.length() + 1][pattern.length() + 1];
//         dp[text.length()][pattern.length()] = true;

//         for (int i = text.length(); i >= 0; i--){
//             for (int j = pattern.length() - 1; j >= 0; j--){
//                 boolean first_match = (i < text.length() &&
//                                        (pattern.charAt(j) == text.charAt(i) ||
//                                         pattern.charAt(j) == '.'));
//                 if (j + 1 < pattern.length() && pattern.charAt(j+1) == '*'){
//                     dp[i][j] = dp[i][j+2] || first_match && dp[i+1][j];
//                 } else {
//                     dp[i][j] = first_match && dp[i+1][j+1];
//                 }
//             }
//         }
//         return dp[0][0];
