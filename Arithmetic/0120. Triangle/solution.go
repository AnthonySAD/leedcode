package test_solution

import (
	"math"
	"testing"
)


func minimumTotal(triangle [][]int) int {
    height := len(triangle)
    f := make([][]int, height)

    for i := 0; i < height; i++ {
        f[i] = make([]int, height)
    }
    f[0][0] = triangle[0][0]
    for i := 1; i < height; i++ {
        f[i][0] = f[i - 1][0] + triangle[i][0]
        for j := 1; j < i; j++ {
            f[i][j] = min(f[i - 1][j - 1], f[i - 1][j]) + triangle[i][j]
        }
        f[i][i] = f[i - 1][i - 1] + triangle[i][i]
    }
    ans := math.MaxInt32
    for i := 0; i < height; i++ {
        ans = min(ans, f[height-1][i])
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func minimumTotalGood(triangle [][]int) int {
    n := len(triangle)
    f := make([]int, n)
    f[0] = triangle[0][0]
    for i := 1; i < n; i++ {
        f[i] = f[i - 1] + triangle[i][i]
        for j := i - 1; j > 0; j-- {
            f[j] = min(f[j - 1], f[j]) + triangle[i][j]
        }
        f[0] += triangle[i][0]
    }
    ans := math.MaxInt32
    for i := 0; i < n; i++ {
        ans = min(ans, f[i])
    }
    return ans
}


func TestSo(t *testing.T) {
	triangle := [][]int{[2],[3,4],[6,5,7],[4,1,8,3]}
	t.Log(minimuminimumTotal(triangle))
}