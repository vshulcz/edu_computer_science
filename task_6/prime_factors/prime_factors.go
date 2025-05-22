package main

import (
	"fmt"
	"math/rand"
)

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
			x = (x*x + c) % n
			y = (y*y + c) % n
			y = (y*y + c) % n
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
