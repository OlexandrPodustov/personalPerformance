package main

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {
		// 	name: "1",
		// 	args: args{sentence: "The quick brown fox jumped over a lazy dog"},
		// 	want: "dog lazy a over jumped fox brown quick The",
		// },

		{
			name: "2",
			args: args{sentence: "We love Go"},
			want: "Go love We",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.sentence); got != tt.want {
				t.Errorf("reverseWords() = %q, want %v", got, tt.want)
			}
		})
	}
}
