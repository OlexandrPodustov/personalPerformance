package main

import (
	"reflect"
	"testing"
)

func Test_extractColumns(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				input: []string{
					"OOO",
					"X.X",
					"XXO",
				},
			},
			want: []string{"OXX", "O.X", "OXO"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractColumns(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solution(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				input: []string{"XO", ".."},
			},
			want: "1 1",
		},
		{
			name: "2",
			args: args{
				input: []string{"X.", ".O"},
			},
			want: "1 2",
		},
		{
			name: "3",
			args: args{
				input: []string{
					"...",
					"...",
					"...",
				},
			},
			want: "3 6",
		},
		{
			name: "4",
			args: args{
				input: []string{
					".OX",
					"X..",
					"..O",
				},
			},
			want: "2 2",
		},
		{
			name: "5",
			args: args{
				input: []string{
					"OXO",
					"X.X",
					"OXO",
				},
			},
			want: "1 1",
		},
		{
			name: "6",
			args: args{
				input: []string{
					".XO",
					"O.X",
					"XO.",
				},
			},
			want: "Impossible",
		},
		{
			name: "7",
			args: args{
				input: []string{
					"X...",
					".O.O",
					".XX.",
					"O.XO",
				},
			},
			want: "2 2",
		},
		{
			name: "8",
			args: args{
				input: []string{
					"OX.OO",
					"X.XXX",
					"OXOOO",
					"OXOOO",
					"XXXX.",
				},
			},
			want: "1 2",
		},
		{
			name: "mine1",
			args: args{
				input: []string{"XO", "OO"},
			},
			want: "Impossible",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.input); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
