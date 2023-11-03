// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1634/F
// https://codeforces.com/problemset/status/1634/problem/F
func TestCF1634F(t *testing.T) {
	testCases := [][2]string{
		{
			`3 5 3
2 2 1
0 0 0
A 1 3
A 1 3
B 1 1
B 2 2
A 3 3`,
			`YES
NO
NO
NO
YES`,
		},
		{
			`5 3 10
2 5 0 3 5
3 5 8 2 5
B 2 3
B 3 4
A 1 2`,
			`NO
NO
YES`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, CF1634F)
}
