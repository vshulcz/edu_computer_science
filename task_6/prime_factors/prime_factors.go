package main

import (
	"fmt"
	"math/rand"
)

func mulmod(a, b, m int64) int64 {
	var res int64 = 0
	for a != 0 {
		if a&1 == 1 {
			res = (res + b) % m
		}
		a >>= 1
		b = (b << 1) % m
	}
	return res
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func pollardRho(n int64) int64 {
	const iterations = 100000
	for range iterations {
		c := rand.Int63n(n-1) + 1
		x := rand.Int63n(n)
		y := x
		var d int64 = 1
		for d == 1 {
			x = (mulmod(x, x, n) + c) % n
			y = (mulmod(y, y, n) + c) % n
			y = (mulmod(y, y, n) + c) % n
			if x > y {
				d = gcd(x-y, n)
			} else {
				d = gcd(y-x, n)
			}
		}
		if d < n {
			return d
		}
	}
	return n
}

func main() {
	var n int64
	fmt.Scan(&n)

	d := pollardRho(n)
	other := n / d
	if d > other {
		d, other = other, d
	}
	fmt.Println(d, other)
}
