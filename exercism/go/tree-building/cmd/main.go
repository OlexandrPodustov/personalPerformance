package main

import (
	t "../../tree-building"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

// Define a function Build(records []Record) (*Node, error)
// where Record is a struct containing int fields ID and Parent
// and Node is a struct containing int field ID and []*Node field Children.

var successTestCases = []struct {
	name     string
	input    []t.Record
	expected *t.Node
}{
	{
		name:     "empty input",
		input:    []t.Record{},
		expected: nil,
	},
	{
		name: "one node",
		input: []t.Record{
			{ID: 0},
		},
		expected: &t.Node{
			ID: 0,
		},
	},
	{
		name: "three nodes in order",
		input: []t.Record{
			{ID: 0},
			{ID: 1, Parent: 0},
			{ID: 2, Parent: 0},
		},
		expected: &t.Node{
			ID: 0,
			Children: []*t.Node{
				{ID: 1},
				{ID: 2},
			},
		},
	},
	{
		name: "three nodes in reverse order",
		input: []t.Record{
			{ID: 2, Parent: 0},
			{ID: 1, Parent: 0},
			{ID: 0},
		},
		expected: &t.Node{
			ID: 0,
			Children: []*t.Node{
				{ID: 1},
				{ID: 2},
			},
		},
	},
	{
		name: "more than two children",
		input: []t.Record{
			{ID: 3, Parent: 0},
			{ID: 2, Parent: 0},
			{ID: 1, Parent: 0},
			{ID: 0},
		},
		expected: &t.Node{
			ID: 0,
			Children: []*t.Node{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
		},
	},
	{
		name: "binary tree",
		input: []t.Record{
			{ID: 5, Parent: 1},
			{ID: 3, Parent: 2},
			{ID: 2, Parent: 0},
			{ID: 4, Parent: 1},
			{ID: 1, Parent: 0},
			{ID: 0},
			{ID: 6, Parent: 2},
		},
		expected: &t.Node{
			ID: 0,
			Children: []*t.Node{
				{
					ID: 1,
					Children: []*t.Node{
						{ID: 4},
						{ID: 5},
					},
				},
				{
					ID: 2,
					Children: []*t.Node{
						{ID: 3},
						{ID: 6},
					},
				},
			},
		},
	},
	{
		name: "unbalanced tree",
		input: []t.Record{
			{ID: 5, Parent: 2},
			{ID: 3, Parent: 2},
			{ID: 2, Parent: 0},
			{ID: 4, Parent: 1},
			{ID: 1, Parent: 0},
			{ID: 0},
			{ID: 6, Parent: 2},
		},
		expected: &t.Node{
			ID: 0,
			Children: []*t.Node{
				{
					ID: 1,
					Children: []*t.Node{
						{ID: 4},
					},
				},
				{
					ID: 2,
					Children: []*t.Node{
						{ID: 3},
						{ID: 5},
						{ID: 6},
					},
				},
			},
		},
	},
}

var failureTestCases = []struct {
	name  string
	input []t.Record
}{
	{
		name: "root node has parent",
		input: []t.Record{
			{ID: 0, Parent: 1},
			{ID: 1, Parent: 0},
		},
	},
	{
		name: "no root node",
		input: []t.Record{
			{ID: 1, Parent: 0},
		},
	},
	{
		name: "non-continuous",
		input: []t.Record{
			{ID: 2, Parent: 0},
			{ID: 4, Parent: 2},
			{ID: 1, Parent: 0},
			{ID: 0},
		},
	},
	{
		name: "cycle directly",
		input: []t.Record{
			{ID: 5, Parent: 2},
			{ID: 3, Parent: 2},
			{ID: 2, Parent: 2},
			{ID: 4, Parent: 1},
			{ID: 1, Parent: 0},
			{ID: 0},
			{ID: 6, Parent: 3},
		},
	},
	{
		name: "cycle indirectly",
		input: []t.Record{
			{ID: 5, Parent: 2},
			{ID: 3, Parent: 2},
			{ID: 2, Parent: 6},
			{ID: 4, Parent: 1},
			{ID: 1, Parent: 0},
			{ID: 0},
			{ID: 6, Parent: 3},
		},
	},
	{
		name: "higher id parent of lower id",
		input: []t.Record{
			{ID: 0},
			{ID: 2, Parent: 0},
			{ID: 1, Parent: 2},
		},
	},
}

func (n t.Node) String() string {
	return fmt.Sprintf("%d:%s", n.ID, n.Children)
}

func main() {
	TestMakeTreeSuccess()
}
func TestMakeTreeSuccess(t2 *testing.T) {
	for _, tt := range successTestCases {
		actual, err := t.Build(tt.input)
		if err != nil {
			var _ error = err
			t2.Fatalf("Build for test case %q returned error %q. Error not expected.",
				tt.name, err)
		}
		if !reflect.DeepEqual(actual, tt.expected) {
			t2.Fatalf("Build for test case %q returned %s but was expected to return %s.",
				tt.name, actual, tt.expected)
		}
	}
}

func TestMakeTreeFailure(t2 *testing.T) {
	for _, tt := range failureTestCases {
		actual, err := t.Build(tt.input)
		if err == nil {
			t2.Fatalf("Build for test case %q returned %s but was expected to fail.",
				tt.name, actual)
		}
	}
}

func shuffleRecords(records []t.Record) []t.Record {
	rand := rand.New(rand.NewSource(42))
	newRecords := make([]t.Record, len(records))
	for i, idx := range rand.Perm(len(records)) {
		newRecords[i] = records[idx]
	}
	return newRecords
}

// Binary tree
func makeTwoTreeRecords() []t.Record {
	records := make([]t.Record, 1<<16)
	for i := range records {
		if i == 0 {
			records[i] = t.Record{ID: 0}
		} else {
			records[i] = t.Record{ID: i, Parent: i >> 1}
		}
	}
	return shuffleRecords(records)
}

var twoTreeRecords = makeTwoTreeRecords()

func BenchmarkTwoTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t.Build(twoTreeRecords)
	}
}

// Each node but the root node and leaf nodes has ten children.
func makeTenTreeRecords() []t.Record {
	records := make([]t.Record, 10000)
	for i := range records {
		if i == 0 {
			records[i] = t.Record{ID: 0}
		} else {
			records[i] = t.Record{ID: i, Parent: i / 10}
		}
	}
	return shuffleRecords(records)
}

var tenTreeRecords = makeTenTreeRecords()

func BenchmarkTenTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t.Build(tenTreeRecords)
	}
}

func makeShallowRecords() []t.Record {
	records := make([]t.Record, 10000)
	for i := range records {
		records[i] = t.Record{ID: i, Parent: 0}
	}
	return shuffleRecords(records)
}

var shallowRecords = makeShallowRecords()

func BenchmarkShallowTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t.Build(shallowRecords)
	}
}
