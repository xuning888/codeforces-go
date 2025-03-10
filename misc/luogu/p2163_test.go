// Code generated by copypasta/template/generator_test.go
package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://www.luogu.com.cn/problem/P2163
func Test_p2163(t *testing.T) {
	testCases := [][2]string{
		{
			`3 1
0 0 
0 1
1 0
0 0 1 1`,
			`3`,
		},
		{
			`3 1
3 3
1 2
1 0
0 0 2 3`,
			`2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p2163)
}

func TestCompare_p2163(_t *testing.T) {
	//return
	testutil.DebugTLE = 0
	rg := testutil.NewRandGenerator()
	inputGenerator := func() string {
		//return ``
		rg.Clear()
		n := rg.Int(1, 10)
		m := rg.Int(1, 5)
		rg.NewLine()
		for range n {
			rg.Int(0,10)
			rg.Int(0,10)
			rg.NewLine()
		}
		for range m {
			x1 := rg.Int(0,10)
			y1 := rg.Int(0,10)
			rg.Int(x1,10)
			rg.Int(y1,10)
			rg.NewLine()
		}
		return rg.String()
	}

	runBF := func(in io.Reader, out io.Writer) {
		var n, m int
		Fscan(in, &n, &m)
		a := make([]struct{ x, y int }, n)
		for i := range a {
			Fscan(in, &a[i].x, &a[i].y)
		}
		for range m {
		    var x1,y1,x2,y2 int
			Fscan(in, &x1, &y1, &x2, &y2)
			ans := 0
			for _, p := range a {
				if x1<=p.x && p.x <= x2 && y1 <= p.y && p.y <= y2 {
					ans++
				}
			}
			Fprintln(out, ans)
		}
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runBF, p2163)
}
