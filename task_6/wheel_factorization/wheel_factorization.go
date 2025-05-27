package main

import (
	"fmt"
	"math/big"
)

func wheelGenerator(primes []int) <-chan *big.Int {
	mod := 1
	for _, p := range primes {
		mod *= p
	}

	pattern := []int{}
firstLoop:
	for i := 1; i < mod; i++ {
		for _, p := range primes {
			if i%p == 0 {
				continue firstLoop
			}
		}
		pattern = append(pattern, i)
	}

	ch := make(chan *big.Int)
	go func() {
		base := 0
		for {
			for _, offset := range pattern {
				ch <- big.NewInt(int64(base + offset))
			}
			base += mod
		}
	}()
	return ch
}

func factor(n *big.Int, primes []int, gen <-chan *big.Int) []*big.Int {
	var factors []*big.Int

	for _, p := range primes {
		bp := big.NewInt(int64(p))
		for {
			mod := new(big.Int).Mod(n, bp)
			if mod.Sign() != 0 {
				break
			}
			factors = append(factors, bp)
			n.Div(n, bp)
		}
	}

	one := big.NewInt(1)
	if n.Cmp(one) == 0 {
		return factors
	}

	<-gen
	sqrtN := new(big.Int)
	for c := range gen {
		sqrtN.Mul(c, c)
		if sqrtN.Cmp(n) > 0 {
			break
		}

		for {
			mod := new(big.Int).Mod(n, c)
			if mod.Sign() != 0 {
				break
			}
			factors = append(factors, new(big.Int).Set(c))
			n.Div(n, c)
		}
	}

	if n.Cmp(one) > 0 {
		factors = append(factors, n)
	}

	return factors
}

func main() {
	var s string
	fmt.Scan(&s)

	n := new(big.Int)
	n.SetString(s, 10)

	primes := []int{2, 3, 5}

	gen := wheelGenerator(primes)
	fac := factor(n, primes, gen)

	for i, f := range fac {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(f)
	}
	fmt.Println()
}
