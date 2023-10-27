package main

import "testing"

func Test_canAttendMeetings(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[0,30],[5,10],[15,20]",
			args: args{
				intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			},
			want: false,
		},
		{
			name: "[[7,10],[2,4]]",
			args: args{
				intervals: [][]int{{7, 10}, {2, 4}},
			},
			want: true,
		},
		{
			name: "{0, 30}, {40, 50}, {29, 35}",
			args: args{
				intervals: [][]int{{0, 30}, {40, 50}, {29, 35}},
			},
			want: false,
		},
		{
			name: "{13, 15}, {1, 13}",
			args: args{
				intervals: [][]int{{13, 15}, {1, 13}},
			},
			want: true,
		},
		{
			name: "[[9,14],[5,6],[3,5],[12,19]]",
			args: args{
				intervals: [][]int{{9, 14}, {5, 6}, {3, 5}, {12, 19}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := canAttendMeetings(tt.args.intervals); got != tt.want {
				t.Errorf("canAttendMeetings(%v) = %v, want %v", tt.args.intervals, got, tt.want)
			}
		})
	}
}
