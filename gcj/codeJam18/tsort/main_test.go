package main

import (
	"testing"
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
	for _, test := range tt {
		t.Log(test)

		res := tSort(test.rawDataSet)

		actualResult := checkAscending(res)

		if test.expectedResult != actualResult {
			t.Errorf("actualResult: %s", actualResult)
			t.Fail()
		}

	}

}
