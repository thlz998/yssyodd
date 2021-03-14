package app

func RomanToInteger(strs string) int {
	if len(strs) == 0 {
		return 0
	}
	type MapV struct {
		I int8
		V int8
		X int8
		L int8
		C int8
		D int8
		M int8
	}
	v := MapV{
		I: 1,
		V: 5,
		X: 10,
		L: 50,
		C: 100,
		D: 500,
		M: 1000,
	}

	println(v)
	return 0
}
