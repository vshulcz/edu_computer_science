package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func insert(sorted []int, val int) []int {
	left, right := 0, len(sorted)
	for left < right {
		mid := (left + right) / 2
		if sorted[mid] < val {
			left = mid + 1
		} else {
			right = mid
		}
	}
	sorted = append(sorted, 0)
	copy(sorted[left+1:], sorted[left:])
	sorted[left] = val
	return sorted
}

func main() {
	fin, err := os.Open("minmax.in")
	if err != nil {
		log.Fatalf("cannot open input file: %v", err)
	}
	defer fin.Close()

	fout, err := os.Create("minmax.out")
	if err != nil {
		log.Fatalf("cannot create output file: %v", err)
	}
	defer fout.Close()

	in := bufio.NewReader(fin)
	out := bufio.NewWriter(fout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	in.ReadString('\n')

	var data []int

	for range n {
		line, _ := in.ReadString('\n')
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "Insert") {
			numStr := strings.TrimSuffix(strings.TrimPrefix(line, "Insert("), ")")
			num, _ := strconv.Atoi(numStr)
			data = insert(data, num)
		} else if line == "GetMin" {
			fmt.Fprintln(out, data[0])
			data = data[1:]
		} else if line == "GetMax" {
			fmt.Fprintln(out, data[len(data)-1])
			data = data[:len(data)-1]
		}
	}
}
