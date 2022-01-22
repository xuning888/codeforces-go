// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][3]string{
		{
			`["LockingTree", "lock", "unlock", "unlock", "lock", "upgrade", "lock"]`,
			`[[[-1, 0, 0, 1, 1, 2, 2]], [2, 2], [2, 3], [2, 2], [4, 5], [0, 1], [0, 1]]`,
			`[null, true, false, true, true, true, false]`,
		},
		{
			`["LockingTree","upgrade","upgrade","upgrade","upgrade","lock","upgrade","lock","upgrade","lock","lock","lock","upgrade","upgrade","upgrade","upgrade","lock","upgrade","lock","upgrade","unlock"]`,
`[[[-1,0,8,0,7,4,2,3,3,1]],[8,39],[5,28],[6,33],[9,24],[5,22],[1,3],[5,20],[0,38],[5,14],[6,34],[6,28],[3,23],[4,45],[8,7],[2,18],[3,35],[2,16],[3,21],[1,41],[5,22]]`,
			`[null,false,false,false,false,true,false,false,true,true,true,false,false,false,false,false,true,false,false,false,false]`,
		},
		//{
		//	`["LockingTree","lock","upgrade","lock","upgrade"  ,"lock",   "lock","lock","upgrade","upgrade","upgrade","upgrade","lock","upgrade","lock","upgrade","unlock"]`,
		//	`[[[-1,0,8,0,7,4,2,3,3,1]],,[5,22],[1,3],[5,20],[0,38],  [5,14],  [6,34],[6,28],[3,23],[4,45],[8,7],[2,18],[3,35],[2,16],[3,21],[1,41],[5,22]]`,
		//	`[null,,true,false,false,true,true,true,false,false,false,false,false,true,false,false,false,false]`,
		//},
	}
	targetCaseNum := -1
	if err := testutil.RunLeetCodeClassWithExamples(t, Constructor, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-60/problems/operations-on-tree/
