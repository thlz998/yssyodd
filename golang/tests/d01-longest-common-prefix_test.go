package testApp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thlz998/yssyodd/golang/app"
)

type question struct {
	p []string
	a string
}

var qs = []question{
	{
		p: []string{"flower", "flow", "flight"},
		a: "fl",
	},
	// {
	// 	p: []string{"dog", "racecar", "car"},
	// 	a: "",
	// },
	// {
	// 	p: []string{},
	// 	a: "",
	// },
	// {
	// 	p: []string{"1234", "123", "12343", "123", "12356"},
	// 	a: "123",
	// },
	// {
	// 	p: []string{"1234", "123", "12343", "123", "12356"},
	// 	a: "123",
	// },
	// {
	// 	p: []string{"1234"},
	// 	a: "1234",
	// },
}

func Test_LongestCommonPrefix(t *testing.T) {
	ast := assert.New(t)
	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a, app.LongestCommonPrefix(p), "输入:%v", p)
	}
}
