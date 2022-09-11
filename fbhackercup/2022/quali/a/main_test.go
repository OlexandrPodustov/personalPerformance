package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		preOwnedClockParts         []int
		preOwnedClockPartsCount    map[int]int
		capable_of_holding_at_most int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				preOwnedClockParts:         []int{1, 2, 2},
				preOwnedClockPartsCount:    map[int]int{1: 1, 2: 2},
				capable_of_holding_at_most: 2,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.preOwnedClockParts, tt.args.preOwnedClockPartsCount, tt.args.capable_of_holding_at_most); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
