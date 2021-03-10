package problem0657

import "strings"

func judgeCircle(moves string) bool {
	x := []int{0, 0}
	ar := strings.Split(moves, "")
	for _, a := range ar {
		switch true {
		case a == "U":
			x[0]++
			break
		case a == "D":
			x[0]--
			break
		case a == "L":
			x[1]++
			break
		case a == "R":
			x[1]--
			break
		}
	}

	if x[0] == 0 && x[1] == 0 {
		return true
	}
	return false
}
