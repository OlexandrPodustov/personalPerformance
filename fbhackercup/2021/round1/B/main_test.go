package main

import (
	"fmt"
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		nRows     int
		mCols     int
		aTopLeft  int
		bTopRight int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {
		// 	name: "1",
		// 	args: args{
		// 		nRows:     2,
		// 		mCols:     2,
		// 		aTopLeft:  999,
		// 		bTopRight: 999,
		// 	},
		// 	want: fmt.Sprintf("%v%v%v", "Possible\n", "333 333\n", "333 333"),
		// },
		// {
		// 	name: "3", //  4 3 6 6
		// 	args: args{
		// 		nRows:     4,
		// 		mCols:     3,
		// 		aTopLeft:  6,
		// 		bTopRight: 6,
		// 	},
		// 	want: fmt.Sprintf("%v%v%v%v%v", "Possible\n", "1 1 1\n", "1 2 1\n", "1 2 1\n", "1 2 1"),
		// },
		// {
		// 	name: "4",
		// 	args: args{
		// 		nRows:     50,
		// 		mCols:     50,
		// 		aTopLeft:  1,
		// 		bTopRight: 1,
		// 	},
		// 	want: `Impossible`,
		// },

		{
			name: "2",
			args: args{
				nRows:     2,
				mCols:     3,
				aTopLeft:  12,
				bTopRight: 11,
			},
			want: fmt.Sprintf("%v%v%v", "Possible\n", "2 2 2\n", "5 9 6"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.nRows, tt.args.mCols, tt.args.aTopLeft, tt.args.bTopRight); got != tt.want {
				t.Errorf("solution() = \n%q\n%q", got, tt.want)
				t.FailNow()
			}
		})
	}
}
