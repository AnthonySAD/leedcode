package main

func main() {
	number := 3999
	ans := intToRoman(number)
	print(ans)
}

func intToRoman(num int) string {
	ans := ""
	value := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	str := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	i := 0
	for num > 0 {
		if num < value[i] {
			i++
		} else {
			ans += str[i]
			num -= value[i]
		}
	}
	return ans
}

func intToRoman1(num int) string {
	ans := ""
	temp := 0
	q := [3]string{"M", "MM", "MMM"}
	b := [9]string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	s := [9]string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	g := [9]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	for num > 0 {
		switch {
		case num >= 1000:
			temp = num / 1000
			num -= temp * 1000
			ans += q[temp-1]
		case num >= 100:
			temp = num / 100
			num -= temp * 100
			ans += b[temp-1]
		case num >= 10:
			temp = num / 10
			num -= temp * 10
			ans += s[temp-1]
		case num > 0:
			ans += g[num-1]
			num = 0
		}
	}
	return ans
}
