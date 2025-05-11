package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type SegmentTree struct {
	t []int
	n int
}

func New(n int) SegmentTree {
	size := 1
	for size < n {
		size <<= 1
	}
	return SegmentTree{
		t: make([]int, 2*size),
		n: size,
	}
}

//	             t[1]
//				/    \
//			 t[2]    t[3]
//			 /  \     /  \
//		  t[4] t[5] t[6] t[7]
func (st *SegmentTree) Modify(i, x int) {
	i += st.n - 1
	st.t[i] = x
	for i > 1 {
		st.t[i>>1] = st.t[i] + st.t[i^1] // t[3] = t[6] + t[7]
		i >>= 1
	}
}

func (st *SegmentTree) Query(l, r int) int {
	l += st.n - 1
	r += st.n - 1
	res := 0
	for l < r {
		if l&1 == 1 {
			res += st.t[l]
			l++
		}
		if r&1 == 1 {
			r--
			res += st.t[r]
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func main() {
	fin, err := os.Open("sum.in")
	if err != nil {
		log.Fatalf("cannot open input file: %v", err)
	}
	defer fin.Close()

	fout, err := os.Create("sum.out")
	if err != nil {
		log.Fatalf("cannot create output file: %v", err)
	}
	defer fout.Close()

	in := bufio.NewReader(fin)
	out := bufio.NewWriter(fout)
	defer out.Flush()

	line, _ := in.ReadString('\n')
	parts := strings.Fields(line)
	n, _ := strconv.Atoi(parts[0])
	k, _ := strconv.Atoi(parts[1])

	st := New(n)

	for range k {
		line, _ := in.ReadString('\n')
		parts := strings.Fields(line)
		if parts[0] == "A" {
			idx, _ := strconv.Atoi(parts[1])
			val, _ := strconv.Atoi(parts[2])
			st.Modify(idx-1, val)
		} else if parts[0] == "Q" {
			l, _ := strconv.Atoi(parts[1])
			r, _ := strconv.Atoi(parts[2])
			sum := st.Query(l-1, r)
			fmt.Fprintln(out, sum)
		}
	}
}
