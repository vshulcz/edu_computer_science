package main

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func runTest(commands []string) []int {
	var result []int
	var data []int

	for _, line := range commands {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Insert") {
			numStr := strings.TrimSuffix(strings.TrimPrefix(line, "Insert("), ")")
			num, _ := strconv.Atoi(numStr)
			data = insert(data, num)
		} else if line == "GetMin" {
			result = append(result, data[0])
			data = data[1:]
		} else if line == "GetMax" {
			result = append(result, data[len(data)-1])
			data = data[:len(data)-1]
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
