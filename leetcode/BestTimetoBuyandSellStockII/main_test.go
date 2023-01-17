package main

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			// Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
			// Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
			// Total profit is 4 + 3 = 7.
			name: "1",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 7,
		},
		{
			name: "2",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
		{
			name: "4",
			args: args{
				prices: []int{1, 4, 3, 4, 1},
			},
			want: 4,
		},
		{
			name: "3",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
