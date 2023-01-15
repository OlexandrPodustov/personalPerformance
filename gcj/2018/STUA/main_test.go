package main

import (
	"testing"
)

type univ struct {
	testCaseNumber uint
	maxWithstand   uint
	expectedResult string
	ss             string
}

var tt = []univ{
	{
		1,
		1,
		"1",
		"CS",
	},
	{
		2,
		2,
		"0",
		"CS",
	},
	{
		3,
		1,
		"IMPOSSIBLE",
		"SS",
	},
	{
		4,
		6,
		"2",
		"SCCSSC",
	},
	{
		5,
		2,
		"0",
		"CC",
	},
	{
		6,
		3,
		"5",
		"CSCSS",
	},
	{
		7,
		0,
		"0",
		"CCCCCCCCC",
	},
}

func TestRun(t *testing.T) {
	for _, test := range tt {
		t.Log(test)

		ff := []byte(test.ss)
		actualResult := calc(test.maxWithstand, &ff)
		if test.expectedResult != actualResult {
			t.Fail()
		}
	}
}
