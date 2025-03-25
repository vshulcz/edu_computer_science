package main

import (
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Case 1",
			nums:     []int{4, 2, 1, 3},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Case 2",
			nums:     []int{-1, 5, 3, 4, 0},
			expected: []int{-1, 0, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := sliceToList(tt.nums)
			result := sortList(input)
			output := listToSlice(result)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("expected %v, but got %v", tt.expected, output)
			}
		})
	}
}

func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for _, n := range nums[1:] {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
