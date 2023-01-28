package main

import (
	"log"
	"reflect"
	"testing"
)

func TestCalc(t *testing.T) {
	type c struct {
		in   string
		want []int
	}

	testCases := []c{
		// {
		// 	in:   "[ SUM [ SUM a.txt b.txt ] [ SUM b.txt c.txt ] ]",
		// 	want: []int{1, 2, 3, 4, 5},
		// },
		// {
		// 	in:   "[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]",
		// 	want: []int{1, 3, 4},
		// },
	}

	for _, tc := range testCases {
		got, err := calc(tc.in)
		_ = err
		// assert.NoError(t, err)

		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("for input: %v - want: %v, got: %v", tc.in, tc.want, got)
			t.FailNow()
		}
	}
}

func TestSum(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	type c struct {
		in   string
		want []int
	}

	testCases := []c{
		{
			in:   "a.txt b.txt",
			want: []int{1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		got, err := sum(tc.in)
		_ = err

		// assert.NoError(t, err)

		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("for input: %v - want: %v, got: %v", tc.in, tc.want, got)
			t.FailNow()
		}
	}
}

func TestDif(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	type c struct {
		in   string
		want []int
	}

	testCases := []c{
		{
			in:   "a.txt b.txt c.txt",
			want: []int{1},
		},
	}

	for _, tc := range testCases {
		got, err := diff(tc.in)
		_ = err

		// assert.NoError(t, err)

		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("for input: %v - want: %v, got: %v", tc.in, tc.want, got)
			t.FailNow()
		}
	}
}

// INT b.txt c.txt
func TestIntersect(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	type c struct {
		in   string
		want []int
	}
	testCases := []c{
		{
			in:   "b.txt c.txt",
			want: []int{3, 4},
		},
	}

	for _, tc := range testCases {
		got, err := intersect(tc.in)
		_ = err

		// assert.NoError(t, err)

		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("for input: %v - want: %v, got: %v", tc.in, tc.want, got)
			t.FailNow()
		}
	}
}

// Sets Calculator
// Write a sets calculator. It should calculate union, intersection and difference of sets of integers for given expression. Grammar of calculator is given:
// expression := “[“ operator sets “]” sets := set | set sets
// set := file | expression
// operator := “SUM” | “INT” | “DIF”
// Each file contains sorted integers, one integer in a line. Meaning of operators:
// SUM - returns union of all sets
// INT - returns intersection of all sets
// DIF - returns difference of first set and the rest ones
// Program should print result on standard output: sorted integers, one integer in a line.
// Solution should include: source code, building script (if building is needed).
// Final program should be able to be run on Linux. Solution should be delivered in tar archive file: “solution.tgz”.

// The task will be assessed taking into consideration the following criteria (sorted by severity):
// 1. compliance with specification and correctness
// 2. good programming practices (clean code)
// 3. clearness and readability of implemented algorithm
// 4. computational complexity

// Example:
// $ cat a.txt 1
// 2
// 3
// $ cat b.txt 2
// 3
// 4
// $ cat c.txt 3
// 4
// 5

// ./scalc [ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]
// 1
// 3
// 4
