package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Node struct {
	A   int // BST
	B   int // Heap
	Idx int
}

func BuildTreap(nodes []Node) (parent, left, right []int) {
	n := len(nodes)

	arr := make([]Node, n)
	copy(arr, nodes)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].A < arr[j].A // BST
	})

	parS := make([]int, n)
	ls := make([]int, n)
	rs := make([]int, n)
	for i := range n {
		parS[i], ls[i], rs[i] = -1, -1, -1
	}

	stack := make([]int, 0, n)
	for i := range n {
		last := -1
		for len(stack) > 0 && arr[stack[len(stack)-1]].B > arr[i].B {
			last = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			parS[i] = stack[len(stack)-1]
			rs[stack[len(stack)-1]] = i
		}
		if last != -1 {
			parS[last] = i
			ls[i] = last
		}
		stack = append(stack, i)
	}

	parent = make([]int, n)
	left = make([]int, n)
	right = make([]int, n)
	for si := range n {
		orig := arr[si].Idx - 1
		if parS[si] != -1 {
			parent[orig] = arr[parS[si]].Idx
		}
		if ls[si] != -1 {
			left[orig] = arr[ls[si]].Idx
		}
		if rs[si] != -1 {
			right[orig] = arr[rs[si]].Idx
		}
	}
	return parent, left, right
}

func main() {
	fin, err := os.Open("tree.in")
	if err != nil {
		log.Fatalf("cannot open input file: %v", err)
	}
	defer fin.Close()

	fout, err := os.Create("tree.out")
	if err != nil {
		log.Fatalf("cannot create output file: %v", err)
	}
	defer fout.Close()

	in := bufio.NewReader(fin)
	out := bufio.NewWriter(fout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	nodes := make([]Node, n)
	for i := range n {
		fmt.Fscan(in, &nodes[i].A, &nodes[i].B)
		nodes[i].Idx = i + 1
	}

	parent, left, right := BuildTreap(nodes)
	fmt.Fprintln(out, "YES")
	for i := range n {
		fmt.Fprintf(out, "%d %d %d\n", parent[i], left[i], right[i])
	}
}
