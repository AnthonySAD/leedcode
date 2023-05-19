package solution_test

import (
	"testing"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	tmp := make([]int, numCourses)
	for _, v := range prerequisites {
		tmp[v[0]]++
	}
	ok := []int{}

	for i, v := range tmp {
		if v == 0 {
			ok = append(ok, i)
		}
	}

	for i := 0; i < len(ok); i++ {
		for _, v := range prerequisites {
			if v[1] == ok[i] {
				tmp[v[0]]--
				if tmp[v[0]] == 0 {
					ok = append(ok, v[0])
				}
			}
		}
	}
	return len(ok) == numCourses
}

func TestFunc(t *testing.T) {
	t.Log(canFinish(2, [][]int{[]int{1, 0}, []int{0, 1}}))

	// ok := []int{1}
	// fmt.Println(1)
}
