package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int64
	fmt.Scan(&n)

	factors := make(map[int64]int)
	factor(n, factors)

	var res int64 = 1
	for _, d := range factors {
		res *= int64(d + 1)
	}
	fmt.Println(res)
}

func factor(n int64, factors map[int64]int) {
	if n < 2 {
		return
	}
	if isPrime(n) {
		factors[n]++
		return
	}
	d := pollardRho(n)
	factor(d, factors)
	factor(n/d, factors)
}

func isPrime(n int64) bool {
	if n == 2 || n == 3 || n == 5 || n == 7 {
		return true
	}
	if n < 2 || n%2 == 0 || n%3 == 0 {
		return false
	}

	d := n - 1
	s := int64(0)
	for d%2 == 0 {
		d /= 2
		s++
	}
	for _, a := range []int64{2, 3, 5, 7, 11} {
		if a >= n {
			continue
		}
		x := powmod(a, d, n)
		if x == 1 || x == n-1 {
			continue
		}
		ok := false
		for r := int64(1); r < s; r++ {
			x = mulmod(x, x, n)
			if x == n-1 {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}
	return true
}

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

func powmod(a, b, m int64) int64 {
	res := int64(1)
	a %= m
	for b > 0 {
		if b&1 == 1 {
			res = mulmod(res, a, m)
		}
		a = mulmod(a, a, m)
		b >>= 1
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
	if n <= 1 {
		return n
	}
	if n%2 == 0 {
		return 2
	}
	if isPrime(n) {
		return n
	}
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
