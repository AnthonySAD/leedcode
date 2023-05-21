package solution

func rotate(matrix [][]int) {
	n := len(matrix)
	//分成4个扇区，只遍历一个扇区，依次旋转4个点即可
	width := (n + 1) / 2
	height := n / 2

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] = matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}
