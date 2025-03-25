package main

import (
	"reflect"
	"testing"
)

func TestHeapOperations(t *testing.T) {
	tests := []struct {
		name           string
		heapSize       int
		operations     [][]int
		expectedOutput []string
		finalHeap      []int
	}{
		{
			name:     "Case 1",
			heapSize: 4,
			operations: [][]int{
				{1},
				{2, 9},
				{2, 4},
				{2, 9},
				{2, 9},
				{2, 7},
				{1},
			},
			expectedOutput: []string{"-1", "1", "2", "3", "2", "-1", "2 9"},
			finalHeap:      []int{9, 4, 9},
		},
		{
			name:     "Case 2",
			heapSize: 1,
			operations: [][]int{
				{1},
			},
			expectedOutput: []string{"-1"},
			finalHeap:      []int{},
		},
		{
			name:     "Case 3",
			heapSize: 3,
			operations: [][]int{
				{2, 5},
				{2, 10},
				{2, 3},
				{2, 8},
			},
			expectedOutput: []string{"1", "1", "3", "-1"},
			finalHeap:      []int{10, 5, 3},
		},
		{
			name:     "Case 4",
			heapSize: 2,
			operations: [][]int{
				{2, 15},
				{1},
				{2, 20},
				{1},
				{1},
				{1},
			},
			expectedOutput: []string{"1", "1 15", "1", "1 20", "-1", "-1"},
			finalHeap:      []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap := &Heap{data: make([]int, 0, tt.heapSize)}
			outputs := []string{}
			for _, op := range tt.operations {
				if op[0] == 1 {
					outputs = append(outputs, heap.extractMax())
				} else if op[0] == 2 {
					outputs = append(outputs, heap.insert(op[1]))
				}
			}
			if !reflect.DeepEqual(outputs, tt.expectedOutput) {
				t.Errorf("expected output %v, but got %v", tt.expectedOutput, outputs)
			}
			if !reflect.DeepEqual(heap.data, tt.finalHeap) {
				t.Errorf("expected final heap %v, but got %v", tt.finalHeap, heap.data)
			}
		})
	}
}
