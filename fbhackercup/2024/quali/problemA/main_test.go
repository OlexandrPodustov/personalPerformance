package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		travelers uint64
		timeLimit uint64
		times     []uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				travelers: 4,
				timeLimit: 17,
				times:     []uint64{1, 2, 5, 10},
			},
			want: possible,
		},
		{
			name: "2",
			args: args{
				travelers: 4,
				timeLimit: 4,
				times:     []uint64{1, 2, 5, 10},
			},
			want: impossible,
		},
		{
			name: "3",
			args: args{
				travelers: 2,
				timeLimit: 22,
				times:     []uint64{22, 22},
			},
			want: possible,
		},
		{
			name: "4",
			args: args{
				travelers: 3,
				timeLimit: 1000000000,
				times:     []uint64{1000000000, 1000000000, 1000000000},
			},
			want: impossible,
		},
		{
			name: "5",
			args: args{
				travelers: 1,
				timeLimit: 10,
				times:     []uint64{12},
			},
			want: impossible,
		},

		{
			name: "5",
			args: args{
				travelers: 1,
				timeLimit: 100,
				times:     []uint64{12},
			},
			want: possible,
		},
		{
			name: "6",
			args: args{
				travelers: 1,
				timeLimit: 10,
				times:     []uint64{10},
			},
			want: possible,
		},
		{
			name: "7",
			args: args{
				travelers: 1,
				timeLimit: 10,
				times:     []uint64{9},
			},
			want: possible,
		},
		{
			name: "8",
			args: args{
				travelers: 1,
				timeLimit: 10,
				times:     []uint64{11},
			},
			want: impossible,
		},

		{
			name: "9",
			args: args{
				travelers: 2,
				timeLimit: 10,
				times:     []uint64{10, 15},
			},
			want: possible,
		},
		{
			name: "10",
			args: args{
				travelers: 5,
				timeLimit: 100,
				times:     []uint64{15, 14, 20, 28, 16},
			},
			want: impossible,
		},
		{
			name: "11",
			args: args{
				travelers: 5,
				timeLimit: 100,
				times:     []uint64{16, 15, 20, 28, 17},
			},
			want: impossible,
		},
		{
			name: "12",
			args: args{
				travelers: 5,
				timeLimit: 987654321,
				times:     []uint64{345678912, 456789123, 567891234, 678912345, 789123456},
			},
			want: impossible,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.travelers, tt.args.timeLimit, tt.args.times); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
