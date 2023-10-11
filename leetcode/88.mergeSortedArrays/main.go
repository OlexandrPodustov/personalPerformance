package main

// min-max size 0...1000
// elements 10^-9...10^9

func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	insertionPoint := len(nums1) - 1
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[insertionPoint] = nums1[m]
			m--
		} else {
			nums1[insertionPoint] = nums2[n]
			n--
		}
		insertionPoint--

	}

	if m <= 0 {
		for i := n; i >= 0; i-- {
			nums1[insertionPoint] = nums2[i]
			insertionPoint--
		}
	}

	if n <= 0 {
		for i := m; i >= 0; i-- {
			nums1[insertionPoint] = nums1[i]
			insertionPoint--
		}
	}

	return
}
