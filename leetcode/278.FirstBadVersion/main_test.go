package main

import "testing"

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 5,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
