package solution_test

import "testing"
import "strconv"
import "fmt"

func getHint(secret string, guess string) string {
	secretNum := make([]int, 10)
	guessNum := make([]int, 10)
	sameQuantity := 0
	for i := range secret {
		if secret[i] == guess[i] {
			sameQuantity++
		} else {
			secretNum[secret[i]-'0']++
			guessNum[guess[i]-'0']++
		}
	}
	fmt.Printf("%T_%d", secret[0]-'0', int(secret[0]))
	num := 0
	for i := 0; i < 10; i++ {
		if secretNum[i] > guessNum[i] {
			num += guessNum[i]
		} else {
			num += secretNum[i]
		}

	}
	return strconv.Itoa(sameQuantity) + "A" + strconv.Itoa(num) + "B"

}

func TestFunc(t *testing.T) {
	t.Log(getHint("1123", "0111"))

}
