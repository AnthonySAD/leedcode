package main

func main() {
	str := "MCMXCIV"
	ans := romanToInt(str)
	print(ans)
}

func romanToInt(s string) int {
	ans := 0
	length := len(s)
	tempV := 0
	tempV2 := 0
	value := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	i := 0
	for {
		tempV = value[s[i]]
		i++
		if i >= length {
			return ans + tempV
		}

		tempV2 = value[s[i]]
		if tempV2 > tempV {
			ans -= tempV
		} else {
			ans += tempV
		}
	}
}
