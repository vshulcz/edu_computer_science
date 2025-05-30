package main

import (
	"testing"
)

func Test_factor(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want int64
	}{
		{
			name: "Case 1",
			n:    12,
			want: 6,
		},
		{
			name: "Case 2",
			n:    999999999999999989,
			want: 2,
		},
		{
			name: "Case 3",
			n:    100000007700000049,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factors := make(map[int64]int)
			factor(tt.n, factors)

			var res int64 = 1
			for _, e := range factors {
				res *= int64(e + 1)
			}
			if res != tt.want {
				t.Errorf("factor() = %v, want %v", res, tt.want)
			}
		})
	}
}
