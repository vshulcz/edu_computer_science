package main

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name: "Case 1",
			input: []string{
				"5 9",
				"A 2 2",
				"A 3 1",
				"A 4 2",
				"Q 1 1",
				"Q 2 2",
				"Q 3 3",
				"Q 4 4",
				"Q 5 5",
				"Q 1 5",
			},
			expected: []string{
				"0", // a[0]
				"2", // a[1]
				"1", // a[2]
				"2", // a[3]
				"0", // a[4]
				"5", // sum of a[0..4] = 0+2+1+2+0
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(strings.Join(tt.input, "\n")))
			scanner.Scan()
			header := strings.Fields(scanner.Text())
			n, _ := strconv.Atoi(header[0])
			k, _ := strconv.Atoi(header[1])

			st := New(n)
			var result []string

			for i := 0; i < k && scanner.Scan(); i++ {
				parts := strings.Fields(scanner.Text())
				switch parts[0] {
				case "A":
					idx, _ := strconv.Atoi(parts[1])
					val, _ := strconv.Atoi(parts[2])
					st.Modify(idx-1, val)
				case "Q":
					l, _ := strconv.Atoi(parts[1])
					r, _ := strconv.Atoi(parts[2])
					sum := st.Query(l-1, r)
					result = append(result, strconv.Itoa(sum))
				}
			}

			if len(result) != len(tt.expected) {
				t.Fatalf("expected %d results, got %d", len(tt.expected), len(result))
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("query %d: expected %s, got %s", i+1, tt.expected[i], result[i])
				}
			}
		})
	}
}
