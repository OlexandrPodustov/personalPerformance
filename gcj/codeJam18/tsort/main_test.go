package main

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
)

type univ struct {
	testCaseNumber uint
	dataSetSize    uint
	rawDataSet     []int
	expectedResult string
}

var tt = []univ{
	{
		1,
		5,
		[]int{5, 6, 8, 4, 3},
		"OK",
	},
	{
		2,
		3,
		[]int{8, 9, 7},
		"2",
	},
}

func TestRun(t *testing.T) {
	it := is.New(t)
	for _, test := range tt {
		t.Log(test)
		var b []byte
		for _, v := range test.rawDataSet {
			b = append(b, byte(v))
		}

		res := tSort(b)

		actualResult := checkAscending(res)

		it.Equal(test.expectedResult, actualResult) // test.expectedResult != actualResult
	}

}

func BenchmarkBbbbb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
