package main

import (
	"math/big"
	"reflect"
	"testing"
)

var N *big.Int = big.NewInt(21)

func Benchmark4Primes(b *testing.B) {
	primes := []int{2, 3, 5, 7}
	for i := 0; i < b.N; i++ {
		factor(N, primes)
	}
}

func Benchmark5Primes(b *testing.B) {
	primes := []int{2, 3, 5, 7, 11}
	for i := 0; i < b.N; i++ {
		factor(N, primes)
	}
}

func Benchmark6Primes(b *testing.B) {
	primes := []int{2, 3, 5, 7, 11, 13}
	for i := 0; i < b.N; i++ {
		factor(N, primes)
	}
}

func Benchmark7Primes(b *testing.B) {
	primes := []int{2, 3, 5, 7, 11, 13, 17}
	for i := 0; i < b.N; i++ {
		factor(N, primes)
	}
}

func extractChannel(ch <-chan *big.Int, n int) []*big.Int {
	result := make([]*big.Int, n)
	for i := range n {
		result[i] = <-ch
	}
	return result
}

func Test_wheelGenerator(t *testing.T) {
	tests := []struct {
		name   string
		primes []int
		want   []*big.Int
	}{
		{
			name:   "Case 1",
			primes: []int{2, 3},
			want: []*big.Int{
				big.NewInt(1), big.NewInt(5), big.NewInt(7), big.NewInt(11), big.NewInt(13),
				big.NewInt(17), big.NewInt(19), big.NewInt(23), big.NewInt(25), big.NewInt(29),
			},
		},
		{
			name:   "Case 2",
			primes: []int{2, 3, 5},
			want: []*big.Int{
				big.NewInt(1), big.NewInt(7), big.NewInt(11), big.NewInt(13), big.NewInt(17),
				big.NewInt(19), big.NewInt(23), big.NewInt(29), big.NewInt(31), big.NewInt(37),
			},
		},
		{
			name:   "Case 3",
			primes: []int{2, 3, 5, 7},
			want: []*big.Int{
				big.NewInt(1), big.NewInt(11), big.NewInt(13), big.NewInt(17), big.NewInt(19),
				big.NewInt(23), big.NewInt(29), big.NewInt(31), big.NewInt(37), big.NewInt(41),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractChannel(wheelGenerator(tt.primes), 10)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wheelGenerator(%v) = %v, want %v", tt.primes, got, tt.want)
			}
		})
	}
}
