package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/vshulcz/edu_computer_science/task_4/avl"
)

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

	var root *avl.TreeNode

	for range n {
		line, _ := in.ReadString('\n')
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "Insert") {
			numStr := strings.TrimSuffix(strings.TrimPrefix(line, "Insert("), ")")
			num, _ := strconv.Atoi(numStr)
			root = root.Insert(num)
		} else if line == "GetMin" {
			node := root
			for node.Left != nil {
				node = node.Left
			}
			fmt.Fprintln(out, node.Val)
			root = root.Erase(node.Val)
		} else if line == "GetMax" {
			node := root
			for node.Right != nil {
				node = node.Right
			}
			fmt.Fprintln(out, node.Val)
			root = root.Erase(node.Val)
		}
	}
}
