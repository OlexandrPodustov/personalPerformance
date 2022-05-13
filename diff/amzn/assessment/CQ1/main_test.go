package main

import "testing"

func Test_getMinimumDays(t *testing.T) {
	type args struct {
		parcels []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: ",",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinimumDays(tt.args.parcels); got != tt.want {
				t.Errorf("getMinimumDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
