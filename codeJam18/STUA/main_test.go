package main

import (
	"fmt"
	"git.4tree.de/UWF/uwf-vendor/github.com/matryer/is"
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
	it := is.New(t)
	for _, test := range tt {
		t.Log(test)

		ff := []byte(test.ss)
		actualResult := calc(test.maxWithstand, &ff)
		it.Equal(test.expectedResult, actualResult) // test.expectedResult != actualResult
	}

}

func BenchmarkBbbbb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}
