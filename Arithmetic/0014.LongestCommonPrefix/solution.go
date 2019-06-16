package main

func main(){
	strs := []string{"aa","a"}
	print(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	count := len(strs)
	pre := ""
	if count == 0 {
		return pre
	}
	length := len(strs[0])
	eachLen := 0
	for _, item := range strs {
		eachLen = len(item)
		if eachLen < length {
			length = eachLen
		}
	}

    for i := 0; i < length; i++ {
		for _, item := range strs {
			if item[i] != strs[0][i] {
				return pre
			}
		}
		pre += string(strs[0][i])
	}

	return pre
}

func longestCommonPrefix1(strs []string) string {
	count := len(strs)
	pre := ""
	if count == 0 {
		return pre
	}
	length := len(strs[0])
    for i := 0; i < length; i++ {
		for j := 1; j < count; j++ {
			if len(strs[j]) - 1 < i || strs[j][i] != strs[0][i] {
				return pre
			}
		}
		pre += string(strs[0][i])
	}

	return pre
}