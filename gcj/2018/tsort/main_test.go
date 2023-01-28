package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	type univ struct {
		rawDataSet     []int
		expectedResult string
	}

	tt := []univ{
		{
			rawDataSet:     []int{5, 6, 8, 4, 3},
			expectedResult: "OK",
		},
		{
			rawDataSet:     []int{8, 9, 7},
			expectedResult: "2",
		},
	}

	for _, test := range tt {
		t.Log(test)
		res := tSort(test.rawDataSet)
		actualResult := checkAscending(res)
		if test.expectedResult != actualResult {
			t.Errorf("actualResult: %s", actualResult)
			t.FailNow()
		}
	}
}
