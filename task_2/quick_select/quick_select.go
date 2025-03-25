package main

import (
	"slices"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	return quickSelect(nums, length-k)
}

func quickSelect(nums []int, k int) int {
	pivot := medianOfMedians(nums)
	low, mid, high := partition(nums, pivot)

	if k < len(low) {
		return quickSelect(low, k)
	} else if k < len(low)+len(mid) {
		return pivot
	} else {
		return quickSelect(high, k-len(low)-len(mid))
	}
}

func partition(nums []int, pivot int) (low, mid, high []int) {
	for _, v := range nums {
		switch {
		case v < pivot:
			low = append(low, v)
		case v > pivot:
			high = append(high, v)
		default:
			mid = append(mid, v)
		}
	}
	return
}

// BFPRT
func medianOfMedians(nums []int) int {
	if len(nums) <= 5 {
		sort.Ints(nums)
		return nums[len(nums)/2]
	}

	var medians []int
	for i := 0; i < len(nums); i += 5 {
		end := min(i+5, len(nums))
		group := slices.Clone(nums[i:end])
		sort.Ints(group)
		medians = append(medians, group[len(group)/2])
	}

	return medianOfMedians(medians)
}
