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
	for {
		pivot := medianOfMedians(nums)
		lt, gt := partition(nums, pivot)

		if k < lt {
			nums = nums[:lt]
		} else if k < gt {
			return pivot
		} else {
			nums = nums[gt:]
			k -= gt
		}
	}
}

func partition(nums []int, pivot int) (lt int, gt int) {
	n := len(nums)
	lt, i, gt := 0, 0, n

	for i < gt {
		switch {
		case nums[i] < pivot:
			nums[lt], nums[i] = nums[i], nums[lt]
			lt++
			i++
		case nums[i] > pivot:
			gt--
			nums[i], nums[gt] = nums[gt], nums[i]
		default:
			i++
		}
	}
	return lt, gt
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
