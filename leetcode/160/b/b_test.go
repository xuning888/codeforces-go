// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	exampleIns := [][]string{{`2`, `3`}, {`3`, `2`}}
	exampleOuts := [][]string{{`[3,2,0,1]`}, {`[2,6,7,5,4,0,1,3]`}}
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	if err := testutil.RunLeetCodeFunc(t, circularPermutation, exampleIns, exampleOuts); err != nil {
		t.Fatal(err)
	}
}
