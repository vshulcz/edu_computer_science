package main

import (
	"math/big"
	"reflect"
	"runtime"
	"testing"
)

var N, _ = new(big.Int).SetString("19590340644999083431262508198206381046123972390589368223882605328968666316379870661851951648789482321596229559115436019149189529725215266728292282990852649023362731392404017939142010958261393634959471483757196721672243410067118516227661133135192488848989914892157188308679896875137439519338903968094905549750386407106033836586660683539201011635917900039904495065203299749542985993134669814805318474080581207891125910", 10)

const M = 10000

func Benchmark3Primes(b *testing.B) {
	runtime.GOMAXPROCS(1)
	primes := []int{2, 3, 5}
	gen := wheelGenerator(primes)
	b.ResetTimer()
	for range M {
		factor(N, primes, gen)
	}
}

func Benchmark4Primes(b *testing.B) {
	runtime.GOMAXPROCS(1)
	primes := []int{2, 3, 5, 7}
	gen := wheelGenerator(primes)
	b.ResetTimer()
	for range M {
		factor(N, primes, gen)
	}
}

func Benchmark5Primes(b *testing.B) {
	runtime.GOMAXPROCS(1)
	primes := []int{2, 3, 5, 7, 11}
	gen := wheelGenerator(primes)
	b.ResetTimer()
	for range M {
		factor(N, primes, gen)
	}
}

func Benchmark6Primes(b *testing.B) {
	runtime.GOMAXPROCS(1)
	primes := []int{2, 3, 5, 7, 11, 13}
	gen := wheelGenerator(primes)
	b.ResetTimer()
	for range M {
		factor(N, primes, gen)
	}
}

func Benchmark7Primes(b *testing.B) {
	runtime.GOMAXPROCS(1)
	primes := []int{2, 3, 5, 7, 11, 13, 17}
	gen := wheelGenerator(primes)
	b.ResetTimer()
	for range M {
		factor(N, primes, gen)
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
