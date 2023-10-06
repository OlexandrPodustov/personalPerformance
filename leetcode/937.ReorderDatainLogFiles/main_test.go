package main

import (
	"reflect"
	"testing"
)

func Test_reorderLogFiles(t *testing.T) {
	type args struct {
		logs []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				logs: []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"},
			},
			want: []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"},
		},
		{
			name: "2",
			args: args{
				logs: []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"},
			},
			want: []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
		},
		{
			name: "3",
			args: args{
				logs: []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo", "a2 act car"},
			},
			want: []string{"a2 act car", "g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
		},
		{
			name: "4",
			args: args{
				logs: []string{"j mo", "5 m w", "g 07", "o 2 0", "t q h"},
			},
			want: []string{"5 m w", "j mo", "t q h", "g 07", "o 2 0"},
		},
		{
			name: "5",
			args: args{
				logs: []string{"7 96", "d 0 5", "r 439", "1 bw", "6 dkf"},
			},
			want: []string{"1 bw", "6 dkf", "7 96", "d 0 5", "r 439"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderLogFiles(tt.args.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderLogFiles() = %v\nwant\t\t\t%v", got, tt.want)
			}
		})
	}
}
