package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Heap struct {
	data []int
}

// First type
func (h *Heap) extractMax() string {
	if len(h.data) == 0 {
		return fmt.Sprintf("%d", -1)
	}
	res := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	idx := siftDown(h.data, 0)

	return fmt.Sprintf("%d %d", idx+1, res)
}

func siftDown(data []int, idx int) int {
	left := 2*idx + 1
	right := 2*idx + 2
	largest := idx

	if left < len(data) && data[left] > data[largest] {
		largest = left
	}
	if right < len(data) && data[right] > data[largest] {
		largest = right
	}
	if largest != idx {
		data[idx], data[largest] = data[largest], data[idx]
		return siftDown(data, largest)
	}
	return idx
}

// Second type
func (h *Heap) insert(x int) string {
	if len(h.data) >= cap(h.data) {
		return fmt.Sprintf("%d", -1)
	}
	h.data = append(h.data, x)
	return fmt.Sprintf("%d", siftUp(h.data, len(h.data)-1))
}

func siftUp(data []int, idx int) int {
	parent := (idx - 1) / 2
	for idx > 0 && data[idx] > data[parent] {
		data[idx], data[parent] = data[parent], data[idx]
		idx = parent
	}
	return idx + 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	params := scanner.Text()
	var N, M int
	fmt.Sscanf(params, "%d %d", &N, &M)

	heap := &Heap{data: make([]int, 0, N)}
	for range M {
		scanner.Scan()
		line := scanner.Text()
		values := []int{}
		sc := bufio.NewScanner(strings.NewReader(line))
		sc.Split(bufio.ScanWords)
		for sc.Scan() {
			val, _ := strconv.Atoi(sc.Text())
			values = append(values, val)
		}
		if values[0] == 1 {
			value := heap.extractMax()
			fmt.Println(value)
		} else if values[0] == 2 {
			index := heap.insert(values[1])
			fmt.Println(index)
		}
	}
	for _, val := range heap.data {
		fmt.Print(val, " ")
	}
}
