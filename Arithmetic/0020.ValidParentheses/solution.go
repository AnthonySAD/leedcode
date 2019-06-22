package main

func main() {
	s := "()]"
	ans := isValid(s)
	print(ans)
}


func isValid(s string) bool {
	count := len(s)
	temp := []byte{}
	for i := 0; i < count; i ++ {
		if s[i] == '(' {
			temp = append(temp, ')')
		}else if s[i] == '{' {
			temp = append(temp, '}')
		}else if s[i] == '[' {
			temp = append(temp, ']')
		}else if len(temp) > 0 && temp[len(temp) - 1] == s[i]{
			temp = temp[:len(temp) - 1]
		}else{
			return false
		}
	}

	if len(temp) == 0 {
		return true
	}else{
		return false
	}
}