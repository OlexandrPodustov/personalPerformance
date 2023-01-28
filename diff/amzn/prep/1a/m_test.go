package main

import (
	"reflect"
	"testing"
)

func Test_numberOfItems(t *testing.T) {
	type args struct {
		s            string
		startIndices []int32
		endIndices   []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		// {
		// 	name: "1",
		// 	args: args{
		// 		s:            "*|*|",
		// 		startIndices: []int32{1, 1},
		// 		endIndices:   []int32{1, 3},
		// 	},
		// 	want: []int32{1},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfItems(tt.args.s, tt.args.startIndices, tt.args.endIndices); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberOfItems() = %v, want %v", got, tt.want)
				t.FailNow()
			}
		})
	}
}
