package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findRepeatedSequences(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				s: "ATATATATATATATAT",
				k: 6,
			},
			want: []string{"TATATA", "ATATAT"},
		},

		{
			name: "AAAAACCCCCAAAAACCCCCC",
			args: args{
				s: "AAAAACCCCCAAAAACCCCCC",
				k: 8,
			},
			want: []string{"AAAAACCC", "AAAACCCC", "AAACCCCC"},
		},
		{
			name: "GGGGGGGGGGGGGGGGGGGGGGGGG",
			args: args{
				s: "GGGGGGGGGGGGGGGGGGGGGGGGG",
				k: 12,
			},
			want: []string{"GGGGGGGGGGGG"},
		},
		{
			name: "TTTTTCCCCCCCTTTTTTCCCCCCCTTTTTTT",
			args: args{
				s: "TTTTTCCCCCCCTTTTTTCCCCCCCTTTTTTT",
				k: 10,
			},
			want: []string{"CCCCCCCTTT", "CCCCCCTTTT", "CCCCCTTTTT", "CCCCTTTTTT", "TCCCCCCCTT", "TTCCCCCCCT", "TTTCCCCCCC", "TTTTCCCCCC", "TTTTTCCCCC"},
		},
		{
			name: "AAAAAACCCCCCCAAAAAAAACCCCCCCTG",
			args: args{
				s: "AAAAAACCCCCCCAAAAAAAACCCCCCCTG",
				k: 10,
			},
			want: []string{"AAAAAACCCC", "AAAAACCCCC", "AAAACCCCCC", "AAACCCCCCC"},
		},

		{
			name: "AAA",
			args: args{
				s: "AAA",
				k: 6,
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findRepeatedSequences(tt.args.s, tt.args.k)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRepeatedSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
