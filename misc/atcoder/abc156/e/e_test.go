// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc156/tasks/abc156_e
// 提交：https://atcoder.jp/contests/abc156/submit?taskScreenName=abc156_e
// 对拍：https://atcoder.jp/contests/abc156/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc156_e&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc156/submissions?f.Status=AC&f.Task=abc156_e&orderBy=source_length
func Test_e(t *testing.T) {
	testCases := [][2]string{
		{
			`3 2`,
			`10`,
		},
		{
			`200000 1000000000`,
			`607923868`,
		},
		{
			`15 6`,
			`22583772`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
