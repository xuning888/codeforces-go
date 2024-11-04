// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1793/D
// https://codeforces.com/problemset/status/1793/problem/D?friends=on
func Test_cf1793D(t *testing.T) {
	testCases := [][2]string{
		{
			`3
1 3 2
2 1 3`,
			`2`,
		},
		{
			`7
7 3 6 2 1 5 4
6 7 2 5 3 1 4`,
			`16`,
		},
		{
			`6
1 2 3 4 5 6
6 5 4 3 2 1`,
			`11`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1793D)
}
