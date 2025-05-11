package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const N = 100000

type Node struct {
	min, max int
}

type SegmentTree struct {
	t []Node
	n int
}

func New(a []int) SegmentTree {
	n := len(a)
	size := 1
	for size < n {
		size <<= 1
	}
	t := make([]Node, 2*size)
	st := SegmentTree{t: t, n: size}

	for i := range n {
		st.t[i+st.n] = Node{min: a[i], max: a[i]}
	}

	for i := st.n - 1; i > 0; i-- {
		st.t[i] = Node{
			min: min(st.t[i<<1].min, st.t[i<<1|1].min),
			max: max(st.t[i<<1].max, st.t[i<<1|1].max),
		}
	}
	return st
}

func (st *SegmentTree) Modify(i, x int) {
	i += st.n - 1
	st.t[i] = Node{min: x, max: x}
	for i > 1 {
		i >>= 1
		st.t[i] = Node{
			min: min(st.t[i<<1].min, st.t[i<<1|1].min),
			max: max(st.t[i<<1].max, st.t[i<<1|1].max),
		}
	}
}

func (st *SegmentTree) Query(l, r int) int {
	l += st.n - 1
	r += st.n - 1
	resMin := N + 1
	resMax := -N - 1

	for l <= r {
		if l&1 == 1 {
			resMin = min(resMin, st.t[l].min)
			resMax = max(resMax, st.t[l].max)
			l++
		}
		if r&1 == 0 {
			resMin = min(resMin, st.t[r].min)
			resMax = max(resMax, st.t[r].max)
			r--
		}
		l >>= 1
		r >>= 1
	}
	return resMax - resMin
}

func main() {
	fin, err := os.Open("rvq.in")
	if err != nil {
		log.Fatalf("cannot open input file: %v", err)
	}
	defer fin.Close()

	fout, err := os.Create("rvq.out")
	if err != nil {
		log.Fatalf("cannot create output file: %v", err)
	}
	defer fout.Close()

	in := bufio.NewReader(fin)
	out := bufio.NewWriter(fout)
	defer out.Flush()

	line, _ := in.ReadString('\n')
	k, _ := strconv.Atoi(strings.TrimSpace(line))

	a := make([]int, N)
	for i := 1; i <= N; i++ {
		a[i-1] = (i*i)%12345 + (i*i*i)%23456
	}

	st := New(a)

	for range k {
		line, _ := in.ReadString('\n')
		parts := strings.Fields(line)
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])

		if x > 0 {
			// max - min on [x, y]
			result := st.Query(x, y)
			fmt.Fprintln(out, result)
		} else {
			// a[abs(x)] = y
			st.Modify(-x, y)
		}
	}
}
