// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1923/C
// https://codeforces.com/problemset/status/1923/problem/C?friends=on
func Test_cf1923C(t *testing.T) {
	testCases := [][2]string{
		{
			`1
5 4
1 2 1 4 5
1 5
4 4
3 4
1 3`,
			`YES
NO
YES
NO`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1923C)
}
