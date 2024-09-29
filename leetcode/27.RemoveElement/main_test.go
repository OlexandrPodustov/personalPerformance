package main

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantData []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{3, 3, 3, 3, 2},
				val:  3,
			},
			want:     1,
			wantData: []int{2},
		},
		{
			name: "2",
			args: args{
				nums: []int{3, 3, 3, 2, 2},
				val:  3,
			},
			want:     2,
			wantData: []int{2, 2},
		},

		{
			name: "3",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want:     2,
			wantData: []int{2, 2},
		},

		{
			name: "4",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want:     5,
			wantData: []int{0, 1, 3, 0, 4},
		},
		{
			name: "5",
			args: args{
				nums: []int{2},
				val:  3,
			},
			want:     1,
			wantData: []int{2},
		},

		{
			name: "6",
			args: args{
				nums: []int{3, 3},
				val:  5,
			},
			want:     2,
			wantData: []int{3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			if got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.nums[:got], tt.wantData) {
				t.Errorf("removeElement() = %v, want %v", tt.args.nums[:got], tt.wantData)
			}
		})
	}
}
