package revert

// swap = n/2 + k/2 + (n-k)/2 = n
func rotate1(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// swap = n
func rotate2(nums []int, k int) {
	n := len(nums)
	k = k % n
	count := 0
	for start := 0; count < n; start++ {
		current := start
		prev := nums[start]

		for {
			next := (current + k) % n
			nums[next], prev = prev, nums[next]
			current = next
			count++

			if start == current {
				break
			}
		}
	}
}
