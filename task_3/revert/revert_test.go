package revert

import (
	"reflect"
	"testing"
)

func Test_rotate1(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Case 1",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "Case 2",
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate1(tt.nums, tt.k)
			if !reflect.DeepEqual(tt.nums, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, tt.nums)
			}
		})
	}
}

func Test_rotate2(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Case 1",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "Case 2",
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate2(tt.nums, tt.k)
			if !reflect.DeepEqual(tt.nums, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, tt.nums)
			}
		})
	}
}
