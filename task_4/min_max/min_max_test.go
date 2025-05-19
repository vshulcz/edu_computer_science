package main

import (
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/vshulcz/edu_computer_science/task_4/avl"
)

func runTest(commands []string) []int {
	var result []int
	var root *avl.TreeNode

	for _, line := range commands {
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
			result = append(result, node.Val)
			root = root.Erase(node.Val)
		} else if line == "GetMax" {
			node := root
			for node.Right != nil {
				node = node.Right
			}
			result = append(result, node.Val)
			root = root.Erase(node.Val)
		}
	}

	return result
}

func Test_Main(t *testing.T) {
	tests := []struct {
		name     string
		commands []string
		expected []int
	}{
		{
			name: "Case 1",
			commands: []string{
				"Insert(100)",
				"Insert(99)",
				"Insert(1)",
				"Insert(2)",
				"GetMin",
				"GetMax",
				"Insert(1)",
				"GetMin",
				"GetMin",
				"GetMax",
			},
			expected: []int{1, 100, 1, 2, 99},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := runTest(tt.commands)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("runTest() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
