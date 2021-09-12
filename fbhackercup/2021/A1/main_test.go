package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		// {
		// 	name: "AAAEEEUUUHCCKRP",
		// 	args: args{
		// 		input: "AAAEEEUUUHCCKRP",
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "HAAACKEEERCUUUP",
		// 	args: args{
		// 		input: "HAAACKEEERCUUUP",
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "OOEUEUIUUOOAUEIOIEEUIAIUAUOOAUUIIEUUOAAIOAEAUEIAOAOOAOOAAIEOAEOEIOIAAEAAIOEEOIIIEEIOIUEIEEEAAEEEEOAA",
		// 	args: args{
		// 		input: "OOEUEUIUUOOAUEIOIEEUIAIUAUOOAUUIIEUUOAAIOAEAUEIAOAOOAOOAAIEOAEOEIOIAAEAAIOEEOIIIEEIOIUEIEEEAAEEEEOAA",
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "QPQWJXRJJXBTKKGBKVXNSCQBHZTSFZRYCDZFYQJQWHWHYJVDRXSGWRLJNTHBXYBRBTVXBBPPCXRBFVXVNDQQTHHBKXZDPQZGSHWF",
		// 	args: args{
		// 		input: "QPQWJXRJJXBTKKGBKVXNSCQBHZTSFZRYCDZFYQJQWHWHYJVDRXSGWRLJNTHBXYBRBTVXBBPPCXRBFVXVNDQQTHHBKXZDPQZGSHWF",
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "EDHHDGFGCFHBHCAFGHHCBDHHECAEGGAGAGDGEGGFDGBCDFDCBHFGBDBCGHEACCGFAGEFBFGECFGFCGGBEDGDBFCEHEDHEHHFGCEB",
		// 	args: args{
		// 		input: "EDHHDGFGCFHBHCAFGHHCBDHHECAEGGAGAGDGEGGFDGBCDFDCBHFGBDBCGHEACCGFAGEFBFGECFGFCGGBEDGDBFCEHEDHEHHFGCEB",
		// 	},
		// 	want: 2,
		// },

		{
			name: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			args: args{
				input: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			},
			want: 29,
		},
		{
			name: "FBHCA",
			args: args{
				input: "FBHCA",
			},
			want: 4,
		},

		{
			name: "ABC",
			args: args{
				input: "ABC",
			},
			want: 2,
		},
		{
			name: "F",
			args: args{
				input: "F",
			},
			want: 0,
		},
		{
			name: "BANANA",
			args: args{
				input: "BANANA",
			},
			want: 3,
		},
		{
			name: "BANANAO",
			args: args{
				input: "BANANAO",
			},
			want: 5,
		},

		{
			name: "FOXEN",
			args: args{
				input: "FOXEN",
			},
			want: 5,
		},
		{
			name: "CONSISTENCY",
			args: args{
				input: "CONSISTENCY",
			},
			want: 12,
		},

		{
			name: "AO mfl none",
			args: args{
				input: "AO",
			},
			want: 2,
		},

		{
			name: "FBHC mfl none",
			args: args{
				input: "FBHC",
			},
			want: 4,
		},

		{
			name: "AO mfl none",
			args: args{
				input: "AOVH",
			},
			want: 4,
		},

		{
			name: "AAAEEEUUUOOOHCCKRP",
			args: args{
				input: "AAAEEEUUUOOOHCCKRP",
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.input); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
