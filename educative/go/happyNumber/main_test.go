package main

import "testing"

func Test_happyNumber(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				num: 1,
			},
			want: true,
		},
		{
			name: "28",
			args: args{
				num: 28,
			},
			want: true,
		},
		{
			name: "2147483646",
			args: args{
				num: 2147483646,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := happyNumber(tt.args.num); got != tt.want {
				t.Errorf("happyNumber() = %v, want %v", got, tt.want)
				t.FailNow()
			}
		})
	}
}
