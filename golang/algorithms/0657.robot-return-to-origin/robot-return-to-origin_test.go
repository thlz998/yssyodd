package problem0657

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	s   string
	ans bool
}{
	{
		"UD",
		true,
	},
	{
		"LL",
		false,
	},
	{
		"RLUURDDDLU",
		true,
	},
}

func Test_reverseWords(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, judgeCircle(tc.s), "输入:%v", tc)
	}
}
