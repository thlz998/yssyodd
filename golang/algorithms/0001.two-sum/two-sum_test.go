package problem0001

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one []int
}

type question struct {
	p para
	a ans
}

var qs = []question{
	{
		p: para{
			one: []int{3, 2, 4},
			two: 6,
		},
		a: ans{
			one: []int{1, 2},
		},
	},
	{
		p: para{
			one: []int{3, 2, 4},
			two: 8,
		},
		a: ans{
			one: nil,
		},
	},
	{
		p: para{
			one: []int{3, 2, 6, 5, 6, 7, 0, 55, 13, 234, 342},
			two: 15,
		},
		a: ans{
			one: []int{1, 8},
		},
	},
	{
		p: para{
			one: []int{3, 8, 9, 13, 234, 34, 22, 4, 5, 11},
			two: 18,
		},
		a: ans{
			one: []int{3, 8},
		},
	},
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, twoSum(p.one, p.two), "输入:%v", p)
	}
}

func Test_OK2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, twoSum2(p.one, p.two), "输入:%v", p)
	}
}
