package main

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
)

func TestSegmentTreeMinMax(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name: "Case 1",
			input: []string{
				"7",
				"1 3",
				"2 4",
				"-2 -100",
				"1 5",
				"8 9",
				"-3 -101",
				"2 3",
			},
			expected: []string{
				"34",
				"68",
				"250",
				"234",
				"1",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(strings.Join(tt.input, "\n")))
			scanner.Scan()
			k, _ := strconv.Atoi(scanner.Text())

			a := make([]int, N)
			for i := 1; i <= N; i++ {
				a[i-1] = (i * i % 12345) + (i * i * i % 23456)
			}

			st := New(a)
			var output []string

			for i := 0; i < k && scanner.Scan(); i++ {
				parts := strings.Fields(scanner.Text())
				x, _ := strconv.Atoi(parts[0])
				y, _ := strconv.Atoi(parts[1])

				if x > 0 {
					result := st.Query(x, y)
					output = append(output, strconv.Itoa(result))
				} else {
					st.Modify(-x, y)
				}
			}

			if len(output) != len(tt.expected) {
				t.Fatalf("expected %d lines, got %d", len(tt.expected), len(output))
			}
			for i := range output {
				if output[i] != tt.expected[i] {
					t.Errorf("line %d: expected %s, got %s", i+1, tt.expected[i], output[i])
				}
			}
		})
	}
}
