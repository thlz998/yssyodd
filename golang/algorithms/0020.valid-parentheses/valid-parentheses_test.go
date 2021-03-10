package problem0020

import (
"testing"

"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	s   string
	ans bool
}{
	{
		"()",
		true,
	},
	{
		"()[]{}",
		true,
	},
	{
		"(]",
		false,
	},
	{
		"([)]",
		false,
	},
	{
		"{[]}",
		true,
	},
}

func Test_reverseWords(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, isValid(tc.s), "输入:%v", tc)
	}
}
