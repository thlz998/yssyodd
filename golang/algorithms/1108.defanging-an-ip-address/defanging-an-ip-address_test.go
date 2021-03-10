package problem1108

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans string
}{
	{
		"1.1.1.1",
		"1[.]1[.]1[.]1",
	},
	{
		"2.1.3.1",
		"2[.]1[.]3[.]1",
	},
}

func Test_reverseWords(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, defangIPaddr(tc.s), "输入:%v", tc)
	}
}
