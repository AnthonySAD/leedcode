package solution_test

import "testing"
import "fmt"

func canCompleteCircuitOld(gas []int, cost []int) int {
	if len(gas) == 1 && gas[0] >= cost[0] {
		return 0
	}

	for i := 0; i < len(gas); i++ {
		curGas := gas[i] - cost[i]
		if curGas <= 0 {
			continue
		}

		j := i + 1
		end := i
		if j >= len(gas) {
			j = 0
		}

		for {

			curGas += gas[j] - cost[j]
			if curGas < 0 {
				break
			}
			j++
			if j == len(gas) {
				j = 0
			}
			if j == end {
				return i
			}
		}
	}

	return -1
}

func canCompleteCircuitOld2(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		gas[i] -= cost[i]
	}
	current := gas[0]
	i, j := 0, 0
	for i < len(gas) {
		if current >= 0 {
			//结束条件
			if (i == 0 && j == len(gas)-1) || j == i-1 {
				return i
			}
			j++
			if j == len(gas) {
				j = 0
			}
			current += gas[j]

		} else {
			if i == j {
				i++
				j = i
				//一轮循环玩没找到答案
				if i == len(gas) {
					return -1
				}
				current = gas[i]
			} else {
				current -= gas[i]
				i++
			}

		}

		//一轮循环玩没找到答案
		if i == len(gas) {
			return -1
		}
	}

	if len(gas) == 1 && gas[0] >= cost[0] {
		return 0
	}

	for i := 0; i < len(gas); i++ {
		curGas := gas[i] - cost[i]
		if curGas <= 0 {
			continue
		}

		j := i + 1
		end := i
		if j >= len(gas) {
			j = 0
		}

		for {

			curGas += gas[j] - cost[j]
			if curGas < 0 {
				break
			}
			j++
			if j == len(gas) {
				j = 0
			}
			if j == end {
				return i
			}
		}
	}

	return -1
}
func canCompleteCircuit(gas []int, cost []int) int {

	n := len(gas)
	for i := 0; i < n; i++ {
		gas[i] -= cost[i]
	}
	curGas, i, j := 0, 0, 0

	//还是用双循环的结构写清晰一点
	for i < n {
		curGas = gas[i]
		j = i + 1
		for {
			if curGas < 0 {
				break
			}

			if i == j%n {
				return i
			}
			curGas += gas[j%n]
			j++

		}
		i = j

	}

	return -1
}
func TestFunc(t *testing.T) {
	t.Log(canCompleteCircuit([]int{1, 1, 3}, []int{2, 2, 1}))
}
