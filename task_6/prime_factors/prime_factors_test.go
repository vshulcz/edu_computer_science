package main

import (
	"reflect"
	"testing"
)

func Test_getPrimeFactors(t *testing.T) {
	tests := []struct {
		name  string
		N     int64
		want  int64
		want1 int64
	}{
		{
			name:  "Case 1",
			N:     15,
			want:  3,
			want1: 5,
		},
		{
			name:  "Case 2",
			N:     77,
			want:  7,
			want1: 11,
		},
		{
			name:  "Case 3",
			N:     221,
			want:  13,
			want1: 17,
		},
		{
			name:  "Case 4",
			N:     899,
			want:  29,
			want1: 31,
		},
		{
			name:  "Case 5",
			N:     1763,
			want:  41,
			want1: 43,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pollardRho(tt.N)
			got1 := tt.N / got
			if got > got1 {
				got, got1 = got1, got
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPrimeFactors() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getPrimeFactors() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
