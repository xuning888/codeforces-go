// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/arc100/submit?taskScreenName=arc100_c
func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`2
1 2 3 1`,
			`3
4
5`,
		},
		{
			`3
10 71 84 33 6 47 23 25`,
			`81
94
155
155
155
155
155`,
		},
		{
			`4
75 26 45 72 81 47 97 97 2 2 25 82 84 17 56 32`,
			`101
120
147
156
156
178
194
194
194
194
194
194
194
194
194`,
		},
		// TODO 测试参数的下界和上界
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/arc100/tasks/arc100_c
